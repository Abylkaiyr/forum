package store

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

func (r *UserRepository) CreatePostByUsername(s *model.Post, buf bytes.Buffer, header *multipart.FileHeader) error {
	timer := s.CreatedTime.Format(time.Kitchen)
	s.FilePath = "/uploads/" + "temp_" + header.Filename
	statement, _ := r.store.db.Prepare("INSERT INTO post (owner,title, content, type, createdTime, timer, image, filepath) VALUES (?,?,?,?,?,?,?, ?)")
	_, err := statement.Exec(s.Owner, s.Title, s.Content, s.Type, s.CreatedTime, timer, buf.Bytes(), s.FilePath)
	return err
}

// Get all posts if available

func (r *UserRepository) GetAllPosts() ([]model.Post1, error) {
	// Finding this user from users table and returning it
	var post model.Post
	var posts []model.Post1
	var post1 model.Post1

	rows, err := r.store.db.Query("SELECT * FROM post")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&post.ID, &post.Owner, &post.Title, &post.Content, &post.Type, &post.Likes, &post.Dislikes, &post.CreatedTime, &post.Timer, &post.Image, &post.FilePath)
		// Created new post model because in category,
		// we can not save []string in db
		post1.ID = post.ID
		post1.Owner = post.Owner
		post1.Title = post.Title
		post1.Content = post.Content
		post1.Type = strings.Split(post.Type, ",")
		post1.Likes = post.Likes
		post1.Dislikes = post.Dislikes
		post1.CreatedTime = post.CreatedTime
		post1.Timer = post.Timer
		post1.FilePath = post.FilePath
		posts = append(posts, post1)
	}
	// posts = append(posts, post)
	return posts, nil
}

func (r *UserRepository) GetPostByCategory(category string) ([]model.Post1, error) {
	// Finding this user from users table and returning it
	var post model.Post
	var posts []model.Post1
	var post1 model.Post1

	query := "SELECT * FROM post WHERE type like ?"
	rows, err := r.store.db.Query(query, "%"+category+"%")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(&post.ID, &post.Owner, &post.Title, &post.Content, &post.Type, &post.Likes, &post.Dislikes, &post.CreatedTime, &post.Timer, &post.Image, &post.FilePath)
		// Created new post model because in category,
		// we can not save []string in db
		post1.ID = post.ID
		post1.Owner = post.Owner
		post1.Title = post.Title
		post1.Content = post.Content
		post1.Type = strings.Split(post.Type, ",")
		post1.Likes = post.Likes
		post1.Dislikes = post.Dislikes
		post1.CreatedTime = post.CreatedTime
		post1.Timer = post.Timer
		post1.Image = post.Image
		post1.FilePath = post.FilePath

		posts = append(posts, post1)

	}

	// posts = append(posts, post)
	return posts, err
}

func (r *UserRepository) GetPostByPostID(id int) (model.Post1, error) {
	// Finding this user from users table and returning it
	var post model.Post
	//var posts []model.Post1
	var post1 model.Post1

	query := "SELECT * FROM post WHERE id = ?"
	rows, err := r.store.db.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(&post.ID, &post.Owner, &post.Title, &post.Content, &post.Type, &post.Likes, &post.Dislikes, &post.CreatedTime, &post.Timer, &post.Image, &post.FilePath)
		// Created new post model because in category,
		// we can not save []string in db
		post1.ID = post.ID
		post1.Owner = post.Owner
		post1.Title = post.Title
		post1.Content = post.Content
		post1.Type = strings.Split(post.Type, ",")
		post1.Likes = post.Likes
		post1.Dislikes = post.Dislikes
		post1.CreatedTime = post.CreatedTime
		post1.Timer = post.Timer
		post1.Image = post.Image
		post1.FilePath = post.FilePath

	}

	// posts = append(posts, post)

	return post1, err
}

func (r *UserRepository) InsertImageToPost(buf bytes.Buffer) error {
	_, err := r.store.db.Exec("INSERT INTO post (image) VALUES (?)", buf.Bytes())
	if err != nil {
		fmt.Println("error inserting image:", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetPostDisLikeById(id int) (int, error) {
	likes := 0
	query1 := "select dislikes from post where id = $1"
	rows1 := r.store.db.QueryRow(query1, id)
	if err := rows1.Scan(&likes); err != nil {
		return 0, err
	}

	return likes, nil
}
