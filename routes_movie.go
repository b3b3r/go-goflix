package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type jsonMovie struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerURL  string `json:"trailer_url"`
}

func mapMovieToJSON(m *Movie) jsonMovie {
	return jsonMovie{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerURL:  m.TrailerURL,
	}
}

func (s *server) handleMovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := s.store.GetMovies()
		if err != nil {
			log.Printf("Cannot load movies, err=%v\n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		var resp = make([]jsonMovie, len(movies))
		for i, m := range movies {
			resp[i] = mapMovieToJSON(m)
		}

		s.respond(w, r, resp, http.StatusOK)
	}
}

func (s *server) handleMovieDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Cannot parse id to int, err=%v\n", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}

		movie, err := s.store.GetMovieByID(id)
		if err != nil {
			log.Printf("Cannot load movie, err=%v\n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapMovieToJSON(movie)
		s.respond(w, r, resp, http.StatusOK)
	}
}
