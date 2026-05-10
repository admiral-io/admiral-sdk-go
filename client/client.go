package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
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

// Compile-time check that Client implements AdmiralClient
var _ AdmiralClient = (*Client)(nil)

// Client is the Admiral API client.
type Client struct {
	conn           *grpc.ClientConn
	logger         Logger
	authToken      string
	tokenValidator TokenValidator
	application applicationv1.ApplicationAPIClient
	authentication authenticationv1.AuthenticationAPIClient
	changeSet changesetv1.ChangeSetAPIClient
	cluster clusterv1.ClusterAPIClient
	credential credentialv1.CredentialAPIClient
	environment environmentv1.EnvironmentAPIClient
	healthcheck healthcheckv1.HealthcheckAPIClient
	module modulev1.ModuleAPIClient
	run runv1.RunAPIClient
	runner runnerv1.RunnerAPIClient
	source sourcev1.SourceAPIClient
	state statev1.StateAPIClient
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
		application: applicationv1.NewApplicationAPIClient(conn),
		authentication: authenticationv1.NewAuthenticationAPIClient(conn),
		changeSet: changesetv1.NewChangeSetAPIClient(conn),
		cluster: clusterv1.NewClusterAPIClient(conn),
		credential: credentialv1.NewCredentialAPIClient(conn),
		environment: environmentv1.NewEnvironmentAPIClient(conn),
		healthcheck: healthcheckv1.NewHealthcheckAPIClient(conn),
		module: modulev1.NewModuleAPIClient(conn),
		run: runv1.NewRunAPIClient(conn),
		runner: runnerv1.NewRunnerAPIClient(conn),
		source: sourcev1.NewSourceAPIClient(conn),
		state: statev1.NewStateAPIClient(conn),
		user: userv1.NewUserAPIClient(conn),
	}, nil
}

// Application returns the ApplicationAPI client.
func (c *Client) Application() applicationv1.ApplicationAPIClient {
	return c.application
}

// Authentication returns the AuthenticationAPI client.
func (c *Client) Authentication() authenticationv1.AuthenticationAPIClient {
	return c.authentication
}

// ChangeSet returns the ChangeSetAPI client.
func (c *Client) ChangeSet() changesetv1.ChangeSetAPIClient {
	return c.changeSet
}

// Cluster returns the ClusterAPI client.
func (c *Client) Cluster() clusterv1.ClusterAPIClient {
	return c.cluster
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

// Module returns the ModuleAPI client.
func (c *Client) Module() modulev1.ModuleAPIClient {
	return c.module
}

// Run returns the RunAPI client.
func (c *Client) Run() runv1.RunAPIClient {
	return c.run
}

// Runner returns the RunnerAPI client.
func (c *Client) Runner() runnerv1.RunnerAPIClient {
	return c.runner
}

// Source returns the SourceAPI client.
func (c *Client) Source() sourcev1.SourceAPIClient {
	return c.source
}

// State returns the StateAPI client.
func (c *Client) State() statev1.StateAPIClient {
	return c.state
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
