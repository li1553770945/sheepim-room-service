package service

import (
	"context"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-room-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-room-service/kitex_gen/room"
)

type RoomService struct {
	Repo       repo.IRepository
	AuthClient authservice.Client
}

type IRoomService interface {
	CreateRoom(ctx context.Context) (resp *room.CreateRoomResp, err error)
	JoinRoom(ctx context.Context, req *room.JoinRoomReq) (resp *room.JoinRoomResp, err error)
	GetRoomMembers(ctx context.Context, req *room.GetRoomMembersReq) (resp *room.GetRoomMembersResp, err error)
	CheckIsInRoom(ctx context.Context, req *room.CheckIsInRoomReq) (resp *room.CheckIsInRoomResp, err error)
}

func NewRoomService(repo repo.IRepository, authClient authservice.Client) IRoomService {
	return &RoomService{
		Repo:       repo,
		AuthClient: authClient,
	}
}
