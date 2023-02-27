package store

import (
	"log"

	"github.com/Abylkaiyr/forum/internal/app/model"
	"github.com/Abylkaiyr/forum/internal/app/sessions"
)

func (r *UserRepository) CreateSession(s *sessions.Sessions) error {
	statement, _ := r.store.db.Prepare("INSERT INTO sessions (owner,uuid, expireTime, status) VALUES (?,?,?,?)")
	_, err := statement.Exec(s.Owner, s.UUID, s.ExpireTime, s.Status)
	return err
}

func (r *UserRepository) UpdateSession(s *sessions.Sessions) error {
	query := "UPDATE sessions SET uuid = ?, expireTime = ?, status = ? WHERE owner = ?"
	_, err := r.store.db.Exec(query, s.UUID, s.ExpireTime, s.Status, s.Owner)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (r *UserRepository) FindSessionByUUID(uuid string) (*sessions.Sessions, error) {
	// Finding username by his uuid
	var jk int
	s := sessions.NewSession()
	query := "select * from sessions where uuid = $1"
	rows := r.store.db.QueryRow(query, uuid)
	if err := rows.Scan(&jk, &s.Owner, &s.UUID, &s.ExpireTime, &s.Status); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *UserRepository) FindSessionByName(owner string) (*sessions.Sessions, error) {
	// Finding username by his uuid
	var jk int
	s := sessions.NewSession()
	query := "select * from sessions where owner = $1"
	rows := r.store.db.QueryRow(query, owner)
	if err := rows.Scan(&jk, &s.Owner, &s.UUID, &s.ExpireTime, &s.Status); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *UserRepository) FindUserBySession(owner string) (*model.User, error) {
	// Finding this user from users table and returning it
	u := model.NewUser()
	query1 := "select * from users where username = $1"
	rows1 := r.store.db.QueryRow(query1, owner)
	if err := rows1.Scan(&u.ID, &u.Email, &u.Username, &u.Password); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) DeleteUserSessionByUUID(uuid string) error {
	// Finding this user from users table and returning it
	statement, err := r.store.db.Prepare("DELETE FROM sessions WHERE uuid = $1")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(uuid)
	return err
}
