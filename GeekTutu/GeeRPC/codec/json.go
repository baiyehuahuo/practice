package codec

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
)

var _ Codec = (*JsonCodec)(nil)

type JsonCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *json.Decoder
	enc  *json.Encoder
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

func (j *JsonCodec) Close() error {
	return j.conn.Close()
}

func (j *JsonCodec) ReadHeader(header *Header) (err error) {
	return j.dec.Decode(header)
}

func (j *JsonCodec) ReadBody(body interface{}) (err error) {
	return j.dec.Decode(body)
}

func (j *JsonCodec) Write(header *Header, body interface{}) (err error) {
	defer func() {
		j.buf.Flush()
		if err != nil {
			j.Close()
		}
	}()
	if err = j.enc.Encode(header); err != nil {
		log.Println(codecPrefix, "json encoding header error: ", err)
		return err
	}
	if err = j.enc.Encode(body); err != nil {
		log.Println(codecPrefix, "json encoding body error: ", err)
		return err
	}
	return nil
}
