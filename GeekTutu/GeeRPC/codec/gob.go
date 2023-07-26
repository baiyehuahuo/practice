package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

var _ Codec = (*GobCodec)(nil)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (g *GobCodec) Close() (err error) {
	return g.conn.Close()
}

func (g *GobCodec) ReadHeader(header *Header) (err error) {
	return g.dec.Decode(header)
}

func (g *GobCodec) ReadBody(body interface{}) (err error) {
	return g.dec.Decode(body)
}

func (g *GobCodec) Write(header *Header, body interface{}) (err error) {
	defer func() {
		g.buf.Flush()
		if err != nil {
			g.Close()
		}
	}()

	if err = g.enc.Encode(header); err != nil {
		log.Println(codecPrefix, "gob error encoding header: ", err)
		return
	}
	if err = g.enc.Encode(body); err != nil {
		log.Println(codecPrefix, "gob error encoding body: ", err)
		return
	}
	return nil
}
