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
}
struct JoinRoomResp{
    1: required base.BaseResp baseResp
    2: optional string ClientId
    3: optional string clientToken
}

struct GetRoomMembersReq{
    1: required string roomId
}
struct GetRoomMembersResp{
    1: required base.BaseResp baseResp
    2: optional list<string> members
}
struct CheckIsInRoomReq{
    1: required string roomId
    2: required string clientId
}
struct CheckIsInRoomResp{
    1: required base.BaseResp baseResp
    2: optional bool isInRoom
}


service RoomService {
    CreateRoomResp CreateRoom()
    JoinRoomResp JoinRoom(JoinRoomReq req)
    GetRoomMembersResp GetRoomMembers(GetRoomMembersReq req)
    CheckIsInRoomResp CheckIsInRoom(CheckIsInRoomReq req)
}
