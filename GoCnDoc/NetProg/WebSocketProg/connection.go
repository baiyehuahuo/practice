package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type connection struct {
	conn        *websocket.Conn
	messageChan chan []byte
	data        *Data
}

var (
	// Upgrader 指定用于将 HTTP 连接升级为 WebSocket 连接的参数。 线程安全
	wu = &websocket.Upgrader{
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	userList []string
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 升级将 HTTP 服务器连接升级为 WebSocket 协议。
	// responseHeader 包含在对客户端升级请求的响应中。
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	conn := &connection{
		messageChan: make(chan []byte, 256),
		conn:        ws,
		data:        &Data{},
	}
	h.r <- conn
	go conn.writer()
	conn.reader()
	defer func() {
		conn.data.Type = "logout"
		userList = del(userList, conn.data.User)
		conn.data.UserList = userList
		conn.data.Content = conn.data.User
		message, err := json.Marshal(conn.data)
		if err != nil {
			log.Println("marshal c.data failed:", conn.data, err)
		}
		h.messageChan <- message
		h.r <- conn
	}()
}

func (conn *connection) writer() {
	for message := range conn.messageChan {
		conn.conn.WriteMessage(websocket.TextMessage, message)
	}
	conn.conn.Close()
}

func (conn *connection) reader() {
	var message []byte
	for {
		_, connData, err := conn.conn.ReadMessage()
		if err != nil {
			h.r <- conn
			break
		}
		if err = json.Unmarshal(connData, &conn.data); err != nil {
			h.r <- conn
			break
		}
		switch conn.data.Type {
		case "login":
			conn.data.User = conn.data.Content
			conn.data.From = conn.data.User
			userList = append(userList, conn.data.User)
			conn.data.UserList = userList
			message, _ = json.Marshal(conn.data)
			h.messageChan <- message
		case "user":
			message, _ = json.Marshal(conn.data)
			h.messageChan <- message
		case "logout":
			userList = del(userList, conn.data.User)
			message, _ = json.Marshal(conn.data)
			h.messageChan <- message
			h.r <- conn
		default:
			log.Println("unknown message type: ", conn.data.Type)
		}
	}
}

// 从用户队列中删除 用户 delUser
func del(userList []string, delUser string) []string {
	length := len(userList)
	for i := 0; i < length; i++ {
		if userList[i] == delUser {
			return append(userList[:i], userList[i+1:]...)
		}
	}
	return userList
}
