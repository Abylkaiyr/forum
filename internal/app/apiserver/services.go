package apiserver

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

func (c *APIServer) ProcessImageFile(w http.ResponseWriter, header *multipart.FileHeader, file multipart.File) (*os.File, error) {

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		return nil, err
	}
	tempFile, err := os.Create(fmt.Sprintf("./uploads/%s", "temp_"+header.Filename))
	if err != nil {
		http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
		return nil, fmt.Errorf("error creating temporary file")
	}
	defer tempFile.Close()

	// Copy the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return nil, fmt.Errorf("error copying file")
	}

	return tempFile, nil
}

func (c *APIServer) ProcessImageToSaveInDb(tempFile *os.File, header *multipart.FileHeader) (bytes.Buffer, error) {
	var buf bytes.Buffer

	file, err := os.Open("./uploads/" + "temp_" + header.Filename)
	if err != nil {

		fmt.Println("error opening file:", err)
		return buf, fmt.Errorf("error opening file")
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("error decoding image:", err)
		return buf, fmt.Errorf("error decoding image")
	}

	err = jpeg.Encode(&buf, img, nil)
	if err != nil {
		fmt.Println("error encoding image:", err)
		return buf, fmt.Errorf("error encoding image")
	}
	return buf, nil
}

func (c *APIServer) ProcessPostLikes(post *model.Post1, user model.User) {
	// Find reeaction by post ID

	reactions, err := c.store.User().GetReactions(post.ID)
	if err != nil {
		c.logger.ErrLog.Fatalf("error in getting reactions from reactions table, %s", err)
		return
	}

	// Check if this user has already disliked the same post

	if reactions.PostLiker == "" {
		if reactions.TotalDislikes != 0 {
			c.checkUserDislikehas(reactions, user)
		}

		postLiker := user.Username + ","
		reactions.Likes = 1
		reactions.TotalLikes++
		reactions.LikeState += user.Username + ","
		err := c.store.User().CreateReactionLike(reactions, user, post, postLiker, reactions.PostDisLiker)
		if err != nil {
			c.logger.ErrLog.Fatalf("error in inserting reactions for reactions table, %s", err)
			return
		}

	} else {
		if reactions.TotalDislikes != 0 {
			c.checkUserDislikehas(reactions, user)
		}

		if !strings.Contains(reactions.PostLiker, user.Username) && len(reactions.PostLiker) != 0 {
			reactions.PostLiker += user.Username
			reactions.TotalLikes += 1
			reactions.LikeState += "," + user.Username + ","
			err := c.store.User().UpdatePostLikesInitial(reactions, user)
			if err != nil {
				c.logger.ErrLog.Fatalf("error in updating reactions for reactions table, %s", err)
				return
			}

		} else {
			arr := strings.Split(reactions.LikeState, ",")
			if strings.Contains(reactions.LikeState, user.Username) {
				reactions.TotalLikes--
				stateArr := removeStringElement(arr, user.Username)

				if len(stateArr) == 0 {
					reactions.LikeState = ""
				} else {
					reactions.LikeState = strings.Join(stateArr, ",")
				}
				err = c.store.User().UpdatePostLikesState(reactions, user)
				if err != nil {
					c.logger.ErrLog.Fatalf("error in updating reactions Like State for reactions table, %s", err)
					return
				}
			} else {
				reactions.TotalLikes++
				arr = append(arr, user.Username)
				reactions.LikeState = strings.Join(arr, ",")
				err = c.store.User().UpdatePostLikesState(reactions, user)
				if err != nil {
					c.logger.ErrLog.Fatalf("error in updating reactions Like State for reactions table, %s", err)
					return
				}

			}

		}

	}

}
func (c *APIServer) ProcessPostDisLikes(post *model.Post1, user model.User) {
	// Find reeaction by post ID

	reactions, err := c.store.User().GetReactions(post.ID)
	if err != nil {
		c.logger.ErrLog.Fatalf("error in getting reactions from reactions table, %s", err)
		return
	}
	if reactions.PostDisLiker == "" && reactions.PostLiker == "" {
		if reactions.TotalLikes != 0 {
			c.checkUserLikehas(reactions, user)
		}
		postDisliker := user.Username + ","
		reactions.Dislikes = 1
		reactions.TotalDislikes++
		reactions.DisLikeState += user.Username + ","
		err := c.store.User().CreateReactionLike(reactions, user, post, reactions.PostLiker, postDisliker)
		if err != nil {
			c.logger.ErrLog.Fatalf("error in inserting reactions for reactions table, %s", err)
			return
		}

	} else {
		if reactions.TotalLikes != 0 {
			c.checkUserLikehas(reactions, user)
		}
		if !strings.Contains(reactions.PostDisLiker, user.Username) && len(reactions.PostDisLiker) != 0 {

			reactions.PostDisLiker += user.Username
			reactions.TotalDislikes += 1
			reactions.DisLikeState += "," + user.Username + ","
			err := c.store.User().UpdatePostLikesInitial(reactions, user)
			if err != nil {
				c.logger.ErrLog.Fatalf("error in updating reactions for reactions table, %s", err)
				return
			}

		} else {
			arr := strings.Split(reactions.DisLikeState, ",")
			if strings.Contains(reactions.DisLikeState, user.Username) {
				reactions.TotalDislikes--
				stateArr := removeStringElement(arr, user.Username)

				if len(stateArr) == 0 {
					reactions.DisLikeState = ""
				} else {
					reactions.DisLikeState = strings.Join(stateArr, ",")
				}
				err = c.store.User().UpdatePostLikesState(reactions, user)
				if err != nil {
					c.logger.ErrLog.Fatalf("error in updating reactions Like State for reactions table, %s", err)
					return
				}
			} else {
				reactions.TotalDislikes++
				arr = append(arr, user.Username)
				reactions.DisLikeState = strings.Join(arr, ",")
				err = c.store.User().UpdatePostLikesState(reactions, user)
				if err != nil {
					c.logger.ErrLog.Fatalf("error in updating reactions Like State for reactions table, %s", err)
					return
				}

			}

		}

	}

}

func removeStringElement(s []string, user string) []string {
	for i, v := range s {
		if v == user {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	return s
}

func (c *APIServer) checkUserDislikehas(reactions *model.Reactions, user model.User) {
	var err error
	if strings.Contains(reactions.DisLikeState, user.Username) {
		arr := strings.Split(reactions.DisLikeState, ",")
		reactions.TotalDislikes--
		stateArr := removeStringElement(arr, user.Username)

		if len(stateArr) == 0 {
			reactions.DisLikeState = ""
		} else {
			reactions.DisLikeState = strings.Join(stateArr, ",")
		}
		err = c.store.User().UpdatePostLikesState(reactions, user)
		if err != nil {
			c.logger.ErrLog.Fatalf("error in updating reactions Like State for reactions table, %s", err)
			return
		}
	}
}

func (c *APIServer) checkUserLikehas(reactions *model.Reactions, user model.User) {
	var err error
	if strings.Contains(reactions.LikeState, user.Username) {
		arr := strings.Split(reactions.LikeState, ",")
		reactions.TotalLikes--
		stateArr := removeStringElement(arr, user.Username)

		if len(stateArr) == 0 {
			reactions.LikeState = ""
		} else {
			reactions.LikeState = strings.Join(stateArr, ",")
		}
		err = c.store.User().UpdatePostLikesState(reactions, user)
		if err != nil {
			c.logger.ErrLog.Fatalf("error in updating reactions Like State for reactions table, %s", err)
			return
		}
	}
}

// Comments

func (c *APIServer) ProcessComments(user *model.User, post model.Post1, comment string) ([]model.Comments, error) {

	// if there is no comments, we should create a comment
	if len(comment) != 0 && len(user.Username) != 0 {
		cn := &model.Comments{}
		cn.Content = comment
		cn.PostID = post.ID
		cn.Owner = user.Username
		cn.Likes = 0
		cn.Dislikes = 0
		cn.LikeState = ""
		cn.DislikeState = ""
		if err := c.store.User().CreateComment(post, cn); err != nil {
			c.logger.ErrLog.Printf("db error: %s", err)
			return nil, err
		}
	}
	comments, err := c.store.User().GetAllComments(post)
	if err != nil {
		c.logger.ErrLog.Printf("error is happened: %s", err)
		return nil, err
	}
	return comments, nil
}
