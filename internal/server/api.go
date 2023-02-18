package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

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

	json.NewEncoder(w).Encode(Room{})
}

// Create a game room
func (s *Server) createRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room Room
	_ = json.NewDecoder(r.Body).Decode(&room)
	room.ID = strconv.Itoa(rand.Intn(1000000))

	s.Rooms = append(s.Rooms, room)
	json.NewEncoder(w).Encode(&room)
}
