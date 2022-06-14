package post

import (
	dataPb "github.com/toshkentov01/task/data_service/genproto/data_service"
)

// Repository ...
type Repository interface {
	Reader
	Writer
}

// Reader ...
type Reader interface {
	GetPost(postID int) (*dataPb.Post, error)
	ListPosts(limit, page uint32) (*dataPb.ListPostsResponse, error)
}

// Writer ...
type Writer interface {
	UpdatePost(postID int, title, body string) error
	DeletePost(postID int) error
	CheckForOwnership(postID, userID int) (*dataPb.CheckForOwnershipResponse, error)
	InsertPosts() error
}
