package main

import (
	"sync"
)

// 用户信息表，主要是存储用户名和fd的对应关系

type UserData struct {
	server    *Server //fd 对应的server结构
	timestamp int64   //登录时间戳，单位ms
}

var userServerMap = make(map[string]*UserData, 10000)
var userServerLock sync.RWMutex

func NewUserData(server *Server, timestamp int64) *UserData {
	return &UserData{
		server:    server,
		timestamp: timestamp,
	}
}

func AddUserData(userData *UserData, user string) {
	userServerLock.Lock()
	defer userServerLock.Unlock()
	if temp, ok := userServerMap[user]; ok {
		if temp.timestamp < userData.timestamp {
			userServerMap[user] = userData
			return
		}
	}
	userServerMap[user] = userData
}

func DelUserData(user string) {
	userServerLock.Lock()
	defer userServerLock.Unlock()
	delete(userServerMap, user)
}

func GetServerByUser(user string) *Server {
	userServerLock.RLock()
	defer userServerLock.RUnlock()
	if temp, ok := userServerMap[user]; ok {
		return temp.server
	}
	return nil
}

// 当前在线用户的总数
func UserCount() int {
	userServerLock.RLock()
	defer userServerLock.RUnlock()
	return len(userServerMap)
}
