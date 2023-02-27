package apiserver

import (
	"net/http"

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
