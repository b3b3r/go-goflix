package main

import (
	"fmt"
)

//Movie is the struct of any movie
type Movie struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	ReleaseDate string `db:"release_date"`
	Duration    int    `db:"duration"`
	TrailerURL  string `db:"trailer_url"`
}

func (m Movie) String() string {
	s := fmt.Sprintf("id=%v, title=%v, releaseDate=%v, duration=%v, trailerURL=%v",
		m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
	return s
}

//User is the struct of any user
type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
