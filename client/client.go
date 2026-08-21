package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	agentv1 "go.admiral.io/sdk/proto/admiral/api/agent/v1"
	applicationv1 "go.admiral.io/sdk/proto/admiral/api/application/v1"
	catalogv1 "go.admiral.io/sdk/proto/admiral/api/catalog/v1"
	changesetv1 "go.admiral.io/sdk/proto/admiral/api/changeset/v1"
	credentialv1 "go.admiral.io/sdk/proto/admiral/api/credential/v1"
	environmentv1 "go.admiral.io/sdk/proto/admiral/api/environment/v1"
	healthcheckv1 "go.admiral.io/sdk/proto/admiral/api/healthcheck/v1"
	runv1 "go.admiral.io/sdk/proto/admiral/api/run/v1"
	sourcev1 "go.admiral.io/sdk/proto/admiral/api/source/v1"
	tenantv1 "go.admiral.io/sdk/proto/admiral/api/tenant/v1"
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
	agentRuntime agentv1.AgentRuntimeAPIClient
	application applicationv1.ApplicationAPIClient
	catalog catalogv1.CatalogAPIClient
	changeSet changesetv1.ChangeSetAPIClient
	credential credentialv1.CredentialAPIClient
	environment environmentv1.EnvironmentAPIClient
	healthcheck healthcheckv1.HealthcheckAPIClient
	run runv1.RunAPIClient
	source sourcev1.SourceAPIClient
	tenant tenantv1.TenantAPIClient
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
		agentRuntime: agentv1.NewAgentRuntimeAPIClient(conn),
		application: applicationv1.NewApplicationAPIClient(conn),
		catalog: catalogv1.NewCatalogAPIClient(conn),
		changeSet: changesetv1.NewChangeSetAPIClient(conn),
		credential: credentialv1.NewCredentialAPIClient(conn),
		environment: environmentv1.NewEnvironmentAPIClient(conn),
		healthcheck: healthcheckv1.NewHealthcheckAPIClient(conn),
		run: runv1.NewRunAPIClient(conn),
		source: sourcev1.NewSourceAPIClient(conn),
		tenant: tenantv1.NewTenantAPIClient(conn),
		user: userv1.NewUserAPIClient(conn),
	}, nil
}

// Agent returns the AgentAPI client.
func (c *Client) Agent() agentv1.AgentAPIClient {
	return c.agent
}

// AgentRuntime returns the AgentRuntimeAPI client.
func (c *Client) AgentRuntime() agentv1.AgentRuntimeAPIClient {
	return c.agentRuntime
}

// Application returns the ApplicationAPI client.
func (c *Client) Application() applicationv1.ApplicationAPIClient {
	return c.application
}

// Catalog returns the CatalogAPI client.
func (c *Client) Catalog() catalogv1.CatalogAPIClient {
	return c.catalog
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

// Run returns the RunAPI client.
func (c *Client) Run() runv1.RunAPIClient {
	return c.run
}

// Source returns the SourceAPI client.
func (c *Client) Source() sourcev1.SourceAPIClient {
	return c.source
}

// Tenant returns the TenantAPI client.
func (c *Client) Tenant() tenantv1.TenantAPIClient {
	return c.tenant
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
