package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Abylkaiyr/forum/pkg/delivery"
	"github.com/Abylkaiyr/forum/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/login" {
		Errors(w, http.StatusNotFound, fmt.Errorf("NOT FOUND REQUEST FROM  %s", r.RemoteAddr))
		return
	}

	switch r.Method {
	case http.MethodGet:
		tpl.ExecuteTemplate(w, "login.html", nil)
	case http.MethodPost:
		r.ParseForm()
		userName := r.FormValue("username")
		password := r.FormValue("password")
		database, err := sql.Open("sqlite3", "./storage.db")
		if err != nil {
			Errors(w, http.StatusInternalServerError, fmt.Errorf("ERROR in opening DataBase"))
			return
		}
		defer database.Close()
		rows, _ := database.Query("select id, username, password from users where username like '" + userName + "' ")
		// if err != nil {
		// 	fmt.Println("Could not find you from Database")
		// }
		user := utils.User{}

		for rows.Next() {
			rows.Scan(&user.ID, &user.UserName, &user.UserPassword)
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password))
		if err != nil || user.UserName == "" { // Means user is not found from database
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			tpl.ExecuteTemplate(w, "login.html", "Username or password incorrect")
		}

		delivery.SetSession(user.ID, w)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
