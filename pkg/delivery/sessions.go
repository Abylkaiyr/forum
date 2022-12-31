package delivery

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Abylkaiyr/forum/pkg/utils"
)

// var u *utils.User

func SetSession(user utils.User, w http.ResponseWriter) {
	if user.SessionID == "" || user.ExpireTime.Before(time.Now()) {
		http.SetCookie(w, &http.Cookie{
			Name:    "COOKIE_NAME",
			Value:   uuid.NewString(),
			Expires: time.Now().Add(120 * time.Second),
		})
	}
}
