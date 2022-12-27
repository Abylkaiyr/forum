package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, fmt.Errorf("NOT FOUND REQUEST FROM  %s", r.RemoteAddr))
		return
	}

	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed, fmt.Errorf("%v METHOD IS NOT ALLOWED FROM  %s", r.Method, r.RemoteAddr))
		return
	}

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
