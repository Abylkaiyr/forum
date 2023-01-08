package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Abylkaiyr/forum/pkg/utils"
)

func MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("cookie")
		if err == nil {
			database, err := sql.Open("sqlite3", "./storage.db")
			if err != nil {
				Errors(w, http.StatusInternalServerError, fmt.Errorf("ERROR in opening DataBase"))
				return
			}
			defer database.Close()
			database.Ping()
			query := "select * from sessions where uuid = $1"
			rows := database.QueryRow(query, c.Value)
			var session = utils.Sessions{}
			rows.Scan(&session.UserID, &session.SessionID, &session.ExpireTime)
			if time.Now().Before(session.ExpireTime) {
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				fmt.Println("You are logged in")
			} else {
				next.ServeHTTP(w, r)
				fmt.Println("Your session is expired")
			}

		} else {
			// w.WriteHeader(http.StatusUnauthorized)
			next.ServeHTTP(w, r)
			fmt.Println("No cookie is found for user")
		}
	}
}
