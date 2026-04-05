package client_test

import (
	"context"
	"fmt"
	"log"

	"go.admiral.io/sdk/client"
)

func ExampleNew() {
	// Create a client configuration
	cfg := client.Config{
		HostPort:  "localhost:9443",
		AuthToken: "admp_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA000000", // Replace with your token
		ConnectionOptions: client.ConnectionOptions{
			Insecure: true, // For testing only
		},
	}

	// Create client
	c, err := client.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer func() { _ = c.Close() }()

	// Access service clients via accessors
	// Example: userClient := c.User()
	// Example: healthClient := c.Healthcheck()

	fmt.Println("Client created successfully")
}

func ExampleClient_ValidateToken() {
	// Create a client configuration
	cfg := client.Config{
		HostPort:  "localhost:9443",
		AuthToken: "admp_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA000000", // Replace with your token
		ConnectionOptions: client.ConnectionOptions{
			Insecure: true, // For testing only
		},
	}

	// Create client
	c, err := client.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer func() { _ = c.Close() }()

	// Validate token format
	if err := c.ValidateToken(); err != nil {
		log.Printf("Token validation failed: %v", err)
		return
	}

	fmt.Println("Token is valid")
}

func ExampleVersion() {
	// Get the client library version
	version := client.Version()
	fmt.Printf("Client version: %s\n", version)

	// Get the User-Agent string for HTTP/gRPC requests
	userAgent := client.ClientUserAgent()
	fmt.Printf("User-Agent: %s\n", userAgent)
}

func ExampleClient_Version() {
	// Create a client configuration
	cfg := client.Config{
		HostPort:  "localhost:9443",
		AuthToken: "admp_AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA000000", // Replace with your token
		ConnectionOptions: client.ConnectionOptions{
			Insecure: true, // For testing only
		},
	}

	// Create client
	c, err := client.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer func() { _ = c.Close() }()

	// Get version from the client instance
	fmt.Printf("Using client version: %s\n", c.Version())
}
