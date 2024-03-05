package main

import (
	"AStream-go/config"
	"AStream-go/entity"
	"AStream-go/utils"
	"math"
	"sync"
	"time"
)

var (
	playerState map[string]bool
	exitState   map[string]bool
)

func init() {
	for _, state := range []string{"INITIALIZED", "INITIAL_BUFFERING", "PLAY", "PAUSE", "BUFFERING", "STOP", "END"} {
		playerState[state] = true
	}
	for _, state := range []string{"STOP", "END"} {
		exitState[state] = true
	}
}

//    def __init__(self, video_length, segment_duration, bitrates):
//        self.player_thread = None

func NewDashPlayer(videoLength float64, segmentDuration int, bitrates []int) (player *entity.DashPlayer) {
	utils.Info("Initializing the Buffer")
	player = &entity.DashPlayer{
		PlaybackDuration: videoLength,
		SegmentDuration:  segmentDuration,
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

	utils.Infof("VideoLength=%v,segmentDuration=%d,segmentCount=%d", player.PlaybackDuration, player.SegmentDuration, player.PlaybackCount)
	return player
}

func play(player *entity.DashPlayer) {

}
