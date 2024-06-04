package models

import (
	"fmt"
	"ginchat/utils"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"log"
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

func Chat(w *http.ResponseWriter, request *http.Request) {
	// 1. 获取参数并校验合法性
	query := request.URL.Query()
	userIDStr := query.Get("userIDStr")
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
	}).Upgrade(*w, request, nil)
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
	// 6. 接收逻辑
}
