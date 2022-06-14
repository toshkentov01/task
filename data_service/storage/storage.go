package storage

import "github.com/toshkentov01/task/data_service/storage/post"

// Interface ...
type Interface interface {
	Data() post.PostRepository
}

// storage ...
type storage struct {
	dataRepo post.PostRepository
}

// NewStorage ...
func NewStorage() Interface {
	return &storage{
		dataRepo: post.NewPostRepo(),
	}
}

func (s storage) Data() post.PostRepository {
	return s.dataRepo
}