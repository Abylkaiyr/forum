package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

func (c *APIServer) Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		return
	}

	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "register.html", nil)
	case http.MethodPost:
		r.ParseForm()
		user := model.NewUser()
		user.Username = r.FormValue("name")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("Password")
		if err := user.Validate(); err != nil {
			tpl.ExecuteTemplate(w, "register.html", err)
			c.logger.ErrLog.Printf("Error type: %v", err)
			return
		}
		// Before create just encrypts password
		if err := user.BeforeCreate(); err != nil {
			tpl.ExecuteTemplate(w, "register.html", err)
			c.logger.ErrLog.Printf("Error type: %v", err)
			return
		}

		// user is created in db here
		if err := c.store.User().Create(user); err != nil {
			tpl.ExecuteTemplate(w, "register.html", "User with the same name already exists")
			c.logger.ErrLog.Printf("Error type: %v", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		tpl.ExecuteTemplate(w, "login.html", nil)
	}
}

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
			tpl.ExecuteTemplate(w, "login.html", fmt.Sprintf("Email or password is incorrect"))
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

func (c *APIServer) Home(w http.ResponseWriter, r *http.Request) {
	// The whole data should be send to the home page

	type Data struct {
		User string
	}
	p := Data{}
	// identify the user
	userID, ok := r.Context().Value(keyUserId).(int)
	if !ok {
		fmt.Println("Could not get user_id from sessions")
		return
	}

	if userID == 0 {
		tpl.ExecuteTemplate(w, "index.html", "adas")
	} else {
		user, err := c.store.User().FindUserByUserID(userID)
		p.User = user.Username
		if err != nil {
			c.logger.ErrLog.Printf("Could not find user from the users table %s", err)
			return
		}
	}
	tpl.ExecuteTemplate(w, "index.html", p)
}

func (c *APIServer) Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "post.html", nil)
	case http.MethodPost:
		c.respond(w, r, http.StatusMethodNotAllowed, nil)
	}
}

func (c *APIServer) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createpost" {
		c.respond(w, r, http.StatusNotFound, nil)
		return
	}

	if r.Method == http.MethodGet {
		c.respond(w, r, http.StatusMethodNotAllowed, nil)
	}

	if r.Method == http.MethodPost {
		// Getting post valyes
		post := model.Post{}
		r.ParseForm()
		post.Title = r.FormValue("newtitle")
		post.Content = r.FormValue("newcontent")
		postType := r.Form["newcategory"]

		for _, l := range postType {
			post.Type += l + " "
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
		post.Owner = user.Username //matching user and post
		// Save user post in db
		if err := c.store.User().CreatePostByUsername(&post); err != nil {
			c.respond(w, r, http.StatusInternalServerError, nil)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
