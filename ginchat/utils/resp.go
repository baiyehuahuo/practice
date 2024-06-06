package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println("Marshal error: ", err)
	}
	if _, err = w.Write(ret); err != nil {
		log.Println("Write error: ", err)
	}
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, msg, nil)
}

func RespOK(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, msg, data)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println("Marshal error: ", err)
	}
	if _, err = w.Write(ret); err != nil {
		log.Println("Write error: ", err)
	}
}

func RespOKList(w http.ResponseWriter, data interface{}, total interface{}) {
	RespList(w, 0, data, total)
}
