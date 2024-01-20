package exit_routine

import (
	"fmt"
	"time"
)

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}

func do2phases(phase1, done chan bool) {
	time.Sleep(time.Second)
	select {
	case phase1 <- true:
	default:
		return
	}
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Microsecond):
		return fmt.Errorf("timeout")
	}
}

func timeoutWithBuffer(f func(chan bool)) error {
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Microsecond):
		return fmt.Errorf("timeout")
	}
}

func timeoutFirstPhase() error {
	phase1 := make(chan bool)
	done := make(chan bool)
	go do2phases(phase1, done)
	select {
	case <-phase1:
		<-done
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}
