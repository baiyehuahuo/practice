package codec

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
)

type JsonCodec struct {
	conn io.Closer     // conn 是由构建函数传入，通常是通过 TCP 或者 Unix 建立 socket 时得到的链接实例
	buf  *bufio.Writer // buf 是为了防止阻塞而创建的带缓冲的 Writer，一般这么做能提升性能。
	dec  *json.Decoder // dec 和 enc 对应 gob 的 Decoder 和 Encoder
	enc  *json.Encoder
}

var _ Codec = (*JsonCodec)(nil)

func (c JsonCodec) Close() error {
	return c.conn.Close()
}

func (c JsonCodec) ReadHeader(header *Header) error {
	return c.dec.Decode(header)
}

func (c JsonCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c JsonCodec) Write(header *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err = c.enc.Encode(header); err != nil {
		log.Println("rpc codec: gob error encoding header: ", err)
		return
	}
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body: ", err)
		return
	}
	return nil
}

func NewJsonCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &JsonCodec{
		conn: conn,
		buf:  buf,
		dec:  json.NewDecoder(conn),
		enc:  json.NewEncoder(buf),
	}
}
