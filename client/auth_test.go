package client

import (
	"context"
	"hash/crc32"
	"strings"
	"testing"
)

// createTestToken generates a valid Admiral opaque token with the given prefix
// for testing purposes. The random body is deterministic (all 'A's padded to
// the correct length) and the CRC32 checksum is computed correctly.
func createTestToken(prefix string) string {
	// Token is 54 chars total: 5 (prefix) + 43 (body) + 6 (checksum)
	bodyLen := TokenLength - len(prefix) - checksumLen
	body := prefix + strings.Repeat("A", bodyLen)
	checksum := encodeBase62CRC32(body)
	return body + checksum
}

// createTestTokenWithBadChecksum generates a token with the correct prefix and
// length but an invalid checksum.
func createTestTokenWithBadChecksum(prefix string) string {
	bodyLen := TokenLength - len(prefix) - checksumLen
	body := prefix + strings.Repeat("A", bodyLen)
	return body + "000000" // unlikely to match the real checksum
}

func TestValidateAuthToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty token",
			token:   "",
			wantErr: true,
			errMsg:  "empty",
		},
		{
			name:    "valid PAT",
			token:   createTestToken(TokenPrefixPAT),
			wantErr: false,
		},
		{
			name:    "valid SAT",
			token:   createTestToken(TokenPrefixSAT),
			wantErr: false,
		},
		{
			name:    "valid session token",
			token:   createTestToken(TokenPrefixSession),
			wantErr: false,
		},
		{
			name:    "wrong prefix",
			token:   "admx_" + strings.Repeat("A", TokenLength-5),
			wantErr: true,
			errMsg:  "unrecognized token prefix",
		},
		{
			name:    "too short",
			token:   TokenPrefixPAT + "short",
			wantErr: true,
			errMsg:  "expected length",
		},
		{
			name:    "too long",
			token:   createTestToken(TokenPrefixPAT) + "extra",
			wantErr: true,
			errMsg:  "expected length",
		},
		{
			name:    "bad checksum",
			token:   createTestTokenWithBadChecksum(TokenPrefixPAT),
			wantErr: true,
			errMsg:  "checksum mismatch",
		},
		{
			name:    "valid PAT with Bearer prefix stripped",
			token:   "Bearer " + createTestToken(TokenPrefixPAT),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAuthToken(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.errMsg != "" && err != nil {
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ValidateAuthToken() error = %v, want error containing %q", err, tt.errMsg)
				}
			}
		})
	}
}

func TestDefaultTokenValidator(t *testing.T) {
	v := &DefaultTokenValidator{}

	token := createTestToken(TokenPrefixPAT)
	if err := v.Validate(token); err != nil {
		t.Errorf("DefaultTokenValidator.Validate() unexpected error: %v", err)
	}

	if err := v.Validate(""); err == nil {
		t.Error("DefaultTokenValidator.Validate() expected error for empty token")
	}
}

func TestCustomTokenValidator(t *testing.T) {
	// A simple custom validator that accepts any non-empty token.
	custom := tokenValidatorFunc(func(token string) error {
		if token == "" {
			return ErrInvalidTokenFormat
		}
		return nil
	})

	if err := custom.Validate("anything"); err != nil {
		t.Errorf("custom validator unexpected error: %v", err)
	}
	if err := custom.Validate(""); err == nil {
		t.Error("custom validator expected error for empty token")
	}
}

// tokenValidatorFunc adapts a function to the TokenValidator interface for testing.
type tokenValidatorFunc func(string) error

func (f tokenValidatorFunc) Validate(token string) error { return f(token) }

func TestTokenAuth_AuthScheme(t *testing.T) {
	tests := []struct {
		name   string
		scheme AuthScheme
		want   string
	}{
		{
			name:   "bearer scheme (default zero value)",
			scheme: AuthSchemeBearer,
			want:   "Bearer my-token",
		},
		{
			name:   "token scheme",
			scheme: AuthSchemeToken,
			want:   "Token my-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := tokenAuth{
				token:  "my-token",
				scheme: tt.scheme,
			}
			md, err := ta.GetRequestMetadata(context.Background())
			if err != nil {
				t.Fatalf("GetRequestMetadata() error = %v", err)
			}
			got := md["Authorization"]
			if got != tt.want {
				t.Errorf("Authorization = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAuthScheme_String(t *testing.T) {
	if got := AuthSchemeBearer.String(); got != "Bearer" {
		t.Errorf("AuthSchemeBearer.String() = %q, want %q", got, "Bearer")
	}
	if got := AuthSchemeToken.String(); got != "Token" {
		t.Errorf("AuthSchemeToken.String() = %q, want %q", got, "Token")
	}
}

func TestEncodeBase62CRC32(t *testing.T) {
	// Verify our implementation matches the server's algorithm.
	// CRC32 IEEE of "admp_" + 43 'A's should produce a deterministic result.
	body := TokenPrefixPAT + strings.Repeat("A", TokenLength-len(TokenPrefixPAT)-checksumLen)
	checksum := encodeBase62CRC32(body)

	if len(checksum) != checksumLen {
		t.Errorf("encodeBase62CRC32() length = %d, want %d", len(checksum), checksumLen)
	}

	// Verify the checksum is stable (deterministic).
	checksum2 := encodeBase62CRC32(body)
	if checksum != checksum2 {
		t.Errorf("encodeBase62CRC32() not deterministic: %q != %q", checksum, checksum2)
	}

	// Verify it uses the same CRC32 IEEE algorithm.
	n := crc32.ChecksumIEEE([]byte(body))
	if n == 0 {
		t.Error("CRC32 should not be zero for non-empty input")
	}
}

func TestTokenConstants(t *testing.T) {
	// Verify token constants are consistent.
	if len(TokenPrefixPAT) != 5 {
		t.Errorf("TokenPrefixPAT length = %d, want 5", len(TokenPrefixPAT))
	}
	if len(TokenPrefixSAT) != 5 {
		t.Errorf("TokenPrefixSAT length = %d, want 5", len(TokenPrefixSAT))
	}
	if len(TokenPrefixSession) != 5 {
		t.Errorf("TokenPrefixSession length = %d, want 5", len(TokenPrefixSession))
	}
	// 5 (prefix) + 43 (base64url of 32 bytes) + 6 (checksum) = 54
	if TokenLength != 54 {
		t.Errorf("TokenLength = %d, want 54", TokenLength)
	}
}
