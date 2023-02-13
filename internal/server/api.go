package server

import (
	"log"
	"net/http"
)

func (s *Server) registerApiRoutes() {
	s.Router.HandleFunc("/rooms", s.getRooms).Methods("GET")
	s.Router.HandleFunc("/rooms", s.createRoom).Methods("POST")
	s.Router.HandleFunc("/rooms/{id}", s.getRoom).Methods("GET")
}

// Get all game rooms
func (s *Server) getRooms(w http.ResponseWriter, r *http.Request) {
	log.Println(s.Rooms)
}

// Create a game room
func (s *Server) createRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Create room")
}

// Get a game room by ID
func (s *Server) getRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Get room")
}
