package apiserver

import (
	"fmt"
	"net/http"
)

func (c *APIServer) Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		c.respond(w, r, http.StatusNotFound, fmt.Sprintf("Not found request from %s", r.RemoteAddr))
		c.logger.ErrLog.Printf("Invalid URL requested: %s", r.RemoteAddr)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "login.html", nil)
	case http.MethodPost:
		r.ParseForm()
		useremail := r.FormValue("email")
		password := r.FormValue("password")
		user, err := c.store.User().FindByEmail(useremail)
		if err != nil || !user.ComparePassword(password) {
			tpl.ExecuteTemplate(w, "login.html", "Email or password is incorrect")
			c.logger.ErrLog.Printf("User is not find or Incorrect Password %s", err)
			return
		}
		_, err = c.store.User().FindSessionByName(user.Username)
		if err != nil { // Meaning user has not any session
			// we have to set new session for the user and save it in db

			s := c.session.SetSession(user, w)
			if err := c.store.User().CreateSession(&s); err != nil {
				c.logger.ErrLog.Printf("Could not create session for the user %s", err)
				return
			}

		} else {
			s := c.session.SetSession(user, w) // set the new session values for userSession
			if err := c.store.User().UpdateSession(&s); err != nil {
				c.logger.ErrLog.Printf("Could not update session for the user %s", err)
				return
			}
		}
		// in other case meaning that user has already session and we should update it
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
