package models

import (
	"encoding/json"
	"fmt"
	"ginchat/utils"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model
	FromID  uint   // sender
	ToID    uint   // receiver
	Type    int    // message type (group or one-to-one)
	Media   int    // message type (text or picture)
	Content string // message content
	Picture string
	Desc    string
	Amount  int // other
}

func (msg *Message) TableName() string {
	return "message"
}

func AutoMigrateMessage() error {
	return utils.GetDB().AutoMigrate(&Message{})
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap = make(map[uint]*Node)
var rwLocker sync.RWMutex

func Chat(w http.ResponseWriter, request *http.Request) {
	// 1. 获取参数并校验合法性
	query := request.URL.Query()
	userIDStr := query.Get("userID")
	//token := query.Get("token")
	//targetIDStr := query.Get("targetID")
	//context := query.Get("context")
	//msgType := query.Get("type")
	isValid := true // todo check token
	conn, err := (&websocket.Upgrader{
		// check token
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(w, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2. 获取 conn
	node := &Node{Conn: conn, DataQueue: make(chan []byte, 50), GroupSets: set.New(set.ThreadSafe)}
	// 3. 用户关系
	// 4. userid 跟 node 绑定并加锁
	rwLocker.Lock()
	userIDInt, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatal(err)
	}
	clientMap[uint(userIDInt)] = node
	rwLocker.Unlock()
	// 5. 发送逻辑
	go sendProc(node)
	// 6. 接收逻辑
	go recvProc(node)
	sendMsg(uint(userIDInt), []byte("欢迎加入新世界"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] receive <<<<<< ", string(data))
	}
}

var udpSendChan = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpSendChan <- data
}

func init() {
	// 完成 UDP 数据发送协程
	go func() {
		con, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IP{192, 168, 0, 255}, Port: 3000})
		if err != nil {
			log.Fatal(err)
		}
		defer con.Close()
		for {
			select {
			case data := <-udpSendChan:
				fmt.Println("[ws] udp write <<<<<< ", string(data))
				_, err = con.Write(data)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}()
	// 完成 UDP 数据接收协程
	go func() {
		con, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 3000})
		if err != nil {
			log.Fatal(err)
		}
		defer con.Close()
		for {
			var buf = make([]byte, 512)
			length, err := con.Read(buf)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println("[ws] udp receive <<<<<< ", string(buf[:length]))
			broadMsg(buf[:length])
		}
	}()
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err)
		return
	}
	switch msg.Type {
	case 1: // 私信
		// sendMsg
		sendMsg(msg.ToID, []byte(msg.Content))
	case 2: // 群发
		// sendGroupMsg
	case 3: // 广播
		// sendAllMsg
	case 4:
		// exit
	default:

	}
}

func sendMsg(ToID uint, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[ToID]
	rwLocker.RUnlock()
	if !ok {
		return
	}
	node.DataQueue <- msg
}