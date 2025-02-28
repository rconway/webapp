import React, { useState, useEffect } from 'react';
import { BrowserRouter, Routes, Route, Link, Navigate } from 'react-router-dom';
import keycloak from './keycloak';
import './App.css';

// Home component
const Home = () => (
  <div>
    <h1>Home</h1>
    <p>This is a public page</p>
  </div>
);

// Secured component
const Secured = () => {
  if (!keycloak.authenticated) {
    return <Navigate to="/" />;
  }
  
  return (
    <div>
      <h1>Secured Content</h1>
      <p>This page is protected and only visible after login</p>
      <div>
        <h2>User Information</h2>
        <p>Username: {keycloak.tokenParsed?.preferred_username}</p>
        <p>Email: {keycloak.tokenParsed?.email}</p>
        <pre>{JSON.stringify(keycloak.tokenParsed, null, 2)}</pre>
      </div>
    </div>
  );
};

function App() {
  const [initialized, setInitialized] = useState(false);
  
  useEffect(() => {
    console.log("zzz: calling keycloak.init()")
    // Initialize Keycloak with PKCE
    keycloak.init({
      onLoad: 'check-sso',
      silentCheckSsoRedirectUri: window.location.origin + '/silent-check-sso.html',
      pkceMethod: 'S256', // Enable PKCE
    }).then(authenticated => {
      setInitialized(true);
      console.log(`User is ${authenticated ? 'authenticated' : 'not authenticated'}`);
    }).catch(error => {
      console.error('Failed to initialize Keycloak', error);
    });
    
    return () => {
      // Optional: Handle cleanup
    };
  }, []);

  if (!initialized) {
    return <div>Loading...</div>;
  }

  return (
    <BrowserRouter>
      <div className="App">
        <nav>
          <ul>
            <li><Link to="/">Home</Link></li>
            <li><Link to="/secured">Secured Page</Link></li>
            {!keycloak.authenticated ? (
              <li><button onClick={() => keycloak.login()}>Login</button></li>
            ) : (
              <li><button onClick={() => keycloak.logout()}>Logout</button></li>
            )}
          </ul>
        </nav>

        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/secured" element={<Secured />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
