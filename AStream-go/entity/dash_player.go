package entity

import (
	"AStream-go/config"
	"AStream-go/consts"
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

func (dp *DashPlayer) PlayerRouting() {

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

func (dp *DashPlayer) ABRSelect(abrType string) (segment int, layer int, state float64) {
	switch abrType {
	case "BLFirst":
		return dp.BLFirstSelect()
	case "DDLFirst":
		return dp.DDLFirstSelect()
	case "BackFilling":
		return dp.BackFillingSelect()
	default:
		utils.Fatalf("%s Unknown ABR Select Method: %s", consts.DashPlayerError, abrType)
		return
	}
}

func (dp *DashPlayer) BLFirstSelect() (segment int, layer int, state float64) {
	dp.BufferLock.Lock()
	defer dp.BufferLock.Unlock()
	state = 0 // 0 is normal download  > 0 is sleep time < 0 is out
	segment, layer = dp.NextSegmentNumber, 0
	bitrateLength := len(dp.Bitrates)

	for layer < bitrateLength {
		segment = dp.NextSegmentNumber
		if segment == -1 && dp.BufferQueue[0][0] {
			segment = 0
		}
		for dp.judgeSegmentInRange(segment) && dp.BufferQueue[segment+1][layer] {
			segment++
		}
		if dp.judgeSegmentInRange(segment) && !dp.BufferQueue[segment+1][layer] {
			break
		}
		layer++
	}

	dp.PlaybackStateLock.Lock()
	defer dp.PlaybackStateLock.Unlock()
	if dp.PlaybackState == "END" || !dp.judgeSegmentInRange(dp.NextSegmentNumber) {
		state = -1
	} else if layer == bitrateLength {
		dp.FutureLock.Lock()
		defer dp.FutureLock.Unlock()
		state = max(dp.Future.Sub(time.Now()).Seconds(), 0.001)
	}
	return segment, layer, state
}

func (dp *DashPlayer) DDLFirstSelect() (segment int, layer int, state float64) {
	dp.BufferLock.Lock()
	defer dp.BufferLock.Unlock()
	state = 0 // 0 is normal download  > 0 is sleep time < 0 is out
	segment, layer = dp.NextSegmentNumber, 0
	bitrateLength := len(dp.Bitrates)

	for dp.judgeSegmentInRange(segment) {
		if segment == -1 && dp.BufferQueue[0][0] {
			segment = 0
			continue
		}

		for layer = 0; layer < bitrateLength && dp.BufferQueue[segment+1][layer]; layer++ {
		}

		if layer < bitrateLength && !dp.BufferQueue[segment+1][layer] {
			break
		}

		segment++
	}

	dp.PlaybackStateLock.Lock()
	defer dp.PlaybackStateLock.Unlock()
	if dp.PlaybackState == "END" || !dp.judgeSegmentInRange(dp.NextSegmentNumber) {
		state = -1
	} else if dp.judgeBufferOut(segment) {
		dp.FutureLock.Lock()
		defer dp.FutureLock.Unlock()
		state = max(dp.Future.Sub(time.Now()).Seconds(), 0.001)
	}
	return segment, layer, state
}

func (dp *DashPlayer) BackFillingSelect() (segment int, layer int, state float64) {
	dp.BufferLock.Lock()
	defer dp.BufferLock.Unlock()
	state = 0 // 0 is normal download  > 0 is sleep time < 0 is out
	segment, layer = dp.NextSegmentNumber, 0
	bitrateLength := len(dp.Bitrates)

	for dp.judgeSegmentInRange(segment) && dp.BufferQueue[segment+1][0] {
		segment++
	}
	if dp.judgeSegmentInRange(segment) && !dp.BufferQueue[segment+1][0] {
		return segment, 0, 0
	}

	maxSegment := segment
	layer = 1
	for layer < bitrateLength {
		segment = maxSegment - 1
		for segment >= dp.NextSegmentNumber && dp.BufferQueue[segment+1][layer] {
			segment--
		}
		if segment >= dp.NextSegmentNumber && !dp.BufferQueue[segment+1][layer] {
			break
		}
		layer++
	}

	dp.PlaybackStateLock.Lock()
	defer dp.PlaybackStateLock.Unlock()
	if dp.PlaybackState == "END" || !dp.judgeSegmentInRange(dp.NextSegmentNumber) {
		state = -1
	} else if layer == bitrateLength {
		dp.FutureLock.Lock()
		defer dp.FutureLock.Unlock()
		state = max(dp.Future.Sub(time.Now()).Seconds(), 0.001)
	}

	return segment, layer, state
}

func (dp *DashPlayer) judgeSegmentInRange(segment int) bool {
	return segment+1 < len(dp.BufferQueue) && !dp.judgeBufferOut(segment)
}

func (dp *DashPlayer) judgeBufferOut(segment int) bool {
	return segment-(dp.NextSegmentNumber-1) > config.NormalThreshold
}

func (dp *DashPlayer) SegmentRemain() time.Duration {
	dp.FutureLock.Lock()
	remain := max(dp.Future.Sub(time.Now()), 0)
	dp.FutureLock.Unlock()
	return remain
}

func (dp *DashPlayer) TotalRemain(segmentNumber int) time.Duration {
	dp.BufferLock.Lock()
	dp.FutureLock.Lock()
	remain := dp.SegmentRemain() + time.Duration(segmentNumber-dp.NextSegmentNumber-1)*dp.SegmentDuration
	dp.FutureLock.Unlock()
	dp.BufferLock.Unlock()
	return remain
}

func (dp *DashPlayer) Write(segment, layer int) {
	if dp.ActualStartTime.IsZero() {
		dp.ActualStartTime = time.Now()
		utils.SetJsonHandleSecondValue("playback_info", "start_time", dp.ActualStartTime)
	}

	dp.BufferLock.Lock()
	dp.BufferQueue[segment+1][layer] = true
	if layer == 0 {
		dp.BufferQSize++
	}
	dp.BufferLock.Unlock()
}

func (dp *DashPlayer) Start() {
	dp.SetState("INITIAL_BUFFERING")
	utils.Info("Starting the player")
	go dp.PlayerRouting()
	dp.LogEntry("Starting")
}

func (dp *DashPlayer) Stop() {
	dp.SetState("STOP")
	dp.LogEntry("Stopped")
	utils.Info("Stopped the playback")
}

func (dp *DashPlayer) LogEntry(action string) {
	var logTime time.Duration
	if !dp.ActualStartTime.IsZero() {
		logTime = time.Now().Sub(dp.ActualStartTime)
	}
	dp.PlaybackStateLock.Lock()
	utils.Infof("BufferStats: EpochTime=%s, CurrentPlaybackTime=%s, CurrentBufferSize=%s, CurrentPlaybackState=%s, Action=%s", logTime.Seconds(), dp.PlaybackTimer.Time().Seconds(), dp.BufferSize(), dp.PlaybackState, action)
	dp.PlaybackStateLock.Unlock()
}
