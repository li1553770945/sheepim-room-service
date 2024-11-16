namespace go room
include "base.thrift"

struct CreateRoomResp{
    1: required base.BaseResp baseResp
    2: optional string roomId
    3: optional string clientId
    4: optional string clientToken
}

struct JoinRoomReq{
    1: required string roomId
    2: optional string clientToken
}
struct JoinRoomResp{
    1: required base.BaseResp baseResp
    2: optional string ClientId
    3: optional string clientToken
}
service RoomService {
    CreateRoomResp CreateRoom()
    JoinRoomResp JoinRoom(JoinRoomReq req)
}
