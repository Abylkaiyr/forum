package apiserver

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

func (c *APIServer) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createpost" {
		c.respond(w, r, http.StatusNotFound, nil)
		return
	}

	if r.Method == http.MethodGet {
		c.respond(w, r, http.StatusMethodNotAllowed, nil)
	}

	if r.Method == http.MethodPost {
		// Getting post values
		post := model.Post{}
		if err := r.ParseMultipartForm(5 << 20); err != nil {
			fmt.Println(err)
			return
		}

		post.Title = r.FormValue("newtitle")
		if post.Title == "" {
			tpl.ExecuteTemplate(w, "post.html", "Post should always have a title")
			c.logger.ErrLog.Printf("User submitted post with empty post title")
			return
		}
		post.Content = r.FormValue("newcontent")
		if post.Content == "" {
			tpl.ExecuteTemplate(w, "post.html", "Post should always have a non emty content")
			c.logger.ErrLog.Printf("User submitted post with empty post content")
			return
		}
		post.CreatedTime = time.Now()
		postType := r.Form["newcategory"]
		if len(postType) == 0 {
			tpl.ExecuteTemplate(w, "post.html", "Post should always have at least one category")
			c.logger.ErrLog.Printf("User submitted post with unselected  post category")
			return
		}

		var p string
		for _, v := range postType {
			p += v + ","
		}
		post.Type = strings.TrimSuffix(p, ",")
		// working with uploaded image file
		file, header, err := r.FormFile("image")

		if err != nil {
			tpl.ExecuteTemplate(w, "post.html", "You should submit non-empty and only image file")
			c.logger.ErrLog.Printf("User submitted wrong file or empty file %s", err)
			return
		}
		defer file.Close()
		tempFile, err := c.ProcessImageFile(w, header, file)
		if err != nil {
			fmt.Println(err)
			return
		}
		buf, err := c.ProcessImageToSaveInDb(tempFile, header)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Saving the post by user id in db
		userID, ok := r.Context().Value(keyUserId).(int)
		if !ok {
			fmt.Println("Could not get user_id from sessions")
			return
		}
		// Finding username by its id value
		user, err := c.store.User().FindUserByUserID(userID)
		if err != nil {
			c.logger.ErrLog.Printf("Could not find user from the users table %s", err)
			return
		}
		post.Owner = user.Username // matching user and post
		// Save user post in db
		if err := c.store.User().CreatePostByUsername(&post, buf, header); err != nil {
			c.respond(w, r, http.StatusInternalServerError, nil)
			return
		}

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
