import Keycloak from 'keycloak-js';

// Configure Keycloak instance
const keycloakConfig = {
  url: 'https://auth.test.eoepca.org', // Change this to your Keycloak server URL
  realm: 'eoepca',
  clientId: 'openeo-public'
};

// Initialize Keycloak instance
const keycloak = new Keycloak(keycloakConfig);

export default keycloak;
