package apiserver

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *APIServer) postDisLike(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(keyUserId).(int)
	if userID == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	postId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/post-dislike/"))
	if err != nil {
		c.logger.ErrLog.Printf("could not get post id %s", err)
		return
	}
	if !ok {
		c.logger.ErrLog.Printf("possibly unauthorized user")
		return
	}
	user, err := c.store.User().FindUserByUserID(userID)

	if err != nil {
		c.logger.ErrLog.Printf("could not find user from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	post, err := c.store.User().GetPostByPostID(postId)
	if err != nil {
		c.logger.ErrLog.Printf("Could not find user from the users table %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		c.logger.ErrLog.Fatalf("get method is not allowed for this path: %s", r.URL.RawPath)
		fmt.Fprint(w, "Method is not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case http.MethodPost:
		c.ProcessPostDisLikes(&post, *user)
	}

	// 	reactions, err := c.store.User().GetPostLikesByPostId(postId)

	http.Redirect(w, r, "/", http.StatusSeeOther)
	// }
}
