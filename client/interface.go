package client

import (
	applicationv1 "go.admiral.io/sdk/proto/admiral/api/application/v1"
	clusterv1 "go.admiral.io/sdk/proto/admiral/api/cluster/v1"
	componentv1 "go.admiral.io/sdk/proto/admiral/api/component/v1"
	connectionv1 "go.admiral.io/sdk/proto/admiral/api/connection/v1"
	deploymentv1 "go.admiral.io/sdk/proto/admiral/api/deployment/v1"
	environmentv1 "go.admiral.io/sdk/proto/admiral/api/environment/v1"
	healthcheckv1 "go.admiral.io/sdk/proto/admiral/api/healthcheck/v1"
	runnerv1 "go.admiral.io/sdk/proto/admiral/api/runner/v1"
	sourcev1 "go.admiral.io/sdk/proto/admiral/api/source/v1"
	statev1 "go.admiral.io/sdk/proto/admiral/api/state/v1"
	userv1 "go.admiral.io/sdk/proto/admiral/api/user/v1"
	variablev1 "go.admiral.io/sdk/proto/admiral/api/variable/v1"
)

// AdmiralClient provides access to Admiral service clients.
type AdmiralClient interface {
	// Application returns the ApplicationAPI client.
	Application() applicationv1.ApplicationAPIClient
	// Cluster returns the ClusterAPI client.
	Cluster() clusterv1.ClusterAPIClient
	// Component returns the ComponentAPI client.
	Component() componentv1.ComponentAPIClient
	// Connection returns the ConnectionAPI client.
	Connection() connectionv1.ConnectionAPIClient
	// Deployment returns the DeploymentAPI client.
	Deployment() deploymentv1.DeploymentAPIClient
	// Environment returns the EnvironmentAPI client.
	Environment() environmentv1.EnvironmentAPIClient
	// Healthcheck returns the HealthcheckAPI client.
	Healthcheck() healthcheckv1.HealthcheckAPIClient
	// Runner returns the RunnerAPI client.
	Runner() runnerv1.RunnerAPIClient
	// Source returns the SourceAPI client.
	Source() sourcev1.SourceAPIClient
	// State returns the StateAPI client.
	State() statev1.StateAPIClient
	// User returns the UserAPI client.
	User() userv1.UserAPIClient
	// Variable returns the VariableAPI client.
	Variable() variablev1.VariableAPIClient

	// ValidateToken validates the client's auth token.
	ValidateToken() error

	// GetTokenInfo returns information about the client's auth token.
	GetTokenInfo() (*TokenInfo, error)

	// Version returns the client library version string.
	Version() string

	// Close closes the underlying connection.
	Close() error
}
