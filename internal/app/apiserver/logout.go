package apiserver

import (
	"log"
	"net/http"
	"time"
)

func (c *APIServer) Logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		c.respond(w, r, http.StatusNotFound, nil)
		return
	}
	if r.Method != http.MethodGet {
		c.respond(w, r, http.StatusMethodNotAllowed, nil)
		return
	}

	if r.Method == http.MethodPost {
		c.respond(w, r, http.StatusMethodNotAllowed, nil)
		return
	}
	cookie, err := r.Cookie("cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			log.Println("error: nil cookies")
			c.respond(w, r, http.StatusUnauthorized, nil)
			return
		}
		log.Println("error: can't get cookie")
		c.respond(w, r, http.StatusBadRequest, nil)
		return
	}

	if err := c.store.User().DeleteUserSessionByUUID(cookie.Value); err != nil {
		log.Println("error: could not find get cookie")
		c.respond(w, r, http.StatusBadRequest, nil)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "cookie",
		Value:   "",
		Expires: time.Now(),
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
