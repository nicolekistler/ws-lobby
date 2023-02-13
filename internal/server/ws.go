package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *Server) registerSocketHandlers() {
	s.Router.HandleFunc("/ws", s.handleConnections)
}

func (s *Server) reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func (s *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
	s.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	err = ws.WriteMessage(1, []byte("Hi Client!"))

	if err != nil {
		log.Println(err)
	}

	// helpful log statement to show connections
	log.Println("Client Connected")

	s.reader(ws)
}
