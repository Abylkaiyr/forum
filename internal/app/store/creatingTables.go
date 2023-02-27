package store

import (
	"database/sql"

	"github.com/Abylkaiyr/forum/migrations"
	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) CreateTables() error {
	if err := s.Open(); err != nil {
		return err
	}

	createTable(s.db, migrations.TableForUsers)
	createTable(s.db, migrations.TableForPosts)
	createTable(s.db, migrations.TableForComments)
	createTable(s.db, migrations.TableForSessions)
	createTable(s.db, migrations.TableForPostReactions)
	return nil
}

func createTable(db *sql.DB, str string) error {
	statement, err := db.Prepare(str)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}
