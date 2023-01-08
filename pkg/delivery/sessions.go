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
	var expireTime time.Time

	for rows.Next() {
		rows.Scan(&sessionID, &expireTime)
	}

	sessionID = uuid.NewString()

	expireTime = time.Now().Add(120 * time.Second)

	http.SetCookie(w, &http.Cookie{
		Name:    "cookie",
		Value:   sessionID,
		Expires: expireTime.Add(time.Hour * 3), // added 3 hours because in browser time is not settinng
	})

	statement, _ := database.Prepare("INSERT INTO sessions (userID, uuid, expireTime) VALUES (?,?,?)")
	statement.Exec(strconv.Itoa(userID), sessionID, expireTime)

	// if sessions is expired delete the previous session and add new one
	// if checkExpireSession(expireTime) {
	// 	sessionID = uuid.NewString()
	// 	expireTime = time.Now().Add(120 * time.Second)
	// 	// expireTimeStr = expireTime.Format(time.RFC3339Nano)
	// 	expireTimeStr = expireTime.String()
	// 	statement, _ := database.Prepare("UPDATE sessions SET uuid = ?, expireTime = ? WHERE userID = ?")
	// 	statement.Exec(sessionID, expireTimeStr, strconv.Itoa(userID))
	// 	http.SetCookie(w, &http.Cookie{
	// 		Name:    "COOKIE_NAME",
	// 		Value:   sessionID,
	// 		Expires: expireTime,
	// 	})
	// }
}

// func checkSession(sessionID string) bool {
// 	checker := false
// 	if sessionID == "" {
// 		checker = true
// 	}

// 	return checker
// }
