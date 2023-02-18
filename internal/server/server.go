package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func NewServer() *Server {
	serv := &Server{
		Rooms: []Room{
			{
				ID:       "1",
				State:    RoomState{Round: 0},
				Settings: Settings{MaxNumberRounds: 10},
				Users: []User{
					{
						ID:       "1",
						Name:     "Stella",
						AvatarID: "4",
						IsHost:   true,
						TeamId:   "1",
					},
				},
			},
			{
				ID:       "2",
				State:    RoomState{Round: 0},
				Settings: Settings{MaxNumberRounds: 7},
				Users: []User{
					{
						ID:       "4",
						Name:     "Possum",
						AvatarID: "3",
						IsHost:   true,
						TeamId:   "2",
					},
					{
						ID:       "6",
						Name:     "Rila",
						AvatarID: "7",
						IsHost:   true,
						TeamId:   "2",
					},
				},
			},
		},
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	return serv
}

func (s *Server) Init() error {
	s.Router = mux.NewRouter().StrictSlash(true)

	s.registerApiRoutes()
	s.registerSocketHandlers()

	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("../public")))

	if err := http.ListenAndServe(":8080", s.Router); err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}

	return nil
}
