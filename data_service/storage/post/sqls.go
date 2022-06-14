package post

const (
	// ListPostsSQL ...
	ListPostsSQL = `
		SELECT
			id,
			user_id,
			title,
			body
		FROM posts
		WHERE deleted_at is NULL
		LIMIT $1
		OFFSET $2
	`

	// GetPostSQL ...
	GetPostSQL = `
		SELECT
			id,
			user_id,
			title,
			body
		FROM posts
		WHERE id = $1 AND deleted_at IS NULL
	`

	// UpdatePostSQL ...
	UpdatePostSQL = `
		UPDATE posts 
		SET title = COALESCE($1, title),
			body = COALESCE($2, body),
			updated_at = NOW()
		WHERE id = $3
	`

	// DeletePostSQL ...
	DeletePostSQL = `
		UPDATE posts SET deleted_at = NOW()
		WHERE id = $1
	`

	// CountPostsSQL ...
	CountPostsSQL = `
		SELECT COUNT(*) FROM posts where deleted_at IS NULL
	`

	// CheckForOwnershipSQL ...
	CheckForOwnershipSQL = `
		SELECT
			EXISTS(SELECT * FROM posts WHERE id = $1 AND user_id = $2)
	`
)