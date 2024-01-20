package exit_routine

import (
	"runtime"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasks()
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
}

func TestDoCheckClose(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasksCheckClose()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}
