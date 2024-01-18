package GeeRPC

import (
	"encoding/json"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicNumebr = 0x3bef5c
const serverPrefix = "rpc server: "

type Option struct {
	MagicNumber int        // marks this is a geerpc request
	CodecType   codec.Type // choose different Codec
}

var DefaultOption = &Option{
	MagicNumber: MagicNumebr,
	CodecType:   codec.GobType,
}

// stores all information of a call
type request struct {
	header       *codec.Header // header of request
	argv, replyv reflect.Value // argv and replyv of request
}

// Server represents an RPC Server
type Server struct{}

// NewServer returns a new Server
func NewServer() *Server {
	return &Server{}
}

// DefaultServer is the default instance of *Server
var DefaultServer = &Server{}

// Accept accepts connections on the listener and serves requests for each incoming connection
func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

// Accept accepts connections on the listener and serves requests for each incoming connection
func (s *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept() // 等待 socket 建立 开启子协程处理
		if err != nil {
			log.Println(serverPrefix, "accept error: ", err)
			return
		}
		go s.ServeConn(conn)
	}
}

// ServeConn runs the server on a single connection
func (s *Server) ServeConn(conn net.Conn) {
	defer conn.Close()

	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println(serverPrefix, "options error: ", err)
		return
	}
	if opt.MagicNumber != MagicNumebr {
		log.Println(serverPrefix, "invalid magic number: ", opt.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Println(serverPrefix, "invalid codec type: ", opt.CodecType)
		return
	}
	log.Println(serverPrefix, "serve conn: ", conn.RemoteAddr())
	s.serveCodec(f(conn))
}

// a placeholder for response body when error occurs
var invalidRequest = struct{}{}

func (s *Server) serveCodec(cc codec.Codec) {
	// 读取 处理 回复
	sending := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for {
		req, err := s.readRequest(cc)
		if err != nil {
			if req == nil {
				break // it's not possible to recover, so close.
			}
			req.header.Error = err.Error()
			s.sendResponse(cc, req.header, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		go s.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	cc.Close()
}

func (s *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var header codec.Header
	if err := cc.ReadHeader(&header); err != nil {
		if err != io.EOF || err != io.ErrUnexpectedEOF {
			log.Println(serverPrefix, "read header error: ", err)
		}
		return nil, err
	}
	return &header, nil
}

func (s *Server) readRequest(cc codec.Codec) (*request, error) {
	header, err := s.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{
		header: header,
	}
	// TODO: fill request
	req.argv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println(serverPrefix, "read argv err: ", err)
	}
	return req, nil
}

func (s *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println(serverPrefix, "write response error: ", err)
	}
}

func (s *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println(req.header, req.argv.Elem())
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc response %d", req.header.Seq))
	s.sendResponse(cc, req.header, req.replyv.Interface(), sending)
}
