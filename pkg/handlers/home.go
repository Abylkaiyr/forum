package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Abylkaiyr/forum/pkg/utils"
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

	c, err := r.Cookie("COOKIE_NAME")
	if err != nil {
		fmt.Println("Could not get cookie for user")
	}

	database, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		Errors(w, http.StatusInternalServerError, fmt.Errorf("ERROR in opening DataBase"))
		return
	}
	defer database.Close()
	database.Ping()
	rows, _ := database.Query("select * from sessions where uuid = '" + c.Value + "' ")
	if err != nil {
		fmt.Println("Could not find you from Database")
	}
	var session = &utils.Sessions{}

	for rows.Next() {
		rows.Scan(session.UserID, session.SessionID, session.ExpireTime)
	}
	fmt.Fprint(w, session.SessionID, session.ExpireTime)
}
