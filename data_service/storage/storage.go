package storage

import "github.com/toshkentov01/task/data_service/storage/post"

// Interface ...
type Interface interface {
	Data() post.Repository
}

// storage ...
type storage struct {
	dataRepo post.Repository
}

// NewStorage ...
func NewStorage() Interface {
	return &storage{
		dataRepo: post.NewPostRepo(),
	}
}

func (s storage) Data() post.Repository {
	return s.dataRepo
}
