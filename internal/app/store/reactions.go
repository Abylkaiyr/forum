package store

import (
	"log"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

// L I K E S

func (r *UserRepository) GetReactions(postID int) (*model.Reactions, error) {
	// Finding this user from users table and returning it
	reactions := &model.Reactions{}

	query := "SELECT * FROM reactions WHERE postID = ?"
	rows, err := r.store.db.Query(query, postID)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&reactions.ID, &reactions.PostID, &reactions.PostLiker, &reactions.PostDisLiker, &reactions.Likes, &reactions.Dislikes, &reactions.TotalLikes, &reactions.TotalDislikes, &reactions.LikeState, &reactions.DisLikeState)
		if err != nil {
			return nil, err
		}
	}

	return reactions, nil
}

func (r *UserRepository) CreateReactionLike(reactions *model.Reactions, user model.User, post *model.Post1, postLiker string, postDisliker string) error {

	statement, _ := r.store.db.Prepare("INSERT INTO reactions (postID,postLiker, postDisLiker,likes, dislikes, totalLikes, totalDislikes, likeState) VALUES (?,?,?,?,?,?,?,?)")
	_, err := statement.Exec(post.ID, postLiker, postDisliker, reactions.Likes, reactions.Dislikes, reactions.TotalLikes, reactions.TotalDislikes, reactions.LikeState)
	return err
}

func (r *UserRepository) UpdatePostLikesInitial(reactions *model.Reactions, user model.User) error {
	query := "UPDATE reactions SET postLiker = ?, postDisLiker = ?, totalLikes = ?, likeState = ?, dislikeState =? WHERE postID = ?"
	_, err := r.store.db.Exec(query, reactions.PostLiker, reactions.PostDisLiker, reactions.TotalLikes, reactions.LikeState, reactions.DisLikeState, reactions.PostID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *UserRepository) UpdatePostLikesState(reactions *model.Reactions, user model.User) error {

	query := "UPDATE reactions SET likes = ?, dislikes = ?, totalLikes = ?, totalDislikes = ?, likeState = ?, dislikeState = ? WHERE postID = ?"
	_, err := r.store.db.Exec(query, reactions.Likes, reactions.Dislikes, reactions.TotalLikes, reactions.TotalDislikes, reactions.LikeState, reactions.DisLikeState, reactions.PostID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
