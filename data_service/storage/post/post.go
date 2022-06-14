package post

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/toshkentov01/task/data_service/pkg/errs"
	"github.com/toshkentov01/task/data_service/pkg/open_api/posts"
	"github.com/toshkentov01/task/data_service/pkg/platform/postgres"
	"github.com/toshkentov01/task/data_service/pkg/utils"

	dataPb "github.com/toshkentov01/task/data_service/genproto/data_service"
)

// postRepo ...
type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo() Repository {
	return &postRepo{
		db: postgres.DB(),
	}
}

// GetPost ...
func (pr *postRepo) GetPost(postID int) (*dataPb.Post, error) {
	var post dbPostModel

	if err := pr.db.Get(&post, GetPostSQL, postID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrNotFound
		}

		return nil, errs.ErrInternal
	}

	return post.toModel(), nil
}

// UpdatePost ...
func (pr *postRepo) UpdatePost(postID int, title, body string) error {
	converterTitle, convertedBody := utils.StringToNullString(title), utils.StringToNullString(body)

	result, err := pr.db.Exec(UpdatePostSQL, converterTitle, convertedBody, postID)

	if err != nil {
		return errs.ErrInternal
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errs.ErrNotFound
	}

	return nil
}

// DeletePost ...
func (pr *postRepo) DeletePost(postID int) error {
	result, err := pr.db.Exec(DeletePostSQL, postID)

	if err != nil {
		return errs.ErrInternal
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errs.ErrNotFound
	}

	return nil
}

// ListPosts ...
func (pr *postRepo) ListPosts(limit, page uint32) (*dataPb.ListPostsResponse, error) {
	var count uint32

	results := []*dataPb.Post{}
	offset := (page - 1) * limit

	if err := pr.db.QueryRow(CountPostsSQL).Scan(&count); err != nil {
		return nil, errs.ErrInternal
	}

	rows, err := pr.db.Queryx(ListPostsSQL, limit, offset)
	if err != nil {
		return nil, errs.ErrInternal
	}

	defer rows.Close()

	for rows.Next() {
		var post dbPostModel

		if err := rows.StructScan(&post); err != nil {
			return nil, errs.ErrInternal
		}

		results = append(results, post.toModel())
	}

	return &dataPb.ListPostsResponse{
		Results: results,
		Count:   count,
	}, nil
}

// CheckForOwnership ...
func (pr *postRepo) CheckForOwnership(postID, userID int) (*dataPb.CheckForOwnershipResponse, error) {
	var isOwner bool

	if err := pr.db.QueryRow(CheckForOwnershipSQL, postID, userID).Scan(&isOwner); err != nil {
		return nil, errs.ErrInternal
	}

	return &dataPb.CheckForOwnershipResponse{
		Owner: isOwner,
	}, nil
}

// InsertPosts - this is function and independent from postRepo.
// Because this function just inserts to database posts which we got from open api.
func (pr *postRepo) InsertPosts() error {
	// In this function I inserted 1000 posts at once (in one connection)
	// because it is too costy to connect to postgres in order to insert for each post

	datas := posts.GetPosts()
	postDatas := []posts.Data{}

	postCount := 0

	if err := pr.db.QueryRow(CountPostsSQL).Scan(&postCount); err != nil {
		return errs.ErrInternal
	}

	if postCount == 0 {

		for _, data := range datas {
			postDatas = append(postDatas, data.Data...)
		}

		_, err := pr.db.NamedExec(`INSERT INTO posts(id, user_id, title, body)
			VALUES (:id, :user_id, :title, :body)`, postDatas,
		)

		if err != nil {
			return errs.ErrInternal
		}
	}

	return nil
}
