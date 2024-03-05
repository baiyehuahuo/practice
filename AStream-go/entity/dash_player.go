package entity

import (
	"AStream-go/utils"
	"strings"
	"sync"
	"time"
)

var (
	playerState map[string]bool
	exitState   map[string]bool
)

func init() {
	playerState = make(map[string]bool, 7)
	exitState = make(map[string]bool, 2)
	for _, state := range []string{"INITIALIZED", "INITIAL_BUFFERING", "PLAY", "PAUSE", "BUFFERING", "STOP", "END"} {
		playerState[state] = true
	}
	for _, state := range []string{"STOP", "END"} {
		exitState[state] = true
	}
}

//        self.player_thread = None

type DashPlayer struct {
	PlaybackDuration  float64
	SegmentDuration   time.Duration
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

func (dp *DashPlayer) SetState(state string) {
	state = strings.ToUpper(state)
	if playerState[state] {
		dp.PlaybackStateLock.Lock()
		utils.Infof("Changing state from %s to %s at %s Playback time ", dp.PlaybackState, state, dp.PlaybackTimer.Time())
		dp.PlaybackState = state
		dp.PlaybackStateLock.Unlock()
	} else {
		utils.Fatalf("Unidentified state: %s", state)
	}
}

func (dp *DashPlayer) BufferSize() int {
	dp.BufferLock.Lock()
	ans := dp.BufferQSize
	dp.BufferLock.Unlock()
	return ans
}

/*
   # todo fwf fill it
   play_segment = {
       'playback_length': 2,
       'bitrate': self.bitrates[display_layer],
       'URI': "fwf_empty",
       'segment_number': self.next_segment_number,
       'download_time': 0.0,
       'display_layer': display_layer,
   }
*/

func (dp *DashPlayer) BufferGet() (playSegment map[string]interface{}) {
	// Read one the segment from the buffer
	// Acquire Lock on the buffer and read a segment for it
	dp.BufferLock.Lock()
	dp.FutureLock.Lock() // Calculate time playback when the segment finishes
	displayLayer := -1   // initial is -1
	for displayLayer+1 < len(dp.Bitrates) && dp.BufferQueue[dp.NextSegmentNumber+1][displayLayer+1] {
		displayLayer++
	}
	// todo fwf fill it
	playSegment = map[string]interface{}{
		"playback_length": dp.SegmentDuration.Seconds(),
		"bitrate":         dp.Bitrates[displayLayer],
		"segment_number":  dp.NextSegmentNumber,
		"display_layer":   displayLayer,
		"URI":             nil, // todo fwf remove it?
		"download_time":   0.0, // todo fwf remove it?
	}

	dp.Future = time.Now().Add(dp.SegmentDuration)

	dp.BufferQSize--
	dp.PlaybackCount--
	dp.NextSegmentNumber++

	dp.FutureLock.Unlock()
	dp.BufferLock.Unlock()

	return playSegment
}
