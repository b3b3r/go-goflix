package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type jsonUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func mapUserToJSON(u *User) jsonUser {
	return jsonUser{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}

func (s *server) handleUserDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Cannot parse id to int, err=%v\n", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}

		user, err := s.store.GetUserByID(id)
		if err != nil {
			log.Printf("Cannot load user, err=%v\n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapUserToJSON(user)
		s.respond(w, r, resp, http.StatusOK)
	}
}

func (s *server) handleUserCreate() http.HandlerFunc {
	type requestUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := requestUser{}
		err := s.decode(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse movie body, err=%v\n", err)
			s.respond(w, r, nil, http.StatusBadRequest)
			return
		}
		u := &User{
			ID:       0,
			Username: req.Username,
			Password: req.Password,
		}
		err = s.store.CreateUser(u)
		if err != nil {
			log.Printf("Cannot create movie, err=%v\n", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		resp := mapUserToJSON(u)
		s.respond(w, r, resp, http.StatusOK)
	}
}
