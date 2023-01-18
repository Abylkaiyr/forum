package sessions

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Abylkaiyr/forum/internal/app/model"
	"github.com/google/uuid"
)

type Sessions struct {
	db         *sql.DB
	Owner      string
	UUID       string
	ExpireTime time.Time
	Status     int
}

func NewSession() *Sessions {
	return &Sessions{}
}

func (s *Sessions) SetSession(user *model.User, w http.ResponseWriter) Sessions {
	s.Owner = user.Username
	s.UUID = uuid.NewString()
	s.ExpireTime = time.Now().Add(time.Minute * 5)
	s.Status = 1
	http.SetCookie(w, &http.Cookie{
		Name:    "cookie",
		Value:   s.UUID,
		Expires: s.ExpireTime,
	})

	return *s
}
