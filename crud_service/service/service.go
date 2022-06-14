package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	crudPb "github.com/toshkentov01/task/crud_service/genproto/crud_service"
	dataPb "github.com/toshkentov01/task/crud_service/genproto/data_service"
	l "github.com/toshkentov01/task/crud_service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client "github.com/toshkentov01/task/crud_service/service/grpc_client"
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
	result, err := client.DataService().ListPosts(ctx, &dataPb.ListPostsRequest{
		Limit: request.Limit,
		Page:  request.Page,
	})

	results := Transformator(result.Results)

	return &crudPb.ListPostsResponse{
		Results: results,
		Count:   result.Count,
	}, errorHandler(s.logger, err, "Failed to list posts: ")
}

// GetPost ...
func (s *CrudService) GetPost(ctx context.Context, request *crudPb.GetPostRequest) (*crudPb.Post, error) {
	result, err := client.DataService().GetPost(ctx, &dataPb.GetPostRequest{
		PostId: request.PostId,
	})

	return &crudPb.Post{
		Id:     result.Id,
		UserId: result.UserId,
		Title:  result.Title,
		Body:   result.Body,
	}, errorHandler(s.logger, err, "Failed to get post: ")
}

// UpdatePost ...
func (s *CrudService) UpdatePost(ctx context.Context, request *crudPb.UpdatePostRequest) (*empty.Empty, error) {
	_, err := client.DataService().UpdatePost(ctx, &dataPb.UpdatePostRequest{
		PostId: request.PostId,
		Title:  request.Title,
		Body:   request.Body,
	})

	return &empty.Empty{}, errorHandler(s.logger, err, "Failed to update post: ")
}

// DeletePost ...
func (s *CrudService) DeletePost(ctx context.Context, request *crudPb.DeletePostRequest) (*empty.Empty, error) {
	_, err := client.DataService().DeletePost(ctx, &dataPb.DeletePostRequest{
		PostId: request.PostId,
		UserId: request.UserId,
	})

	return &empty.Empty{}, errorHandler(s.logger, err, "Failed to delete post: ")
}

// CheckForOwnership ...
func (s *CrudService) CheckForOwnership(ctx context.Context, request *crudPb.CheckForOwnershipRequest) (*crudPb.CheckForOwnershipResponse, error) {
	result, err := client.DataService().CheckForOwnership(ctx, &dataPb.CheckForOwnershipRequest{
		UserId: request.UserId,
		PostId: request.PostId,
	})

	if err != nil {
		s.logger.Error("Error while checking post ownership, error: ")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	return &crudPb.CheckForOwnershipResponse{
		Owner: result.Owner,
	}, nil
}
