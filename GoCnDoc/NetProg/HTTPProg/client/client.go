package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println("connect server failed: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("get err: ", err)
			break
		}
		fmt.Println("Get message: ", string(buf[:n]))
		if err == io.EOF {
			break
		}
	}
}
