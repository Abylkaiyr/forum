package handlers

import (
	"net/http"
	"strconv"
	"text/template"
)

var tpl *template.Template

// type AppServer struct {
// 	InfoLog  *log.Logger
// 	ErrorLog *log.Logger
// }

func init() {
	tpl = template.Must(template.ParseGlob("pkg/static/templates/*"))
}

func Errors(w http.ResponseWriter, status int, err error) {
	// app.ErrorLog.Println(err.Error())
	w.WriteHeader(status)
	if err != nil {
		http.Error(w, strconv.Itoa(status)+" "+http.StatusText(status), status)
		return
	}
	statusint := strconv.Itoa(status) + " " + http.StatusText(status)
	tpl.ExecuteTemplate(w, "errors.html", statusint)
}
