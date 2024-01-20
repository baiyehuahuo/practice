package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"log"
	"math"
	"sync"
	"time"
)

func wasteGoroutine() {
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}

func useChan() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}

func thirdPackage() {
	var wg sync.WaitGroup
	pool := tunny.NewFunc(3, func(i interface{}) interface{} {
		defer wg.Done()
		log.Println(i)
		time.Sleep(time.Second)
		return nil
	})
	defer pool.Close()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go pool.Process(i)
	}
	wg.Wait()
}

func main() {
	thirdPackage()
}
