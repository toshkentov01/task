package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	l "github.com/toshkentov01/task/crud_service/pkg/logger"
	crudPb "github.com/toshkentov01/task/crud_service/genproto/crud_service"
)

// CrudService ...
type CrudService struct {
	logger l.Logger
}

// NewCrudService ...
func NewCrudService(log l.Logger) *CrudService {
	return &CrudService{
		logger: log,
	}
}

// ListPosts ...
func (s *CrudService) ListPosts(ctx context.Context, request *crudPb.ListPostsRequest) (*crudPb.ListPostsResponse, error) {

	return &crudPb.ListPostsResponse{}, nil
}

// GetPost ...
func (s *CrudService) GetPost(ctx context.Context, request *crudPb.GetPostRequest) (*crudPb.Post, error) {

	return &crudPb.Post{}, nil
}

// UpdatePost ...
func (s *CrudService) UpdatePost(ctx context.Context, request *crudPb.UpdatePostRequest) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}

// DeletePost ...
func (s *CrudService) DeletePost(ctx context.Context, request *crudPb.DeletePostRequest) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}