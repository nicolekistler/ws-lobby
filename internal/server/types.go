package server

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	Rooms    []Room
	Router   *mux.Router
	Upgrader *websocket.Upgrader
}

type RoomState struct {
	Round int `json:"Round"`
}

type Settings struct {
	MaxNumberRounds int `json:"maxNumberRounds"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	AvatarID string `json:"avatarId"`
	Host     bool   `json:"Host"`
	TeamId   string `json:"teamId"`
}

type Team struct {
	TeamID string `json:"teamId"`
	Score  int    `json:"score"`
	Users  []string
}

type Room struct {
	ID       string `json:"id"`
	State    RoomState
	Settings Settings
	Users    []User
	Teams    []Team
}
