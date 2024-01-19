//package common_data_struct
//
//import "encoding/json"
//
//type hub struct {
//	c map[*connection]bool
//	b chan []byte
//	r chan *connection
//	u chan *connection
//}
//
//var h = hub{
//	c: make(map[*connection]bool),
//	u: make(chan *connection),
//	b: make(chan []byte),
//	r: make(chan *connection),
//}
//
//func (h *hub) run() {
//	for {
//		select {
//		case c := <-h.r:
//			h.c[c] = true
//			c.data.Ip = c.ws.RemoteAddr().String()
//			c.data.Type = "handshake"
//			c.data.UserList = userList
//			data, _ := json.Marshal(c.data)
//			c.sc <- data
//		case c := <-h.u:
//			if _, ok := h.c[c]; ok {
//				delete(h.c, c)
//				close(c.sc)
//			}
//		case data := <-h.b:
//			for c := range h.c {
//				select {
//				case c.sc <- data:
//				default:
//					delete(h.c, c)
//					close(c.sc)
//				}
//			}
//		}
//	}
//}

package main

import (
	"encoding/json"
)

type hub struct {
	connections map[*connection]bool
	messageChan chan []byte
	r           chan *connection
}

var h = hub{
	connections: make(map[*connection]bool),
	messageChan: make(chan []byte),      // 数据往里传，剩下的交给 web 解决
	r:           make(chan *connection), // 连接出现登录或者登出事件
}

func (h *hub) run() {
	var message []byte
	var conn *connection
	for {
		select {
		case conn = <-h.r:
			h.connections[conn] = true
			conn.data.Ip = conn.conn.RemoteAddr().String()
			conn.data.Type = "handshake"
			conn.data.UserList = userList
			message, _ = json.Marshal(conn.data)
			conn.messageChan <- message
		case message = <-h.messageChan:
			for conn = range h.connections {
				select {
				case conn.messageChan <- message:
				default:
					delete(h.connections, conn)
					close(conn.messageChan)
				}
			}
		}
	}
}
