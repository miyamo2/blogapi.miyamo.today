package provider

import (
	"blogapi.miyamo.today/core/echo/middlewares"
	"blogapi.miyamo.today/federator/internal/infra/aws/cognito"
	"github.com/google/wire"
	"os"
)

// compatibility check
var _ middlewares.Verifier = (*cognito.Verifier)(nil)

func Verifier() *cognito.Verifier {
	return cognito.NewVerifier(
		os.Getenv("AWS_REGION"),
		os.Getenv("COGNITO_APP_CLIENT_ID"),
		os.Getenv("COGNITO_USER_POOL_ID"),
	)
}

var VeryfierSet = wire.NewSet(
	Verifier,
	wire.Bind(new(middlewares.Verifier), new(*cognito.Verifier)))
