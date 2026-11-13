package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
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

// Compile-time check that Client implements AdmiralClient
var _ AdmiralClient = (*Client)(nil)

// Client is the Admiral API client.
type Client struct {
	conn           *grpc.ClientConn
	logger         Logger
	authToken      string
	tokenValidator TokenValidator
	agent agentv1.AgentAPIClient
	job agentv1.JobAPIClient
	agentRuntime agentv1.AgentRuntimeAPIClient
	apiKey apikeyv1.ApiKeyAPIClient
	changeSet changesetv1.ChangeSetAPIClient
	credential credentialv1.CredentialAPIClient
	environment environmentv1.EnvironmentAPIClient
	healthcheck healthcheckv1.HealthcheckAPIClient
	integration integrationv1.IntegrationAPIClient
	invitation invitationv1.InvitationAPIClient
	organization organizationv1.OrganizationAPIClient
	project projectv1.ProjectAPIClient
	registry registryv1.RegistryAPIClient
	run runv1.RunAPIClient
	serviceAccount serviceaccountv1.ServiceAccountAPIClient
	issuer serviceaccountv1.IssuerAPIClient
	source sourcev1.SourceAPIClient
	user userv1.UserAPIClient
}

// New creates a new Admiral client with the given configuration.
func New(ctx context.Context, cfg Config) (*Client, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	dialOpts := cfg.ConnectionOptions.DialOptions

	// Configure transport credentials
	if cfg.ConnectionOptions.Insecure {
		// Opt-in plaintext for local dev or a TLS-terminating sidecar; TLS is the default.
		// nosemgrep
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		tlsConfig := cfg.ConnectionOptions.TLSConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{MinVersion: tls.VersionTLS12}
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	}

	// Add user agent
	dialOpts = append(dialOpts, grpc.WithUserAgent(ClientUserAgent()))

	// Dial with timeout
	dialCtx, cancel := context.WithTimeout(ctx, cfg.ConnectionOptions.DialTimeout)
	defer cancel()

	conn, err := grpc.DialContext(dialCtx, cfg.HostPort, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %s: %w", cfg.HostPort, err)
	}

	cfg.Logger.Debugf("connected to Admiral API at %s", cfg.HostPort)

	return &Client{
		conn:           conn,
		logger:         cfg.Logger,
		authToken:      cfg.AuthToken,
		tokenValidator: cfg.TokenValidator,
		agent: agentv1.NewAgentAPIClient(conn),
		job: agentv1.NewJobAPIClient(conn),
		agentRuntime: agentv1.NewAgentRuntimeAPIClient(conn),
		apiKey: apikeyv1.NewApiKeyAPIClient(conn),
		changeSet: changesetv1.NewChangeSetAPIClient(conn),
		credential: credentialv1.NewCredentialAPIClient(conn),
		environment: environmentv1.NewEnvironmentAPIClient(conn),
		healthcheck: healthcheckv1.NewHealthcheckAPIClient(conn),
		integration: integrationv1.NewIntegrationAPIClient(conn),
		invitation: invitationv1.NewInvitationAPIClient(conn),
		organization: organizationv1.NewOrganizationAPIClient(conn),
		project: projectv1.NewProjectAPIClient(conn),
		registry: registryv1.NewRegistryAPIClient(conn),
		run: runv1.NewRunAPIClient(conn),
		serviceAccount: serviceaccountv1.NewServiceAccountAPIClient(conn),
		issuer: serviceaccountv1.NewIssuerAPIClient(conn),
		source: sourcev1.NewSourceAPIClient(conn),
		user: userv1.NewUserAPIClient(conn),
	}, nil
}

// Agent returns the AgentAPI client.
func (c *Client) Agent() agentv1.AgentAPIClient {
	return c.agent
}

// Job returns the JobAPI client.
func (c *Client) Job() agentv1.JobAPIClient {
	return c.job
}

// AgentRuntime returns the AgentRuntimeAPI client.
func (c *Client) AgentRuntime() agentv1.AgentRuntimeAPIClient {
	return c.agentRuntime
}

// ApiKey returns the ApiKeyAPI client.
func (c *Client) ApiKey() apikeyv1.ApiKeyAPIClient {
	return c.apiKey
}

// ChangeSet returns the ChangeSetAPI client.
func (c *Client) ChangeSet() changesetv1.ChangeSetAPIClient {
	return c.changeSet
}

// Credential returns the CredentialAPI client.
func (c *Client) Credential() credentialv1.CredentialAPIClient {
	return c.credential
}

// Environment returns the EnvironmentAPI client.
func (c *Client) Environment() environmentv1.EnvironmentAPIClient {
	return c.environment
}

// Healthcheck returns the HealthcheckAPI client.
func (c *Client) Healthcheck() healthcheckv1.HealthcheckAPIClient {
	return c.healthcheck
}

// Integration returns the IntegrationAPI client.
func (c *Client) Integration() integrationv1.IntegrationAPIClient {
	return c.integration
}

// Invitation returns the InvitationAPI client.
func (c *Client) Invitation() invitationv1.InvitationAPIClient {
	return c.invitation
}

// Organization returns the OrganizationAPI client.
func (c *Client) Organization() organizationv1.OrganizationAPIClient {
	return c.organization
}

// Project returns the ProjectAPI client.
func (c *Client) Project() projectv1.ProjectAPIClient {
	return c.project
}

// Registry returns the RegistryAPI client.
func (c *Client) Registry() registryv1.RegistryAPIClient {
	return c.registry
}

// Run returns the RunAPI client.
func (c *Client) Run() runv1.RunAPIClient {
	return c.run
}

// ServiceAccount returns the ServiceAccountAPI client.
func (c *Client) ServiceAccount() serviceaccountv1.ServiceAccountAPIClient {
	return c.serviceAccount
}

// Issuer returns the IssuerAPI client.
func (c *Client) Issuer() serviceaccountv1.IssuerAPIClient {
	return c.issuer
}

// Source returns the SourceAPI client.
func (c *Client) Source() sourcev1.SourceAPIClient {
	return c.source
}

// User returns the UserAPI client.
func (c *Client) User() userv1.UserAPIClient {
	return c.user
}

// ValidateToken validates the client's auth token format.
func (c *Client) ValidateToken() error {
	return c.tokenValidator.Validate(c.authToken)
}

// Version returns the client library version string.
func (c *Client) Version() string {
	return Version()
}

// Close closes the underlying gRPC connection.
func (c *Client) Close() error {
	if c.conn != nil {
		c.logger.Debugf("closing connection")
		return c.conn.Close()
	}
	return nil
}
