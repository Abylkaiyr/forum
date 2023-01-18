package apiserver

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("pkg/static/templates/*"))
}

func (c *APIServer) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			c.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := c.store.User().Create(u); err != nil {
			c.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		c.respond(w, r, http.StatusCreated, u)
	}
}

func (c *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	c.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (c *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
