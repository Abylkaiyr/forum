package apiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

func (c *APIServer) Home(w http.ResponseWriter, r *http.Request) {
	// The whole data should be send to the home page

	category := r.URL.Query().Get("category")
	type Data struct {
		User  string
		Posts []model.Post1
	}
	p := Data{}
	// identify the useraf
	userID, ok := r.Context().Value(keyUserId).(int)
	if !ok {
		fmt.Println("Could not get user_id from sessions")
		return
	}

	// find post from db
	var err error

	// Categorizing
	if category == "" || category == "all" {
		p.Posts, err = c.store.User().GetAllPosts()
		if err != nil {
			fmt.Println("Could not get all posts from posts")
			return
		}
	} else {
		p.Posts, err = c.store.User().GetPostByCategory(category)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < len(p.Posts); i++ {
		reactions, err := c.store.User().GetReactions(p.Posts[i].ID)
		if err != nil {
			c.logger.ErrLog.Printf("error in getting post likes from reactions %s", err)
		}
		p.Posts[i].Likes = reactions.TotalLikes
		p.Posts[i].Dislikes = reactions.TotalDislikes
		// add dislikes here
	}

	if userID == 0 {
		tpl.ExecuteTemplate(w, "index.html", p)
	} else {
		user, err := c.store.User().FindUserByUserID(userID)
		p.User = user.Username
		if err != nil {
			c.logger.ErrLog.Printf("Could not find user from the users table %s", err)
			return
		}
		tpl.ExecuteTemplate(w, "index.html", p)
	}

}
