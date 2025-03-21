package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"
	"fmt"

	"blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	"blogapi.miyamo.today/federator/internal/infra/fw/gqlgen"
)

// Noop is the resolver for the noop field.
func (r *mutationResolver) Noop(ctx context.Context, input *model.NoopInput) (*model.NoopPayload, error) {
	panic(fmt.Errorf("not implemented: Noop - noop"))
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
