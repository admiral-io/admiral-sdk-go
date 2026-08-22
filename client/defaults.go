package client

import "time"

// DefaultHostPort is the endpoint used when Config.HostPort is not set.
// Declared by this collection's protorepo.yaml (default_host_port).
const DefaultHostPort = "api.admiral.io:443"

// DefaultDialTimeout is the default timeout for establishing a gRPC connection.
const DefaultDialTimeout = 30 * time.Second

// DefaultKeepAliveTime is the default interval for sending keepalive pings.
const DefaultKeepAliveTime = 30 * time.Second

// DefaultKeepAliveTimeout is the default timeout for keepalive ping responses.
const DefaultKeepAliveTimeout = 90 * time.Second
