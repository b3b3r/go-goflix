package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Store is the store for DB methods
type Store interface {
	Open() error
	Close() error
	GetMovies() ([]*Movie, error)
	GetMovieByID(id int64) (*Movie, error)
	CreateMovie(m *Movie) error
	CreateUser(u *User) error
	GetUserByID(id int64) (*User, error)
	FindUser(username, password string) (bool, error)
}

type dbStore struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
);

CREATE TABLE IF NOT EXISTS user
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT,
	password TEXT
);`

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	store.db = db
	db.MustExec(schema)
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}

func (store *dbStore) GetMovies() ([]*Movie, error) {
	var movies []*Movie
	err := store.db.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (store *dbStore) GetMovieByID(id int64) (*Movie, error) {
	var movie = &Movie{}
	err := store.db.Get(movie, "SELECT * FROM movie WHERE id=$1", id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (store *dbStore) CreateMovie(m *Movie) error {
	res, err := store.db.Exec("INSERT INTO movie (title, release_date, duration, trailer_url) VALUES(?,?,?,?)",
		m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)

	if err != nil {
		return err
	}

	m.ID, err = res.LastInsertId()
	return err
}

func (store *dbStore) CreateUser(u *User) error {
	res, err := store.db.Exec("INSERT INTO user (username, password) VALUES(?,?)",
		u.Username, u.Password)

	if err != nil {
		return err
	}

	u.ID, err = res.LastInsertId()
	return err
}

func (store *dbStore) GetUserByID(id int64) (*User, error) {
	var user = &User{}
	err := store.db.Get(user, "SELECT * FROM user WHERE id=$1", id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (store *dbStore) FindUser(username, password string) (bool, error) {
	var count int
	err := store.db.Get(&count, "SELECT COUNT(id) FROM user WHERE username=$1 AND password=$2", username, password)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
