package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CaiyunRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type Wiki struct {
}
type Prons struct {
}
type Dictionary struct {
	Entry        string        `json:"entry"`
	Explanations []string      `json:"explanations"`
	Related      []interface{} `json:"related"`
	Source       string        `json:"source"`
	Prons        Prons         `json:"prons"`
	Type         string        `json:"type"`
}

type CaiyunResponse struct {
	Rc         int        `json:"rc"`
	Wiki       Wiki       `json:"wiki"`
	Dictionary Dictionary `json:"dictionary"`
}

func caiyunxiaoyi() {
	client := &http.Client{}
	//var data = strings.NewReader(`{"trans_type":"zh2en","source":"good"}`) // 转换成流 (转成流式为了防止大文件直接读入内存)
	request := CaiyunRequest{
		TransType: "en2zh",
		Source:    "你好",
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	data := bytes.NewBuffer(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text[表情]ain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "9066d2e3d3b4d60ee16713ae8a5368a7")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	response := CaiyunResponse{}
	if err = json.Unmarshal(bodyText, &response); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", response) // 加上#能看到仔细结果
	fmt.Println(request.Source, "UK: ")
	for _, item := range response.Dictionary.Explanations {
		fmt.Println(item)
	}
}
