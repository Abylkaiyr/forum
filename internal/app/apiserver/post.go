package apiserver

import "net/http"

func (c *APIServer) Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "post.html", nil)
	case http.MethodPost:
		c.respond(w, r, http.StatusMethodNotAllowed, nil)
	}
}
