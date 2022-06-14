package post

import (
	"github.com/toshkentov01/task/data_service/pkg/open_api/posts"

	dataPb "github.com/toshkentov01/task/data_service/genproto/data_service"
)

// dbPostModel ...
type dbPostModel struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Title  string `db:"title"`
	Body   string `db:"body"`
}

// newDbPostModel ...
func newDbPostModel(post *posts.Data) *dbPostModel {
	return &dbPostModel{
		ID:     post.ID,
		UserID: post.UserID,
		Title:  post.Title,
		Body:   post.Body,
	}
}

// toModel ...
func (d dbPostModel) toModel() *dataPb.Post {
	return &dataPb.Post{
		Id:     int64(d.ID),
		UserId: int64(d.UserID),
		Title:  d.Title,
		Body:   d.Body,
	}
}

type DbData struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Title  string `db:"title"`
	Body   string `db:"body"`
}
