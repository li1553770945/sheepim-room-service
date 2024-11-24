package repo

import (
	"context"
	"errors"
	"time"
)

func (r *Repository) CreateRoom(roomId string, initClientId string) error {
	roomKey := "room:" + roomId + ":users"

	// 使用 SAdd 确保房间存在
	// SAdd 不会重复添加成员，所以即使房间已存在也不会报错
	ctx := context.Background()
	_, err := r.Cache.SAdd(ctx, roomKey, initClientId).Result()
	if err != nil {
		return err
	}

	// 设置 Key 的过期时间（例如 24 小时）
	// 防止长期无人访问的房间占用存储空间
	r.Cache.Expire(ctx, roomKey, time.Duration(r.CacheExpireSeconds)*time.Second) // 24 小时
	return nil
}
func (r *Repository) IsRoomExist(roomId string) (exist bool, err error) {
	roomKey := "room:" + roomId + ":users"

	ctx := context.Background()

	// 确保房间存在（如果房间不存在，则先创建房间）
	exists, err := r.Cache.Exists(ctx, roomKey).Result()
	if err != nil {
		return false, err
	}
	return exists != 0, nil
}

func (r *Repository) JoinRoom(roomId, clientId string) error {
	// Redis Key 格式
	roomKey := "room:" + roomId + ":users"

	ctx := context.Background()

	// 确保房间存在（如果房间不存在，则先创建房间）
	exists, err := r.Cache.Exists(ctx, roomKey).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		// 房间不存在，创建房间
		return errors.New("不存在的房间")
	}

	// 将用户 ID 添加到房间的在线用户列表
	_, err = r.Cache.SAdd(ctx, roomKey, clientId).Result()
	return err
}

func (r *Repository) GetRoomMembers(roomId string) ([]string, error) {
	roomKey := "room:" + roomId + ":users"

	ctx := context.Background()

	// 获取房间成员列表
	members, err := r.Cache.SMembers(ctx, roomKey).Result()
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (r *Repository) CheckIsInRoom(roomId string, clientId string) (bool, error) {
	roomKey := "room:" + roomId + ":users"
	ctx := context.Background()

	isMember, err := r.Cache.SIsMember(ctx, roomKey, clientId).Result()
	if err != nil {
		return false, err
	}

	return isMember, nil
}
