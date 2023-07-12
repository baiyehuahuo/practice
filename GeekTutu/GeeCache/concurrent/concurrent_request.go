package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	url := "http://localhost:9999/api?key=Tom"
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			buffer := &bytes.Buffer{}
			io.Copy(buffer, resp.Body)
			fmt.Printf(buffer.String())
			wg.Done()
		}()
	}
	wg.Wait()
}
