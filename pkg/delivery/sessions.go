package delivery

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// var u *utils.User

func SetSession(userID int, w http.ResponseWriter) {

	rows, _ := DB.Query("select uuid, expireTime from sessions where userID =  '" + strconv.Itoa(userID) + "' ")

	sessionID := ""
	var expireTimeStr string
	for rows.Next() {
		rows.Scan(&sessionID, &expireTimeStr)
	}
	if sessionID == "" {
		fmt.Println("HERE I AM 1")
	}

	fmt.Println(expireTimeStr)
	fmt.Println("I am here 2")

	expireTime, err := time.Parse(time.UnixDate, expireTimeStr)
	if err != nil {
		log.Fatal(err)
	}
	// Inserting values to the db
	if expireTime.After(time.Now()) {
		sessionID = uuid.NewString()
		expireTime = time.Now().Add(120 * time.Second)
		statement, _ := DB.Prepare("INSERT INTO sessions (userID, uuid, expireTime) VALUES (?,?,?)")
		statement.Exec(strconv.Itoa(userID), sessionID, expireTime.String())
		fmt.Println("Reached here")
	}
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "COOKIE_NAME",
	// 	Value:   uuid.NewString(),
	// 	Expires: time.Now().Add(120 * time.Second),
	// })

}
