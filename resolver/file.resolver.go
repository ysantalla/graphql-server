package resolver

import (
	"context"

	"github.com/ysantalla/graphql-server/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *mutationResolver) UploadFile(ctx context.Context, file string) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) UploadFiles(ctx context.Context, files []string) ([]model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) ChangeFile(ctx context.Context, file string, where model.FileWhereUniqueInput) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteFile(ctx context.Context, where model.FileWhereUniqueInput) (*model.File, error) {
	panic("not implemented")
}

func (r *queryResolver) File(ctx context.Context, where model.RoleWhereUniqueInput) (*model.File, error) {
	panic("not implemented")
}
func (r *queryResolver) Files(ctx context.Context, where *model.FileWhereInput, orderBy *model.FileOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]model.File, error) {
	panic("not implemented")
}
