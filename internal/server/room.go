package server

import (
	"math/rand"
	"strconv"
)

func NewRoom(requestData CreateRoomRequest) *Room {

	return &Room{
		ID: strconv.Itoa(rand.Intn(1000000)),
		Users: []User{
			{
				Name:     requestData.Name,
				Host:     true,
				AvatarID: requestData.AvatarID,
			},
		},
		State: RoomState{Round: 0},
		Settings: Settings{
			MaxNumberRounds: requestData.MaxNumberRounds,
			MaxNumberUsers:  requestData.MaxNumberUsers,
		},
	}
}
