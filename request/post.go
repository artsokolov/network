package request

type CreatePostRequest struct {
	Content string `json:"content" binding:"required,url"`
}

type Post struct {
	Id string `uri:"id" binding:"required"`
}
