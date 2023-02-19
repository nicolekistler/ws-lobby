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

	json.NewEncoder(w).Encode(Room{})
}

// Create a game room
func (s *Server) createRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var requestData CreateRoomRequest

	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room := NewRoom(requestData)

	s.Rooms = append(s.Rooms, *room)
	json.NewEncoder(w).Encode(room)

	log.Printf("Received POST request with data: %+v", requestData)
	log.Printf("new arr: %+v", s.Rooms)
}
