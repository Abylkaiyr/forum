package utils

type User struct {
	ID           int
	UserName     string
	UserEmail    string
	UserPassword string
}

type Post struct {
	ID           int
	PostOwner    string
	PostTitle    string
	PostContent  string
	PostCategory []string
	PostLikes    int
	PostDisLikes int
}

type Comments struct {
	ID              int
	PostID          int
	CommentOwner    string
	CommentContent  string
	CommentLikes    int
	CommentDislikes int
}
