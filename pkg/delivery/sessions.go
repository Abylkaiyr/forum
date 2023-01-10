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

}

func DeleteSession(w http.ResponseWriter) {
	// Delete session for the user from db
	database, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		fmt.Println("Here")
	}

	database.Ping()
	statement, err := database.Prepare("DROP TABLE IF EXISTS sessions")
	if err != nil {
		fmt.Println("Could not drop table")
	}
	statement.Exec()
	fmt.Println("Dropped table")
}
