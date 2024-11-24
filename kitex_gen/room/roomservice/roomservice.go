// Code generated by Kitex v0.7.2. DO NOT EDIT.

package roomservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	room "github.com/li1553770945/sheepim-room-service/kitex_gen/room"
)

func serviceInfo() *kitex.ServiceInfo {
	return roomServiceServiceInfo
}

var roomServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RoomService"
	handlerType := (*room.RoomService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateRoom":     kitex.NewMethodInfo(createRoomHandler, newRoomServiceCreateRoomArgs, newRoomServiceCreateRoomResult, false),
		"JoinRoom":       kitex.NewMethodInfo(joinRoomHandler, newRoomServiceJoinRoomArgs, newRoomServiceJoinRoomResult, false),
		"GetRoomMembers": kitex.NewMethodInfo(getRoomMembersHandler, newRoomServiceGetRoomMembersArgs, newRoomServiceGetRoomMembersResult, false),
		"CheckIsInRoom":  kitex.NewMethodInfo(checkIsInRoomHandler, newRoomServiceCheckIsInRoomArgs, newRoomServiceCheckIsInRoomResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "room",
		"ServiceFilePath": `idl/room.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func createRoomHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*room.RoomServiceCreateRoomResult)
	success, err := handler.(room.RoomService).CreateRoom(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoomServiceCreateRoomArgs() interface{} {
	return room.NewRoomServiceCreateRoomArgs()
}

func newRoomServiceCreateRoomResult() interface{} {
	return room.NewRoomServiceCreateRoomResult()
}

func joinRoomHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*room.RoomServiceJoinRoomArgs)
	realResult := result.(*room.RoomServiceJoinRoomResult)
	success, err := handler.(room.RoomService).JoinRoom(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoomServiceJoinRoomArgs() interface{} {
	return room.NewRoomServiceJoinRoomArgs()
}

func newRoomServiceJoinRoomResult() interface{} {
	return room.NewRoomServiceJoinRoomResult()
}

func getRoomMembersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*room.RoomServiceGetRoomMembersArgs)
	realResult := result.(*room.RoomServiceGetRoomMembersResult)
	success, err := handler.(room.RoomService).GetRoomMembers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoomServiceGetRoomMembersArgs() interface{} {
	return room.NewRoomServiceGetRoomMembersArgs()
}

func newRoomServiceGetRoomMembersResult() interface{} {
	return room.NewRoomServiceGetRoomMembersResult()
}

func checkIsInRoomHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*room.RoomServiceCheckIsInRoomArgs)
	realResult := result.(*room.RoomServiceCheckIsInRoomResult)
	success, err := handler.(room.RoomService).CheckIsInRoom(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRoomServiceCheckIsInRoomArgs() interface{} {
	return room.NewRoomServiceCheckIsInRoomArgs()
}

func newRoomServiceCheckIsInRoomResult() interface{} {
	return room.NewRoomServiceCheckIsInRoomResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateRoom(ctx context.Context) (r *room.CreateRoomResp, err error) {
	var _args room.RoomServiceCreateRoomArgs
	var _result room.RoomServiceCreateRoomResult
	if err = p.c.Call(ctx, "CreateRoom", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) JoinRoom(ctx context.Context, req *room.JoinRoomReq) (r *room.JoinRoomResp, err error) {
	var _args room.RoomServiceJoinRoomArgs
	_args.Req = req
	var _result room.RoomServiceJoinRoomResult
	if err = p.c.Call(ctx, "JoinRoom", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetRoomMembers(ctx context.Context, req *room.GetRoomMembersReq) (r *room.GetRoomMembersResp, err error) {
	var _args room.RoomServiceGetRoomMembersArgs
	_args.Req = req
	var _result room.RoomServiceGetRoomMembersResult
	if err = p.c.Call(ctx, "GetRoomMembers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckIsInRoom(ctx context.Context, req *room.CheckIsInRoomReq) (r *room.CheckIsInRoomResp, err error) {
	var _args room.RoomServiceCheckIsInRoomArgs
	_args.Req = req
	var _result room.RoomServiceCheckIsInRoomResult
	if err = p.c.Call(ctx, "CheckIsInRoom", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
