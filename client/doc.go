// Package client provides Admiral API clients with connection management
// and authentication.
//
// # Quick Start
//
// Create a client and access services:
//
//	c, err := client.New(ctx, client.Config{
//	    AuthToken: "admp_...",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer c.Close()
//
//	// Access services via accessors
//	resp, err := c.User().GetUser(ctx, req)
//
// # Configuration
//
// The Config struct provides options for customizing the client:
//
//   - AuthToken: Required Admiral access token (PAT or SAT)
//   - TokenValidator: Optional custom token validator (defaults to opaque token format checks)
//   - ConnectionOptions: TLS, timeouts, keepalive settings
//   - Logger: Custom logger implementation
//
// # Token Validation
//
// The client validates opaque token format on creation (prefix, length, and
// CRC32 checksum). You can also validate explicitly:
//
//	if err := c.ValidateToken(); err != nil {
//	    // Token has invalid format
//	}
package client
