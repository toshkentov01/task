package service

import (
	dataPb "github.com/toshkentov01/task/crud_service/genproto/data_service"
	crudPb "github.com/toshkentov01/task/crud_service/genproto/crud_service"
)

// Transformator ...
func Transformator(result []*dataPb.Post) []*crudPb.Post {
	posts := result
	results := []*crudPb.Post{}

	for _, post := range posts {
		crudPost := crudPb.Post{}

		crudPost.Id = post.Id
		crudPost.UserId = post.UserId
		crudPost.Body = post.Body
		crudPost.Title = post.Title

		results = append(results, &crudPost)
	}

	return results
}