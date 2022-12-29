package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		Errors(w, http.StatusNotFound, fmt.Errorf("NOT FOUND REQUEST FROM  %s", r.RemoteAddr))
		return
	}

	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "register.html", nil)
	case http.MethodPost:
		r.ParseForm()
		userName := r.FormValue("name")
		userEmail := r.FormValue("email")
		userPassword := r.FormValue("Password")
		crpassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), 14)
		if err != nil {
			Errors(w, http.StatusInternalServerError, fmt.Errorf("ERROR in hasing user password"))
			return
		}

		password := string(crpassword)
		database, err := sql.Open("sqlite3", "./storage.db")
		if err != nil {
			Errors(w, http.StatusInternalServerError, fmt.Errorf("ERROR in opening DataBase"))
			return
		}
		defer database.Close()
		DB, err := database.Prepare(`INSERT INTO users(email,username, password) values(?,?,?)`)
		if err != nil {
			Errors(w, http.StatusInternalServerError, fmt.Errorf("ERROR in preparing statement for DB"))
			return
		}
		rows, _ := database.Query("select * from users where email ='" + userEmail + "' or username ='" + userName + "'")
		var id int
		var name string
		var email string
		var password2 string
		for rows.Next() {
			rows.Scan(&id, &email, &name, &password2)
		}
		DB.Exec(userEmail, userName, password)
		err = tpl.ExecuteTemplate(w, "register.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
