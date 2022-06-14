package models

// PostModel ...
type PostModel struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// ListPostsModel ...
type ListPostsModel struct {
	Results []PostModel `json:"results"`
	Count   uint32      `json:"count"`
}

// UpdatePostModel ...
type UpdatePostModel struct {
	PostID int    `json:"post_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

