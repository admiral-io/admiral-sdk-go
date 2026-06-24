package client

import (
	agentv1 "go.admiral.io/sdk/proto/admiral/api/agent/v1"
	applicationv1 "go.admiral.io/sdk/proto/admiral/api/application/v1"
	authenticationv1 "go.admiral.io/sdk/proto/admiral/api/authentication/v1"
	catalogv1 "go.admiral.io/sdk/proto/admiral/api/catalog/v1"
	changesetv1 "go.admiral.io/sdk/proto/admiral/api/changeset/v1"
	credentialv1 "go.admiral.io/sdk/proto/admiral/api/credential/v1"
	environmentv1 "go.admiral.io/sdk/proto/admiral/api/environment/v1"
	healthcheckv1 "go.admiral.io/sdk/proto/admiral/api/healthcheck/v1"
	runv1 "go.admiral.io/sdk/proto/admiral/api/run/v1"
	sourcev1 "go.admiral.io/sdk/proto/admiral/api/source/v1"
	userv1 "go.admiral.io/sdk/proto/admiral/api/user/v1"
)

// AdmiralClient provides access to Admiral service clients.
type AdmiralClient interface {
	// Agent returns the AgentAPI client.
	Agent() agentv1.AgentAPIClient
	// Application returns the ApplicationAPI client.
	Application() applicationv1.ApplicationAPIClient
	// Authentication returns the AuthenticationAPI client.
	Authentication() authenticationv1.AuthenticationAPIClient
	// Catalog returns the CatalogAPI client.
	Catalog() catalogv1.CatalogAPIClient
	// ChangeSet returns the ChangeSetAPI client.
	ChangeSet() changesetv1.ChangeSetAPIClient
	// Credential returns the CredentialAPI client.
	Credential() credentialv1.CredentialAPIClient
	// Environment returns the EnvironmentAPI client.
	Environment() environmentv1.EnvironmentAPIClient
	// Healthcheck returns the HealthcheckAPI client.
	Healthcheck() healthcheckv1.HealthcheckAPIClient
	// Run returns the RunAPI client.
	Run() runv1.RunAPIClient
	// Source returns the SourceAPI client.
	Source() sourcev1.SourceAPIClient
	// User returns the UserAPI client.
	User() userv1.UserAPIClient

	// ValidateToken validates the client's auth token.
	ValidateToken() error

	// Version returns the client library version string.
	Version() string

	// Close closes the underlying connection.
	Close() error
}
