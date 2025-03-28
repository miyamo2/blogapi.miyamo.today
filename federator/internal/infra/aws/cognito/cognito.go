package cognito

import (
	"context"
	"fmt"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"time"
)

var (
	ErrTokenUseUnmatched = errors.New("token_use unmatched")
)

// Verifier implements Verifier with cognito
type Verifier struct {
	jwksURL  string
	clientID string
	issuer   string
}

// Verify verifies the token
func (v *Verifier) Verify(ctx context.Context, tokenStr string) (jwt.Token, error) {
	// OPTIMIZE: enable cache
	keySet, err := jwk.Fetch(ctx, v.jwksURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch jwk")
	}
	token, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
		jwt.WithAudience(v.clientID),
		jwt.WithIssuer(v.issuer),
		withTokenUse(),
		jwt.WithClock(&clock{}),
	)
	return token, nil
}

// NewVerifier creates a new verifier
func NewVerifier(region, clientID, userPoolID string) *Verifier {
	return &Verifier{
		jwksURL:  fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID),
		clientID: clientID,
		issuer:   fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", region, userPoolID),
	}
}

// allowedTokenUse is a map of allowed values for token_use claim
var allowedTokenUse = map[string]bool{
	"id":     true,
	"access": true,
}

// validateTokenUse validates token_use claim
func validateTokenUse(_ context.Context, t jwt.Token) error {
	var tokenUse string
	if err := t.Get("token_use", &tokenUse); err != nil {
		return err
	}
	if !allowedTokenUse[tokenUse] {
		return errors.WithDetail(ErrTokenUseUnmatched, fmt.Sprintf("token_use: %s", tokenUse))
	}
	return nil
}

// withTokenUse enables token_use validation
func withTokenUse() jwt.ValidateOption {
	return jwt.WithValidator(jwt.ValidatorFunc(validateTokenUse))
}

// clock implements jwt.Clock
type clock struct{}

// Now returns the current time
func (c *clock) Now() time.Time {
	now := synchro.Now[tz.UTC]()
	return now.StdTime()
}
