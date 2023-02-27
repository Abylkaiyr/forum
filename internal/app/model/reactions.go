package model

type Reactions struct {
	ID            int
	PostID        int
	PostLiker     string
	PostDisLiker  string
	Likes         int
	Dislikes      int
	TotalLikes    int
	TotalDislikes int
	LikeState     string
	DisLikeState  string
}
