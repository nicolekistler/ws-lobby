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
	Round int `json:"Round" binding:"required"`
}

type Settings struct {
	MaxNumberRounds int `json:"maxNumberRounds" binding:"required"`
	MaxNumberUsers  int `json:"maxNumberUsers" binding:"required"`
}

type User struct {
	ID       string `json:"id"  binding:"required"`
	Name     string `json:"name" binding:"required"`
	AvatarID string `json:"avatarId" binding:"required"`
	Host     bool   `json:"Host" binding:"required"`
}

type Room struct {
	ID       string `json:"id"`
	State    RoomState
	Settings Settings
	Users    []User
}

type CreateRoomRequest struct {
	Name            string `json:"name" binding:"required"`
	MaxNumberRounds int    `json:"maxNumberRounds" binding:"required"`
	MaxNumberUsers  int    `json:"maxNumberUsers" binding:"required"`
	AvatarID        string `json:"avatarId" binding:"required"`
}
