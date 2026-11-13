package client

import (
	agentv1 "go.admiral.io/sdk/proto/admiral/api/agent/v1"
	apikeyv1 "go.admiral.io/sdk/proto/admiral/api/apikey/v1"
	changesetv1 "go.admiral.io/sdk/proto/admiral/api/changeset/v1"
	credentialv1 "go.admiral.io/sdk/proto/admiral/api/credential/v1"
	environmentv1 "go.admiral.io/sdk/proto/admiral/api/environment/v1"
	healthcheckv1 "go.admiral.io/sdk/proto/admiral/api/healthcheck/v1"
	integrationv1 "go.admiral.io/sdk/proto/admiral/api/integration/v1"
	invitationv1 "go.admiral.io/sdk/proto/admiral/api/invitation/v1"
	organizationv1 "go.admiral.io/sdk/proto/admiral/api/organization/v1"
	projectv1 "go.admiral.io/sdk/proto/admiral/api/project/v1"
	registryv1 "go.admiral.io/sdk/proto/admiral/api/registry/v1"
	runv1 "go.admiral.io/sdk/proto/admiral/api/run/v1"
	serviceaccountv1 "go.admiral.io/sdk/proto/admiral/api/serviceaccount/v1"
	sourcev1 "go.admiral.io/sdk/proto/admiral/api/source/v1"
	userv1 "go.admiral.io/sdk/proto/admiral/api/user/v1"
)

// AdmiralClient provides access to Admiral service clients.
type AdmiralClient interface {
	// Agent returns the AgentAPI client.
	Agent() agentv1.AgentAPIClient
	// Job returns the JobAPI client.
	Job() agentv1.JobAPIClient
	// AgentRuntime returns the AgentRuntimeAPI client.
	AgentRuntime() agentv1.AgentRuntimeAPIClient
	// ApiKey returns the ApiKeyAPI client.
	ApiKey() apikeyv1.ApiKeyAPIClient
	// ChangeSet returns the ChangeSetAPI client.
	ChangeSet() changesetv1.ChangeSetAPIClient
	// Credential returns the CredentialAPI client.
	Credential() credentialv1.CredentialAPIClient
	// Environment returns the EnvironmentAPI client.
	Environment() environmentv1.EnvironmentAPIClient
	// Healthcheck returns the HealthcheckAPI client.
	Healthcheck() healthcheckv1.HealthcheckAPIClient
	// Integration returns the IntegrationAPI client.
	Integration() integrationv1.IntegrationAPIClient
	// Invitation returns the InvitationAPI client.
	Invitation() invitationv1.InvitationAPIClient
	// Organization returns the OrganizationAPI client.
	Organization() organizationv1.OrganizationAPIClient
	// Project returns the ProjectAPI client.
	Project() projectv1.ProjectAPIClient
	// Registry returns the RegistryAPI client.
	Registry() registryv1.RegistryAPIClient
	// Run returns the RunAPI client.
	Run() runv1.RunAPIClient
	// ServiceAccount returns the ServiceAccountAPI client.
	ServiceAccount() serviceaccountv1.ServiceAccountAPIClient
	// Issuer returns the IssuerAPI client.
	Issuer() serviceaccountv1.IssuerAPIClient
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
