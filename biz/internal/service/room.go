package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
	"github.com/li1553770945/sheepim-room-service/biz/constant"
	"github.com/li1553770945/sheepim-room-service/kitex_gen/base"
	"github.com/li1553770945/sheepim-room-service/kitex_gen/room"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // 创建随机种子
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func (s *RoomService) CreateRoom(ctx context.Context) (resp *room.CreateRoomResp, err error) {
	roomId := generateRandomString(6)
	exist := true
	for exist == true {
		exist, err = s.Repo.IsRoomExist(roomId)
		if err != nil {
			return nil, err
		}
		if exist {
			roomId = generateRandomString(6)
		}
	}

	clientId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	clientIdStr := clientId.String()
	rpcResp, err := s.AuthClient.GetClientToken(ctx, &auth.GetClientTokenReq{
		ClientId: clientIdStr,
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.BaseResp.Code != 0 {
		return &room.CreateRoomResp{BaseResp: &base.BaseResp{
			Code:    rpcResp.BaseResp.Code,
			Message: rpcResp.BaseResp.Message,
		}}, nil
	}
	err = s.Repo.CreateRoom(roomId, clientIdStr)
	if err != nil {
		return nil, err
	}
	return &room.CreateRoomResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
		RoomId:      &roomId,
		ClientId:    &clientIdStr,
		ClientToken: rpcResp.Token,
	}, nil

}

func (s *RoomService) JoinRoom(ctx context.Context, req *room.JoinRoomReq) (resp *room.JoinRoomResp, err error) {
	exist, err := s.Repo.IsRoomExist(req.RoomId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return &room.JoinRoomResp{
			BaseResp: &base.BaseResp{
				Code:    constant.NotFound,
				Message: "房间不存在",
			},
		}, nil
	}

	// 生成用户的 Client ID
	clientId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	clientIdStr := clientId.String()
	// 获取用户的 Token
	rpcResp, err := s.AuthClient.GetClientToken(ctx, &auth.GetClientTokenReq{
		ClientId: clientIdStr,
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.BaseResp.Code != 0 {
		return &room.JoinRoomResp{BaseResp: &base.BaseResp{
			Code:    rpcResp.BaseResp.Code,
			Message: rpcResp.BaseResp.Message,
		}}, nil
	}
	clientToken := *rpcResp.Token

	// 将用户加入房间
	err = s.Repo.JoinRoom(req.RoomId, clientIdStr)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &room.JoinRoomResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
		ClientId:    &clientIdStr,
		ClientToken: &clientToken,
	}, nil
}

func (s *RoomService) GetRoomMembers(ctx context.Context, req *room.GetRoomMembersReq) (resp *room.GetRoomMembersResp, err error) {
	exist, err := s.Repo.IsRoomExist(req.RoomId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return &room.GetRoomMembersResp{
			BaseResp: &base.BaseResp{
				Code:    constant.NotFound,
				Message: "房间不存在",
			},
		}, nil
	}
	members, err := s.Repo.GetRoomMembers(req.RoomId)
	if err != nil {
		return nil, err
	}
	resp = &room.GetRoomMembersResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
		Members: members,
	}
	return resp, nil
}

func (s *RoomService) CheckIsInRoom(ctx context.Context, req *room.CheckIsInRoomReq) (resp *room.CheckIsInRoomResp, err error) {
	isInRoom, err := s.Repo.CheckIsInRoom(req.RoomId, req.ClientId)
	if err != nil {
		return nil, err
	}
	return &room.CheckIsInRoomResp{
		BaseResp: &base.BaseResp{
			Code: 0,
		},
		IsInRoom: &isInRoom,
	}, nil
}
