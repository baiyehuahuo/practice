// Package codec 编解码相关的包
package codec

import "io"

type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence chosen by client 区分不同的请求
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(header *Header) (err error)
	ReadBody(body interface{}) (err error)
	Write(header *Header, body interface{}) (err error)
}

// NewCodecFunc Codec 构造的函数
type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType     Type   = "application/gob"
	JsonType    Type   = "application/json"
	codecPrefix string = "rpc codec: "
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
	NewCodecFuncMap[JsonType] = NewJsonCodec
}
