package model

type Comments struct {
	ID           int
	PostID       int
	Owner        string
	Content      string
	Likes        int
	Dislikes     int
	LikeState    string
	DislikeState string
}
