package resolver

import (
	"context"

	"github.com/ysantalla/graphql-server/gql"
	"github.com/ysantalla/graphql-server/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

func (r *Resolver) User() gql.UserResolver {
	return &userResolver{r}
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.ID.Hex(), nil
}
