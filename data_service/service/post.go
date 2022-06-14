package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	l "github.com/toshkentov01/task/data_service/pkg/logger"
	"github.com/toshkentov01/task/data_service/storage"

	dataPb "github.com/toshkentov01/task/data_service/genproto/data_service"
)

// DataService ...
type DataService struct {
	storage storage.Interface
	logger  l.Logger
}

// NewDataService ...
func NewDataService(log l.Logger) *DataService {
	return &DataService{
		storage: storage.NewStorage(),
		logger:  log,
	}
}

// ListPosts ...
func (u *DataService) ListPosts(ctx context.Context, request *dataPb.ListPostsRequest) (*dataPb.ListPostsResponse, error) {
	result, err := u.storage.Data().ListPosts(request.Limit, request.Page)

	return result, errorHandler(u.logger, err, "Error while getting posts")
}

// GetPost ...
func (u *DataService) GetPost(ctx context.Context, request *dataPb.GetPostRequest) (*dataPb.Post, error) {
	result, err := u.storage.Data().GetPost(int(request.PostId))

	return result, errorHandler(u.logger, err, "Failed to get post")
}

// UpdatePost ...
func (u *DataService) UpdatePost(ctx context.Context, request *dataPb.UpdatePostRequest) (*empty.Empty, error) {
	err := u.storage.Data().UpdatePost(int(request.PostId), request.Title, request.Body)

	return &empty.Empty{}, errorHandler(u.logger, err, "Failed to update post")
}

// DeletePost ...
func (u *DataService) DeletePost(ctx context.Context, request *dataPb.DeletePostRequest) (*empty.Empty, error) {
	err := u.storage.Data().DeletePost(int(request.PostId))

	return &empty.Empty{}, errorHandler(u.logger, err, "Failed to delete post")
}

// CheckForOwnership ...
func (u *DataService) CheckForOwnership(ctx context.Context, request *dataPb.CheckForOwnershipRequest) (*dataPb.CheckForOwnershipResponse, error) {
	result, err := u.storage.Data().CheckForOwnership(int(request.PostId), int(request.UserId))

	return result, errorHandler(u.logger, err, "Failed to check fo ownership")
}
