package delivery

import (
	"database/sql"

	"github.com/Abylkaiyr/forum/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func newTable(db *sql.DB, str string) *Database {
	statement, _ := db.Prepare(str)
	statement.Exec()
	return &Database{
		DB: db,
	}
}

func StorageInit() {
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		ErrorMsg(err)
		return
	}

	newTable(db, utils.TableForUsers)
	newTable(db, utils.TableForPosts)
	newTable(db, utils.TableForComments)
}
