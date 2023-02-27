package model

import "time"

type Post struct {
	ID          int
	Owner       string
	Title       string
	Content     string
	Type        string
	Likes       int
	Dislikes    int
	CreatedTime time.Time
	Timer       string
	Image       []byte
	FilePath    string
}

type Post1 struct {
	ID          int
	Owner       string
	Title       string
	Content     string
	Type        []string
	Likes       int
	Dislikes    int
	CreatedTime time.Time
	Timer       string
	Image       []byte
	FilePath    string
}
