package main

import (
	"context"
	"github.com/li1553770945/sheepim-room-service/biz/infra/container"
	room "github.com/li1553770945/sheepim-room-service/kitex_gen/room"
)

// RoomServiceImpl implements the last service interface defined in the IDL.
type RoomServiceImpl struct{}

// CreateRoom implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) CreateRoom(ctx context.Context) (resp *room.CreateRoomResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.RoomService.CreateRoom(ctx)
	return
}

// JoinRoom implements the RoomServiceImpl interface.
func (s *RoomServiceImpl) JoinRoom(ctx context.Context, req *room.JoinRoomReq) (resp *room.JoinRoomResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.RoomService.JoinRoom(ctx, req)
	return
}
