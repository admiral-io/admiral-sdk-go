package client

import (
	applicationv1 "go.admiral.io/sdk/proto/admiral/application/v1"
	authenticationv1 "go.admiral.io/sdk/proto/admiral/authentication/v1"
	changesetv1 "go.admiral.io/sdk/proto/admiral/changeset/v1"
	clusterv1 "go.admiral.io/sdk/proto/admiral/cluster/v1"
	credentialv1 "go.admiral.io/sdk/proto/admiral/credential/v1"
	environmentv1 "go.admiral.io/sdk/proto/admiral/environment/v1"
	healthcheckv1 "go.admiral.io/sdk/proto/admiral/healthcheck/v1"
	modulev1 "go.admiral.io/sdk/proto/admiral/module/v1"
	runv1 "go.admiral.io/sdk/proto/admiral/run/v1"
	runnerv1 "go.admiral.io/sdk/proto/admiral/runner/v1"
	sourcev1 "go.admiral.io/sdk/proto/admiral/source/v1"
	statev1 "go.admiral.io/sdk/proto/admiral/state/v1"
	userv1 "go.admiral.io/sdk/proto/admiral/user/v1"
)

// AdmiralClient provides access to Admiral service clients.
type AdmiralClient interface {
	// Application returns the ApplicationAPI client.
	Application() applicationv1.ApplicationAPIClient
	// Authentication returns the AuthenticationAPI client.
	Authentication() authenticationv1.AuthenticationAPIClient
	// ChangeSet returns the ChangeSetAPI client.
	ChangeSet() changesetv1.ChangeSetAPIClient
	// Cluster returns the ClusterAPI client.
	Cluster() clusterv1.ClusterAPIClient
	// Credential returns the CredentialAPI client.
	Credential() credentialv1.CredentialAPIClient
	// Environment returns the EnvironmentAPI client.
	Environment() environmentv1.EnvironmentAPIClient
	// Healthcheck returns the HealthcheckAPI client.
	Healthcheck() healthcheckv1.HealthcheckAPIClient
	// Module returns the ModuleAPI client.
	Module() modulev1.ModuleAPIClient
	// Run returns the RunAPI client.
	Run() runv1.RunAPIClient
	// Runner returns the RunnerAPI client.
	Runner() runnerv1.RunnerAPIClient
	// Source returns the SourceAPI client.
	Source() sourcev1.SourceAPIClient
	// State returns the StateAPI client.
	State() statev1.StateAPIClient
	// User returns the UserAPI client.
	User() userv1.UserAPIClient

	// ValidateToken validates the client's auth token.
	ValidateToken() error

	// Version returns the client library version string.
	Version() string

	// Close closes the underlying connection.
	Close() error
}
