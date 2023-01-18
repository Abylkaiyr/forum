package model

type Post struct {
	ID       int
	Owner    string
	Title    string
	Content  string
	Type     string
	Likes    int
	Dislikes int
}
