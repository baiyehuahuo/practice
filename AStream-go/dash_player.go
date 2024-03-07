package main

import (
	"AStream-go/config"
	"AStream-go/consts"
	"AStream-go/entity"
	"AStream-go/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	playerState map[string]bool
)

func init() {
	playerState = make(map[string]bool, 7)
	for _, state := range []string{"INITIALIZED", "PLAY", "PAUSE", "BUFFERING", "STOP", "END"} {
		playerState[state] = true
	}
}

//        self.player_thread = None

type DashPlayer struct {
	PlaybackDuration  float64
	SegmentDuration   time.Duration
	Bitrates          []int
	PlaybackTimer     *entity.StopWatch
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

func NewDashPlayer(videoLength float64, segmentDuration int, bitrates []int) (player *DashPlayer) {
	utils.Info("Initializing the Buffer")
	player = &DashPlayer{
		PlaybackDuration: videoLength,
		SegmentDuration:  time.Duration(segmentDuration) * time.Second,
		Bitrates:         bitrates,

		// Timers to keep track of playback time and the actual time
		PlaybackTimer:   &entity.StopWatch{},
		ActualStartTime: time.Time{},

		// Playback State
		PlaybackState:     "INITIALIZED",
		PlaybackStateLock: sync.Mutex{},

		// Current video buffer that holds the segment data
		PlaybackCount: int(math.Ceil(videoLength/float64(segmentDuration))) + 1, // add an initialization
		BufferQSize:   0,
		BufferQueue:   nil,
		BufferLock:    sync.Mutex{},

		NextSegmentNumber: -1,
		Interruption:      0,
		BufferLogFile:     config.BufferLogFilename,
		Future:            time.Time{},
		FutureLock:        sync.Mutex{},

		DisplayLayer: nil,
	}
	player.BufferQueue = make([][]bool, player.PlaybackCount)
	player.DisplayLayer = make([]int, player.PlaybackCount)
	highestLayer := len(bitrates)
	for i := range player.BufferQueue {
		player.BufferQueue[i] = make([]bool, highestLayer) // initial is 0, seg 0 is 1, seg 1 is 2 ...
	}

	utils.Infof("VideoLength=%v,segmentDuration=%v,segmentCount=%d", player.PlaybackDuration, player.SegmentDuration.Seconds(), player.PlaybackCount)
	return player
}

func (dp *DashPlayer) PlayerRouting() {
	var (
		startTime           = time.Now()
		interruptionStart   time.Time
		initialWait         float64
		paused, buffering   bool
		overallDownloadSize int64
		overallDownloadTime float64
		layerCount          = make([]int, len(dp.Bitrates))
		overall             int
	)

	utils.Warnf("Initialized player with video length %v", dp.PlaybackDuration)
	for {
		switch dp.PlaybackState {
		case "END":
			// Video stopped by the user
			utils.Warnf("Finished playback of the video: %v seconds of video played for %v seconds", dp.PlaybackDuration, time.Now().Sub(startTime).Seconds())
			utils.SetJsonHandleMultiValue([]string{"playback_info", "end_time"}, time.Now())
			dp.PlaybackTimer.Pause()
			return

		case "STOP":
			//# If video is stopped quit updating the playback time and exit player
			utils.Infof("Player Stopped at time %v", time.Now().Sub(startTime).Seconds())
			utils.SetJsonHandleMultiValue([]string{"playback_info", "end_time"}, time.Now())
			dp.PlaybackTimer.Pause()
			dp.LogEntry("Stopped")
			return

		case "PAUSE":
			if !paused {
				//# do not update the playback time. Wait for the state to change
				utils.Infof("Player Paused after %4.2f seconds of playback", dp.PlaybackTimer.Time().Seconds())
				dp.PlaybackTimer.Pause()
				paused = true
			}
			continue

		case "BUFFERING":
			if !buffering {
				utils.Warnf("Entering buffering stage after %.2f seconds of playback", dp.PlaybackTimer.Time().Seconds())
				dp.PlaybackTimer.Pause()
				buffering = true
				interruptionStart = time.Now()
				utils.SetJsonHandleMultiValueIntIncrease([]string{"playback_info", "interruptions", "count"})
			} else {
				// If the size of the buffer is greater than the RE_BUFFERING_DURATION then start playback
				// If the RE_BUFFERING_DURATION is greater than remain length of the video then do not wait
				remainPlaybackTime := dp.PlaybackDuration - dp.PlaybackTimer.Time().Seconds()
				if (dp.BufferSize() >= config.RebufferingCount) || (float64(config.RebufferingCount)*dp.SegmentDuration.Seconds() >= remainPlaybackTime && dp.BufferSize() > 0) {
					buffering = false
					if !interruptionStart.IsZero() {
						interruptionEnd := time.Now()
						interruption := interruptionEnd.Sub(interruptionStart).Seconds()
						dp.Interruption = interruption

						utils.SetJsonHandleMultiValueSliceAppend([]string{"playback_info", "interruptions", "events"}, []time.Time{interruptionStart, interruptionEnd})
						utils.SetJsonHandleMultiValueFloatAdd([]string{"playback_info", "interruptions", "total_duration"}, interruption)
						utils.Infof("segment_number: %d \t interruption: %.2f", dp.NextSegmentNumber, dp.Interruption)
						interruptionStart = time.Time{}
					}
					dp.SetState("PLAY")
					dp.LogEntry("Buffering-Play")
				}
			}

		case "INITIALIZED":
			if dp.BufferSize() < config.InitialBufferingCount {
				initialWait = time.Now().Sub(startTime).Seconds()
				continue
			} else {
				utils.Infof("Started playing at %v", time.Now())
				utils.Infof("Initial Waiting Time = %.2f", initialWait)
				utils.SetJsonHandleMultiValue([]string{"playback_info", "initial_buffering_duration"}, initialWait)
				utils.SetJsonHandleMultiValue([]string{"playback_info", "start_time"}, time.Now())
				fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! initial time: %.2fs\n", initialWait)
				dp.SetState("PLAY")
				dp.LogEntry("InitialBuffering-Play")
			}

		case "PLAY":
			// Check of the buffer has any segments
			bufferSize := dp.BufferSize()
			if bufferSize == 0 {
				// 没加载出下一个BL层，卡住
				utils.Infof("Buffer empty after %.2f seconds of playback", dp.PlaybackTimer.Time().Seconds())
				dp.PlaybackTimer.Pause()
				dp.SetState("BUFFERING")
				dp.LogEntry("Play-Buffering")
				continue
			}
			playSegment := dp.BufferGet()
			segmentNumber := playSegment["segment_number"].(int)
			utils.Warn("***************PLAYING*****************")
			utils.Warnf("Reading the segment number %d from the buffer at playtime %.2f", segmentNumber, dp.PlaybackTimer.Time().Seconds())
			dp.LogEntry("StillPlaying")
			// Start the playback
			dp.PlaybackTimer.Start()

			layerNo := 0
			var segmentSize int64
			if segmentNumber != -1 {
				layerNo = playSegment["display_layer"].(int)
				dp.DisplayLayer[segmentNumber+1] = layerNo
				var fi os.FileInfo
				var err error
				for i := 0; i <= layerNo; i++ {
					layerCount[i]++
					fi, err = os.Stat(fmt.Sprintf("DownloadedSegment/BBB-I-720p.seg%d-L%d.svc", segmentNumber, i))
					if err != nil {
						utils.Errorf("%s get file %s stats error: %s", consts.DashPlayerError, fmt.Sprintf("DownloadedSegment/BBB-I-720p.seg%d-L%d.svc", segmentNumber, i), err.Error())
						continue
					}
					segmentSize += fi.Size()
				}
			}

			overall++
			overallDownloadSize += segmentSize
			fmt.Printf("display segment %d, layer %d\n", segmentNumber, layerNo)

			// If playback hasn't started yet, wait for the playback_start_time
			time.Sleep(dp.SegmentRemain())

			// Duration for which the video was played in seconds, make sure all segments have been played!
			if dp.PlaybackTimer.Time().Seconds() >= dp.PlaybackDuration && dp.PlaybackCount == 0 {
				dp.PlaybackTimer.Pause()
				deleteFilePath := config.DownloadPath

				overallDownloadTime = dp.PlaybackDuration
				totalDownloaded := utils.CountSuffixFileSize(config.DownloadPath, ".svc") // overall_download_size
				utils.CleanFiles(config.DownloadPath)
				totalDownloadTime := overallDownloadTime

				meanDownloadRate := strconv.FormatFloat(math.Round(float64(overallDownloadSize)/(totalDownloadTime*1024)), 'f', 2, 64) + " / " +
					strconv.FormatFloat(math.Round(float64(totalDownloaded)/(totalDownloadTime*1024)), 'f', 2, 64)
				utils.SetJsonHandleMultiValue([]string{"playback_info", "mean_downrate"}, meanDownloadRate)
				utils.SetJsonHandleMultiValue([]string{"playback_info", "mean_downtime"}, totalDownloadTime/float64(overall))
				for i := range layerCount {
					utils.SetJsonHandleMultiValue([]string{"playback_info", fmt.Sprintf("Layer%d_count", i)}, layerCount[i])
				}

				utils.Warnf("Completed the video playback: %.2f seconds", dp.PlaybackTimer.Time().Seconds())
				utils.Warnf("Mean downrate: %v MBits", float64(totalDownloaded)/(totalDownloadTime*1024))
				utils.Warnf("Interruptions: %v", utils.GetJsonHandleMultiValue([]string{"playback_info", "interruptions", "count"}))
				utils.Warnf("Interruption time total: %v s", utils.GetJsonHandleMultiValue([]string{"playback_info", "interruptions", "total_duration"}))
				for i := range layerCount {
					utils.Warnf("Layer%d total count: %d ", i, layerCount[i])
				}
				utils.Warnf("total count: %d ", overall)

				fmt.Printf("total buffer time %v, bandwidth %v\n", utils.GetJsonHandleMultiValue([]string{"playback_info", "interruptions", "total_duration"}), utils.GetJsonHandleMultiValue([]string{"playback_info", "mean_downrate"}))
				utils.DeleteFiles(deleteFilePath, ".svc")
				utils.DeleteFiles(deleteFilePath, "init")
				dp.SetState("END") // 播放结束
				dp.LogEntry("TheEnd")
			}
		default:
			utils.Fatalf("%s Unknown player state: %s", consts.DashPlayerError, dp.PlaybackState)
		}
	}
}

func (dp *DashPlayer) SetState(state string) {
	state = strings.ToUpper(state)
	if playerState[state] {
		dp.PlaybackStateLock.Lock()
		utils.Infof("Changing state from %s to %s at %s Playback time ", dp.PlaybackState, state, dp.PlaybackTimer.Time())
		dp.PlaybackState = state
		dp.PlaybackStateLock.Unlock()
	} else {
		utils.Errorf("%s Unidentified state: %s", consts.DashPlayerError, state)
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
		utils.Fatalf("%s %s Unknown ABR Select Method: %s", consts.DashPlayerError, utils.GetCallerName(), abrType)
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
	} else if layer == bitrateLength || !dp.judgeSegmentInRange(segment) {
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
	} else if layer == bitrateLength || !dp.judgeSegmentInRange(segment) {
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
	} else if layer == bitrateLength || !dp.judgeSegmentInRange(segment) {
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
	remain := max(dp.Future.Sub(time.Now()), 0) + time.Duration(segmentNumber-dp.NextSegmentNumber-1)*dp.SegmentDuration
	dp.FutureLock.Unlock()
	dp.BufferLock.Unlock()
	return max(remain, 0)
}

func (dp *DashPlayer) Write(segment, layer int) {
	if dp.ActualStartTime.IsZero() {
		dp.ActualStartTime = time.Now()
		utils.SetJsonHandleMultiValue([]string{"playback_info", "start_time"}, dp.ActualStartTime)
	}

	dp.BufferLock.Lock()
	dp.BufferQueue[segment+1][layer] = true
	if layer == 0 {
		dp.BufferQSize++
	}
	dp.BufferLock.Unlock()
}

func (dp *DashPlayer) Start() {
	dp.SetState("INITIALIZED")
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
	utils.Infof("BufferStats: EpochTime=%v, CurrentPlaybackTime=%v, CurrentBufferSize=%v, CurrentPlaybackState=%s, Action=%s", logTime.Seconds(), dp.PlaybackTimer.Time().Seconds(), dp.BufferSize(), dp.PlaybackState, action)
	dp.PlaybackStateLock.Unlock()
}
