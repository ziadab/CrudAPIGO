package requests

var PostCreate struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
