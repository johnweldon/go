package go_stat

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type stat struct {
	DB *sql.DB
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (s *stat) init() {
	var err error
	s.DB, err = sql.Open("sqlite3", "./foo.db")
	check_error(err)
}

func (s *stat) exec(sql string) {
	_, err := s.DB.Exec(sql)
	check_error(err)
}

func (s *stat) cleanup() {
	s.DB.Close()
}

func Open() {

	var s = stat{}
	s.init()
	defer s.cleanup()

	s.exec(`CREATE TABLE IF NOT EXISTS main (
                id integer not null primary key,
                name text )`)

}
