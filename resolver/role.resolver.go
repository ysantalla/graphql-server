package resolver

import (
	"context"

	"github.com/ysantalla/graphql-server/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *mutationResolver) CreateRole(ctx context.Context, data model.RoleCreateInput) (*model.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateRole(ctx context.Context, data model.RoleUpdateInput, where model.RoleWhereUniqueInput) (*model.Role, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteRole(ctx context.Context, where model.RoleWhereUniqueInput) (*model.Role, error) {
	panic("not implemented")
}

func (r *queryResolver) Role(ctx context.Context, where model.RoleWhereUniqueInput) (*model.Role, error) {
	panic("not implemented")
}
func (r *queryResolver) Roles(ctx context.Context, where *model.RoleWhereInput, orderBy *model.RoleOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]model.Role, error) {
	panic("not implemented")
}
