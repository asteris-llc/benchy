package rpc

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// TokenAuthenticator is for use on a client to provide a token
type TokenAuthenticator string

// NewTokenAuthenticator creates a new authentic
func NewTokenAuthenticator(token string) TokenAuthenticator {
	return TokenAuthenticator(token)
}

// GetRequestMetadata gets header data for a given request
func (a TokenAuthenticator) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	fmt.Println(uri)
	return map[string]string{
		"authorization": fmt.Sprintf("BEARER %s", a),
	}, nil
}

// RequireTransportSecurity should be true, but isn't
func (TokenAuthenticator) RequireTransportSecurity() bool { return false }

// TokenVerifier verifies tokens against a collection of tokens
type TokenVerifier struct {
	tokenToProject map[string]string
}

// NewTokenVerifier creates a verifier from a map of tokens to their projects
func NewTokenVerifier(tokens map[string]string) *TokenVerifier {
	return &TokenVerifier{tokens}
}

// VerifyForProject verifies that a token is valid for a given project
func (t *TokenVerifier) VerifyForProject(token, project string) error {
	if should, ok := t.tokenToProject[token]; ok && should == project {
		return nil
	}

	return errors.New("invalid token for project")
}

// VerifyForProjectInContext does VerifyForProject in a context
func (t *TokenVerifier) VerifyForProjectInContext(ctx context.Context, project string) error {
	token, _ := TokenFromContext(ctx)
	return t.VerifyForProject(token, project)
}

// TokenFromContext retrieves a token from a context
func TokenFromContext(ctx context.Context) (token string, ok bool) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return "", false
	}

	tokens, ok := md["authorization"]
	if !ok {
		return "", false
	}

	return strings.TrimLeft(tokens[0], "BEARER "), true
}
