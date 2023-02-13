package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

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

func NewServer() *Server {
	serv := &Server{
		Rooms: []Room{
			{
				ID:    "1",
				State: RoomState{Round: 0},
			},
		},
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	return serv
}
