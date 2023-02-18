package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerApiRoutes() {
	s.Router.HandleFunc("/rooms", s.getRooms).Methods("GET")
	s.Router.HandleFunc("/rooms", s.createRoom).Methods("POST")
	s.Router.HandleFunc("/rooms/{id}", s.getRoom).Methods("GET")
}

// Get all game rooms
func (s *Server) getRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.Rooms)
}

// Get a game room by ID
func (s *Server) getRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range s.Rooms {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	log.Println("Bruh")

	json.NewEncoder(w).Encode(Room{})
}

// Create a game room
func (s *Server) createRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Create room")
}
