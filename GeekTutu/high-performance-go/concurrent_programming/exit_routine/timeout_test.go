package exit_routine

import (
	"runtime"
	"testing"
	"time"
)

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		_ = timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func testWithBuffer(t *testing.T, f func(chan bool)) {
	for i := 0; i < 1000; i++ {
		_ = timeoutWithBuffer(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestBadTimeout(t *testing.T) {
	test(t, doBadthing)
}

func TestGoodTimeout(t *testing.T) {
	test(t, doGoodthing)
}

func TestBadTimeoutWithBuffer(t *testing.T) {
	testWithBuffer(t, doBadthing)
}

func Test2phasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_ = timeoutFirstPhase()
	}
	time.Sleep(time.Second * 3)
	t.Log(runtime.NumGoroutine())
}
