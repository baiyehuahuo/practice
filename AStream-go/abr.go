package main

import (
	"AStream-go/config"
	"AStream-go/consts"
	"AStream-go/utils"
	"time"
)

func ABRSelect(dp *DashPlayer, abrType string) (segment int, layer int, state float64) {
	switch abrType {
	case "BLFirst":
		return BLFirstSelect(dp)
	case "DDLFirst":
		return DDLFirstSelect(dp)
	case "BackFilling":
		return BackFillingSelect(dp)
	default:
		utils.Fatalf("%s %s Unknown ABR Select Method: %s", consts.DashPlayerError, utils.GetCallerName(), abrType)
		return
	}
}

func BLFirstSelect(dp *DashPlayer) (segment int, layer int, state float64) {
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
		for judgeSegmentInRange(dp, segment) && dp.BufferQueue[segment+1][layer] {
			segment++
		}
		if judgeSegmentInRange(dp, segment) && !dp.BufferQueue[segment+1][layer] {
			break
		}
		layer++
	}

	state = getState(dp, segment, layer)

	return segment, layer, state
}

func DDLFirstSelect(dp *DashPlayer) (segment int, layer int, state float64) {
	dp.BufferLock.Lock()
	defer dp.BufferLock.Unlock()
	state = 0 // 0 is normal download  > 0 is sleep time < 0 is out
	segment, layer = dp.NextSegmentNumber, 0
	bitrateLength := len(dp.Bitrates)

	for judgeSegmentInRange(dp, segment) {
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

	state = getState(dp, segment, layer)

	return segment, layer, state
}

func BackFillingSelect(dp *DashPlayer) (segment int, layer int, state float64) {
	dp.BufferLock.Lock()
	defer dp.BufferLock.Unlock()
	state = 0 // 0 is normal download  > 0 is sleep time < 0 is out
	segment, layer = dp.NextSegmentNumber, 0
	bitrateLength := len(dp.Bitrates)

	for judgeSegmentInRange(dp, segment) && dp.BufferQueue[segment+1][0] {
		segment++
	}
	if judgeSegmentInRange(dp, segment) && !dp.BufferQueue[segment+1][0] {
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

	state = getState(dp, segment, layer)

	return segment, layer, state
}

func judgeSegmentInRange(dp *DashPlayer, segment int) bool {
	return segment+1 < len(dp.BufferQueue) && segment-(dp.NextSegmentNumber-1) <= config.NormalThreshold
}

func getState(dp *DashPlayer, segment, layer int) (state float64) {
	dp.PlaybackStateLock.Lock()
	defer dp.PlaybackStateLock.Unlock()
	if dp.PlaybackState == "END" || !judgeSegmentInRange(dp, dp.NextSegmentNumber) {
		state = -1
	} else if layer == len(dp.Bitrates) || !judgeSegmentInRange(dp, segment) {
		dp.FutureLock.Lock()
		defer dp.FutureLock.Unlock()
		state = utils.MaxFloat64(dp.Future.Sub(time.Now()).Seconds(), 0.001)
	}
	return state
}
