package main

import (
	"fmt"
	"sync"
)

var (
	x     int
	mutex sync.Mutex
)

func AddWithLock(wg *sync.WaitGroup) {
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	for i := 0; i < 2000; i++ {
		x++
	}
}

func AddWithoutLock(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2000; i++ {
		x++
	}
}

func main() {
	var wg = &sync.WaitGroup{}
	x = 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go AddWithLock(wg)
	}
	wg.Wait()
	fmt.Println("after add with lock, x = ", x)
	x = 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go AddWithoutLock(wg)
	}
	wg.Wait()
	fmt.Println("after add without lock, x = ", x)
}
