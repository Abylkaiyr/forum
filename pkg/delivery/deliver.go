package delivery

import (
	"database/sql"

	"github.com/Abylkaiyr/forum/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func newTable(db *sql.DB, str string) {
	statement, _ := db.Prepare(str)
	statement.Exec()
}

func StorageInit() {
	DB, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		ErrorMsg(err)
		return
	}
	defer DB.Close()

	newTable(DB, utils.TableForUsers)
	newTable(DB, utils.TableForPosts)
	newTable(DB, utils.TableForComments)
	newTable(DB, utils.TableForSessions)

}
