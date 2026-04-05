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
	clusterv1 "go.admiral.io/sdk/proto/admiral/cluster/v1"
	componentv1 "go.admiral.io/sdk/proto/admiral/component/v1"
	connectionv1 "go.admiral.io/sdk/proto/admiral/connection/v1"
	deploymentv1 "go.admiral.io/sdk/proto/admiral/deployment/v1"
	environmentv1 "go.admiral.io/sdk/proto/admiral/environment/v1"
	healthcheckv1 "go.admiral.io/sdk/proto/admiral/healthcheck/v1"
	runnerv1 "go.admiral.io/sdk/proto/admiral/runner/v1"
	sourcev1 "go.admiral.io/sdk/proto/admiral/source/v1"
	statev1 "go.admiral.io/sdk/proto/admiral/state/v1"
	userv1 "go.admiral.io/sdk/proto/admiral/user/v1"
	variablev1 "go.admiral.io/sdk/proto/admiral/variable/v1"
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
	cluster clusterv1.ClusterAPIClient
	component componentv1.ComponentAPIClient
	connection connectionv1.ConnectionAPIClient
	deployment deploymentv1.DeploymentAPIClient
	environment environmentv1.EnvironmentAPIClient
	healthcheck healthcheckv1.HealthcheckAPIClient
	runner runnerv1.RunnerAPIClient
	source sourcev1.SourceAPIClient
	state statev1.StateAPIClient
	user userv1.UserAPIClient
	variable variablev1.VariableAPIClient
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
		cluster: clusterv1.NewClusterAPIClient(conn),
		component: componentv1.NewComponentAPIClient(conn),
		connection: connectionv1.NewConnectionAPIClient(conn),
		deployment: deploymentv1.NewDeploymentAPIClient(conn),
		environment: environmentv1.NewEnvironmentAPIClient(conn),
		healthcheck: healthcheckv1.NewHealthcheckAPIClient(conn),
		runner: runnerv1.NewRunnerAPIClient(conn),
		source: sourcev1.NewSourceAPIClient(conn),
		state: statev1.NewStateAPIClient(conn),
		user: userv1.NewUserAPIClient(conn),
		variable: variablev1.NewVariableAPIClient(conn),
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

// Cluster returns the ClusterAPI client.
func (c *Client) Cluster() clusterv1.ClusterAPIClient {
	return c.cluster
}

// Component returns the ComponentAPI client.
func (c *Client) Component() componentv1.ComponentAPIClient {
	return c.component
}

// Connection returns the ConnectionAPI client.
func (c *Client) Connection() connectionv1.ConnectionAPIClient {
	return c.connection
}

// Deployment returns the DeploymentAPI client.
func (c *Client) Deployment() deploymentv1.DeploymentAPIClient {
	return c.deployment
}

// Environment returns the EnvironmentAPI client.
func (c *Client) Environment() environmentv1.EnvironmentAPIClient {
	return c.environment
}

// Healthcheck returns the HealthcheckAPI client.
func (c *Client) Healthcheck() healthcheckv1.HealthcheckAPIClient {
	return c.healthcheck
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

// Variable returns the VariableAPI client.
func (c *Client) Variable() variablev1.VariableAPIClient {
	return c.variable
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
