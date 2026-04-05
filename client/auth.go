package client

import (
	"errors"
	"fmt"
	"hash/crc32"
	"strings"
)

const (
	// TokenPrefixPAT is the prefix for Personal Access Tokens.
	TokenPrefixPAT = "admp_"
	// TokenPrefixSAT is the prefix for Service/Agent Tokens.
	TokenPrefixSAT = "adms_"
	// TokenPrefixSession is the prefix for Session Tokens.
	TokenPrefixSession = "adme_"

	// TokenLength is the total length of an Admiral opaque token.
	TokenLength = 54

	// checksumLen is the length of the base62-encoded CRC32 checksum suffix.
	checksumLen = 6
)

// base62 alphabet for CRC32 checksum encoding.
const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var (
	ErrInvalidTokenFormat = errors.New("invalid token format")

	validPrefixes = []string{TokenPrefixPAT, TokenPrefixSAT, TokenPrefixSession}
)

// TokenValidator validates an auth token before use.
// Implementations can add custom validation logic beyond the default
// Admiral opaque token format checks.
type TokenValidator interface {
	Validate(token string) error
}

// DefaultTokenValidator validates Admiral opaque tokens by checking the
// prefix, length, and CRC32 checksum.
type DefaultTokenValidator struct{}

// Validate checks that the token is a well-formed Admiral opaque token.
func (v *DefaultTokenValidator) Validate(token string) error {
	return ValidateAuthToken(token)
}

// ValidateAuthToken validates the format of an Admiral opaque token.
// It checks the prefix, length, and CRC32 checksum without requiring
// a network call.
func ValidateAuthToken(token string) error {
	if token == "" {
		return errors.New("auth token is empty")
	}

	// Strip Bearer prefix if present.
	actual := strings.TrimPrefix(token, "Bearer ")

	if len(actual) != TokenLength {
		return fmt.Errorf("%w: expected length %d, got %d", ErrInvalidTokenFormat, TokenLength, len(actual))
	}

	if !hasValidPrefix(actual) {
		return fmt.Errorf("%w: unrecognized token prefix", ErrInvalidTokenFormat)
	}

	if !validateChecksum(actual) {
		return fmt.Errorf("%w: checksum mismatch", ErrInvalidTokenFormat)
	}

	return nil
}

// hasValidPrefix reports whether the token starts with a known Admiral prefix.
func hasValidPrefix(token string) bool {
	for _, p := range validPrefixes {
		if strings.HasPrefix(token, p) {
			return true
		}
	}
	return false
}

// validateChecksum verifies the CRC32 checksum suffix of an opaque token.
func validateChecksum(token string) bool {
	if len(token) <= checksumLen {
		return false
	}
	body := token[:len(token)-checksumLen]
	expected := encodeBase62CRC32(body)
	return token[len(token)-checksumLen:] == expected
}

// encodeBase62CRC32 computes CRC32 of s and encodes it as a fixed-length
// base62 string (zero-padded to checksumLen characters).
func encodeBase62CRC32(s string) string {
	n := crc32.ChecksumIEEE([]byte(s))
	buf := make([]byte, checksumLen)
	for i := checksumLen - 1; i >= 0; i-- {
		buf[i] = base62[n%62]
		n /= 62
	}
	return string(buf)
}
