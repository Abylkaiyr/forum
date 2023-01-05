package delivery

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func SetSession(userID int, w http.ResponseWriter) {

	database, err := sql.Open("sqlite3", "./storage.db")

	if err != nil {
		fmt.Println("Here")
	}

	database.Ping()

	rows, err := database.Query("select uuid, expireTime from sessions where userID =  '" + strconv.Itoa(userID) + "' ")

	if err != nil {
		fmt.Fprint(w, "could not select")
		return
	}

	sessionID := ""
	var expireTimeStr string

	for rows.Next() {
		rows.Scan(&sessionID, &expireTimeStr)
	}

	expireTime, err := time.Parse(time.RFC3339Nano, expireTimeStr)
	if err != nil {
		fmt.Println("Error in parsing time value here2")
	}

	if checkSession(sessionID, expireTimeStr) || checkExpireSession(expireTime) {
		// if session is empty we set new session and write it to db
		sessionID = uuid.NewString()
		expireTime := time.Now().Add(120 * time.Second)

		http.SetCookie(w, &http.Cookie{
			Name:    "COOKIE_NAME",
			Value:   sessionID,
			Expires: expireTime,
		})
		expireTimeStr = expireTime.Format(time.RFC3339Nano)
		if err != nil {
			fmt.Println("error in parsing time")
		}

		statement, _ := database.Prepare("INSERT INTO sessions (userID, uuid, expireTime) VALUES (?,?,?)")
		statement.Exec(strconv.Itoa(userID), sessionID, expireTimeStr)
	}
}

func checkSession(sessionID string, expireTimeStr string) bool {
	checker := false
	if sessionID == "" || expireTimeStr == "" {
		checker = true
	}

	return checker
}

func checkExpireSession(expireTime time.Time) bool {
	var checker = false
	if expireTime.Before(time.Now()) {
		checker = true
	}
	return checker
}
