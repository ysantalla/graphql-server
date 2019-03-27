package resolver

import (
	"context"

	"github.com/ysantalla/graphql-server/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserCreateInput) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, data model.UserUpdateInput, where model.UserWhereUniqueInput) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, where model.UserWhereUniqueInput) (*model.User, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, where model.UserWhereUniqueInput) (*model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context, where *model.UserWhereInput, orderBy *model.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]model.User, error) {
	panic("not implemented")
}
