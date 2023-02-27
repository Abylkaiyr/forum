package store

import "github.com/Abylkaiyr/forum/internal/app/model"

func (r *UserRepository) Create(u *model.User) error {
	statement, _ := r.store.db.Prepare("INSERT INTO users (email,username, password) VALUES (?,?,?)")
	_, err := statement.Exec(u.Email, u.Username, u.EncryptedPassword)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := model.NewUser()
	query := "select * from users where email = $1"
	rows := r.store.db.QueryRow(query, email)
	if err := rows.Scan(&u.ID, &u.Email, &u.Username, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
func (r *UserRepository) FindUserByUserName(username string) (*model.User, error) {
	u := model.NewUser()
	query := "select * from users where username = $1"
	rows := r.store.db.QueryRow(query, username)
	if err := rows.Scan(&u.ID, &u.Email, &u.Username, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindUserByUserID(id int) (*model.User, error) {
	// Finding this user from users table and returning it
	u := model.NewUser()
	query1 := "select * from users where id = $1"
	rows1 := r.store.db.QueryRow(query1, id)
	if err := rows1.Scan(&u.ID, &u.Email, &u.Username, &u.Password); err != nil {
		return nil, err
	}
	return u, nil
}
