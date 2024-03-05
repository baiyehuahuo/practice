package entity

import (
	"sync"
	"time"
)

//        self.player_thread = None

type DashPlayer struct {
	PlaybackDuration  float64
	SegmentDuration   int
	Bitrates          []int
	PlaybackTimer     *StopWatch
	ActualStartTime   time.Time
	PlaybackState     string
	PlaybackStateLock sync.Mutex
	PlaybackCount     int
	BufferQSize       int
	BufferQueue       [][]bool
	BufferLock        sync.Mutex
	NextSegmentNumber int
	Interruption      float64
	BufferLogFile     string
	Future            time.Time
	FutureLock        sync.Mutex
	DisplayLayer      []int
}
