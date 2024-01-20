package sync_pool_reuse

import (
	"bytes"
	"encoding/json"
	"sync"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{
	Name:   "Geektutu",
	Age:    25,
	Remark: [1024]byte{},
})

func unmarshal() {
	stu := &Student{}
	_ = json.Unmarshal(buf, stu)
}

var studentPool = sync.Pool{New: func() interface{} {
	return new(Student)
}}

var bufferPool = sync.Pool{New: func() interface{} {
	return &bytes.Buffer{}
}}

var data = make([]byte, 10000)
