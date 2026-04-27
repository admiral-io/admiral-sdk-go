# admiral-sdk-go

Go client library for the Admiral API.

## Installation

```bash
go get go.admiral.io/sdk
```

## Usage

```go
package main

import (
	"context"
	"log"

	"go.admiral.io/sdk/client"
)

func main() {
	ctx := context.Background()

	// Create client
	c, err := client.New(ctx, client.Config{
		HostPort:  "api.admiral.io:443",
		AuthToken: "your-token-here",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Access services via accessors
	// c.Application().MethodName(ctx, req)
	// c.Authentication().MethodName(ctx, req)
	// c.ChangeSet().MethodName(ctx, req)
	// c.Cluster().MethodName(ctx, req)
	// c.Component().MethodName(ctx, req)
	// c.Credential().MethodName(ctx, req)
	// c.Deployment().MethodName(ctx, req)
	// c.Environment().MethodName(ctx, req)
	// c.Healthcheck().MethodName(ctx, req)
	// c.Module().MethodName(ctx, req)
	// c.Runner().MethodName(ctx, req)
	// c.Source().MethodName(ctx, req)
	// c.State().MethodName(ctx, req)
	// c.User().MethodName(ctx, req)
	// c.Variable().MethodName(ctx, req)
}
```

## Available Services

| Service | Accessor | Import |
|---------|----------|--------|
| ApplicationAPI | `Application()` | `go.admiral.io/sdk/proto/admiral/application/v1` |
| AuthenticationAPI | `Authentication()` | `go.admiral.io/sdk/proto/admiral/authentication/v1` |
| ChangeSetAPI | `ChangeSet()` | `go.admiral.io/sdk/proto/admiral/changeset/v1` |
| ClusterAPI | `Cluster()` | `go.admiral.io/sdk/proto/admiral/cluster/v1` |
| ComponentAPI | `Component()` | `go.admiral.io/sdk/proto/admiral/component/v1` |
| CredentialAPI | `Credential()` | `go.admiral.io/sdk/proto/admiral/credential/v1` |
| DeploymentAPI | `Deployment()` | `go.admiral.io/sdk/proto/admiral/deployment/v1` |
| EnvironmentAPI | `Environment()` | `go.admiral.io/sdk/proto/admiral/environment/v1` |
| HealthcheckAPI | `Healthcheck()` | `go.admiral.io/sdk/proto/admiral/healthcheck/v1` |
| ModuleAPI | `Module()` | `go.admiral.io/sdk/proto/admiral/module/v1` |
| RunnerAPI | `Runner()` | `go.admiral.io/sdk/proto/admiral/runner/v1` |
| SourceAPI | `Source()` | `go.admiral.io/sdk/proto/admiral/source/v1` |
| StateAPI | `State()` | `go.admiral.io/sdk/proto/admiral/state/v1` |
| UserAPI | `User()` | `go.admiral.io/sdk/proto/admiral/user/v1` |
| VariableAPI | `Variable()` | `go.admiral.io/sdk/proto/admiral/variable/v1` |

## Example

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.admiral.io/sdk/client"
	applicationv1 "go.admiral.io/sdk/proto/admiral/application/v1"
)

func main() {
	ctx := context.Background()

	c, err := client.New(ctx, client.Config{
		HostPort:  "api.admiral.io:443",
		AuthToken: os.Getenv("ADMIRAL_TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Validate token before making requests
	if err := c.ValidateToken(); err != nil {
		log.Fatal("Invalid token:", err)
	}

	// Call a service method
	resp, err := c.Application().ListMethod(ctx, &applicationv1.ListMethodRequest{})
	if err != nil {
		log.Fatal("Request failed:", err)
	}
	fmt.Printf("Response: %+v\n", resp)
}
```

## Configuration

```go
import (
	"crypto/tls"
	"time"

	"go.admiral.io/sdk/client"
)

cfg := client.Config{
	// Required: Server address
	HostPort: "api.admiral.io:443",

	// Required: Admiral access token (PAT or SAT)
	AuthToken: os.Getenv("ADMIRAL_TOKEN"),

	// Optional: Connection options
	ConnectionOptions: client.ConnectionOptions{
		// Use insecure connection (no TLS) - default: false
		Insecure: false,

		// Connection timeout - default: 10s
		DialTimeout: 10 * time.Second,

		// Custom TLS configuration
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},

		// Additional gRPC dial options
		DialOptions: []grpc.DialOption{},
	},

	// Optional: Custom logger (default: no-op logger)
	Logger: client.NewDefaultLogger(),
}

c, err := client.New(ctx, cfg)
```

## Token Validation

The client validates token format on creation (prefix, length, CRC32 checksum).
You can also validate explicitly:

```go
if err := c.ValidateToken(); err != nil {
	log.Fatal("Invalid token:", err)
}
```

A custom `TokenValidator` can be provided to override the default validation:

```go
cfg := client.Config{
	AuthToken:      os.Getenv("ADMIRAL_TOKEN"),
	TokenValidator: &myCustomValidator{},
}
```

## Version Information

```go
// Get library version
fmt.Println("Version:", client.Version())

// Get user agent string (useful for debugging)
fmt.Println("User-Agent:", client.ClientUserAgent())
```

## Requirements

- Go 1.26 or later

## License

Apache-2.0 License - see [LICENSE](LICENSE.txt) for details.
