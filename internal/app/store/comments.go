package store

import (
	"github.com/Abylkaiyr/forum/internal/app/model"
)

func (r *UserRepository) GetAllComments(post model.Post1) ([]model.Comments, error) {
	c := model.Comments{}
	var comments []model.Comments
	query := "SELECT * FROM comment WHERE postID = $1"
	rows, err := r.store.db.Query(query, post.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(&c.ID, &c.PostID, &c.Owner, &c.Content, &c.Likes, &c.Dislikes, &c.LikeState, &c.DislikeState)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func (r *UserRepository) CreateComment(post model.Post1, c *model.Comments) error {
	statement, _ := r.store.db.Prepare("INSERT INTO comment (postID, owner,content, likes, dislikes, likeState, dislikeState) VALUES (?,?,?,?,?,?,?)")
	_, err := statement.Exec(c.PostID, c.Owner, c.Content, c.Likes, c.Dislikes, c.LikeState, c.DislikeState)
	return err
}
