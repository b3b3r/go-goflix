package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/movies", s.handleMovieList()).Methods("GET")
	s.router.HandleFunc("/api/movies/{id}", s.handleMovieList()).Methods("GET")
}
