const clientId = 'openeo-public';
const redirectUri = window.location.origin;
console.log("Using Redirect URI: " + redirectUri)
const scope = 'openid profile email';
const oidcConfigUrl = 'https://auth.test.eoepca.org/realms/eoepca/.well-known/openid-configuration';
let authEndpoint, tokenEndpoint, userInfoEndpoint;

async function fetchOidcConfig() {
  const response = await fetch(oidcConfigUrl);
  const oidcConfig = await response.json();
  authEndpoint = oidcConfig.authorization_endpoint;
  tokenEndpoint = oidcConfig.token_endpoint;
  userInfoEndpoint = oidcConfig.userinfo_endpoint;
}

function generateRandomString(length) {
  const charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  let randomString = '';
  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * charset.length);
    randomString += charset.charAt(randomIndex);
  }
  return randomString;
}

function base64UrlEncode(str) {
  return btoa(str).replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '');
}

function sha256(plain) {
  const shaObj = new jsSHA('SHA-256', 'TEXT');
  shaObj.update(plain);
  const hash = shaObj.getHash('BYTES');
  return hash;
}

async function login() {
  const codeVerifier = generateRandomString(128);
  sessionStorage.setItem('codeVerifier', codeVerifier);

  const codeChallenge = base64UrlEncode(sha256(codeVerifier));
  const state = generateRandomString(16);

  await fetchOidcConfig();

  const authUrl = `${authEndpoint}?response_type=code&client_id=${clientId}&redirect_uri=${redirectUri}&scope=${scope}&code_challenge=${codeChallenge}&code_challenge_method=S256&state=${state}`;
  window.location = authUrl;
}

async function handleCallback(redirectTo) {
  const urlParams = new URLSearchParams(window.location.search);
  const code = urlParams.get('code');
  const state = urlParams.get('state');
  const codeVerifier = sessionStorage.getItem('codeVerifier');

  if (code) {
    await fetchOidcConfig();

    const tokenResponse = await fetch(tokenEndpoint, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: new URLSearchParams({
        grant_type: 'authorization_code',
        code: code,
        client_id: clientId,
        redirect_uri: redirectUri,
        code_verifier: codeVerifier
      })
    });
    const tokenData = await tokenResponse.json();
    sessionStorage.setItem('accessToken', tokenData.access_token);
    await fetchUserInfo(tokenData.access_token);
  }

  window.location = redirectTo;
}

async function fetchUserInfo(accessToken) {
  const userInfoResponse = await fetch(userInfoEndpoint, {
    headers: {
      'Authorization': `Bearer ${accessToken}`
    }
  });
  const userInfo = await userInfoResponse.json();
  localStorage.setItem('userInfo', JSON.stringify(userInfo));
}

async function home() {
  const urlParams = new URLSearchParams(window.location.search);
  const code = urlParams.get('code');

  if (code) {
    await handleCallback(window.location.origin);
  }

  const userInfo = JSON.parse(localStorage.getItem('userInfo'));
  if (userInfo) {
    document.getElementById('userInfo').innerText = `Hello, ${userInfo.name}`;
  } else {
    console.log("Userinfo is not set");
    document.getElementById('userInfo').innerText = '';
  }
}

function logout() {
  sessionStorage.removeItem('accessToken');
  localStorage.removeItem('userInfo');
  document.getElementById('userInfo').innerText = '';
}

document.getElementById('loginButton').addEventListener('click', login);
document.getElementById('logoutButton').addEventListener('click', logout);

window.onload = home;
