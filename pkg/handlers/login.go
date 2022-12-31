package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

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
		rows, _ := database.Query("select * from users where username like '" + userName + "' ")
		// if err != nil {
		// 	fmt.Println("Could not find you from Database")
		// }
		user := &utils.User{}

		for rows.Next() {
			rows.Scan(&user.ID, &user.UserEmail, &user.UserName, &user.UserPassword)
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password))
		if err != nil || user.UserName == "" { // Means user is not found from database
			w.WriteHeader(http.StatusUnauthorized)
			tpl.ExecuteTemplate(w, "login.html", "Username or password incorrect")
		}
		fmt.Fprint(w, "Succeed")
		delivery.SetSession(*user, w)
	}
}
