package apiserver

import (
	"context"
	"net/http"

	"github.com/Abylkaiyr/forum/internal/app/model"
)

type key int

const (
	keyUserId key = iota
)

func (c *APIServer) authMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := model.NewUser()
		sessionCookie, err := r.Cookie("cookie")
		if err != nil {
			if err != nil {
				if err == http.ErrNoCookie {
					u.ID = 0
				}
				u.ID = 0
			}
		} else {
			session, err := c.store.User().FindSessionByUUID(sessionCookie.Value)
			if err != nil {
				c.respond(w, r, http.StatusInternalServerError, nil)
			}
			user, err := c.store.User().FindUserBySession(session.Owner)
			if err != nil {
				user.ID = 0
				c.logger.ErrLog.Printf("Error type: %v", err)
				return
			}
			u.ID = user.ID
		}
		ctx := context.WithValue(r.Context(), keyUserId, u.ID)
		next(w, r.WithContext(ctx))
	}
}
