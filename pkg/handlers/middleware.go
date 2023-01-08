package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Abylkaiyr/forum/pkg/utils"
)

func MiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("here Middleware")
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

			if session.SessionID == "" || session.ExpireTime.After(time.Now()) {
				w.WriteHeader(http.StatusUnauthorized)
				http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			} else {
				next.ServeHTTP(w, r)
			}
		} else {
			// w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
