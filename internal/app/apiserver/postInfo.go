package apiserver

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

func (c *APIServer) PostInfo(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/post-info/"))
	if err != nil {
		c.logger.ErrLog.Printf("Could postID from URL %s", err)
		return
	}
	post, err := c.store.User().GetPostByPostID(id)
	if err != nil {
		c.logger.ErrLog.Printf("Could get post from db %s", err)
		return
	}
	type Data struct {
		User     string
		Post     model.Post1
		UserDat  *model.User
		Comments []model.Comments
	}
	p := Data{}
	p.Post = post
	// identify the useraf
	userID, ok := r.Context().Value(keyUserId).(int)
	if !ok {
		fmt.Println("Could not get user_id from sessions")
		return
	}
	// Dealing with comments
	comment := r.FormValue("comment")

	if userID == 0 {
		cmnts, err := c.store.User().GetAllComments(post)
		if err != nil {
			fmt.Println(err)
		}
		p.Comments = cmnts
		tpl.ExecuteTemplate(w, "post-info.html", p)

	} else {
		user, err := c.store.User().FindUserByUserID(userID)
		if err != nil {
			c.logger.ErrLog.Printf("Could get user from db %s", err)
			return
		}

		p.User = user.Username

		p.UserDat = user
		if err != nil {
			c.logger.ErrLog.Printf("Could not find user from the users table %s", err)
			return
		}
		cmnts, err := c.ProcessComments(p.UserDat, post, comment)
		if err != nil {
			c.logger.ErrLog.Printf("error: %s", err)
			return
		}
		p.Comments = cmnts
		tpl.ExecuteTemplate(w, "post-info.html", p)
	}

}
