package read_mpd

import (
	"AStream-go/config"
	"AStream-go/consts"
	"AStream-go/entity"
	"AStream-go/utils"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getPlaybackTime(playbackTimeStr string) float64 {
	for _, spl := range "PTHMS" {
		playbackTimeStr = strings.ReplaceAll(playbackTimeStr, string(spl), " ")
	}
	numbers := strings.Split(playbackTimeStr, " ")
	slow, fast, length := 0, 0, len(numbers)
	for fast < length {
		if numbers[fast] != "" {
			numbers[slow] = numbers[fast]
			slow++
		}
		fast++
	}
	numbers = numbers[:slow]
	for slow, fast = 0, len(numbers)-1; slow < fast; slow, fast = slow+1, fast-1 {
		numbers[slow], numbers[fast] = numbers[fast], numbers[slow]
	}
	var ans, base float64
	base = 1.0
	for _, v := range numbers {
		vFloat, err := strconv.ParseFloat(v, 64)
		if err != nil {
			utils.Fatalf("%s: getPlaybackTime failed: %s", consts.MPDErrorTag, err.Error())
			return 0
		}
		ans += vFloat * base
		base *= 60
	}
	return ans
}

func ParseMPD(path string) *entity.MPD {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%s read failed: %s", consts.MPDErrorTag, err.Error())
	}

	mpd := new(entity.MPD)
	err = xml.Unmarshal(content, mpd)
	if err != nil {
		log.Fatalf("%s parse failed: %s", consts.MPDErrorTag, err.Error())
	}

	fmt.Println(len(mpd.Periods[0].AdaptationSet[0].Representation))
	fmt.Println(mpd.Periods[0].AdaptationSet[0].SegmentBase.SegmentInitization.SourceURL)
	for _, representation := range mpd.Periods[0].AdaptationSet[0].Representation {
		fmt.Println(len(representation.SegmentList.SegmentURL), representation.SegmentList.SegmentURL[0].Media)
	}
	return mpd
}

func ReadMPD(downloader *entity.DashDownloader, mpd *entity.MPD) (segmentDuration, segmentCount int) {
	utils.Info("Reading the MPD file")

	downloader.PlaybackDuration = getPlaybackTime(mpd.MediaPresentationDuration)
	utils.SetJsonHandleMultiValue([]string{"video_metadata", "playback_duration"}, downloader.PlaybackDuration)
	downloader.MinBufferTime = getPlaybackTime(mpd.MinBufferTime)

	var moMap map[int]*entity.MediaObject
	var bandwidthList []int
	var bandwidth int
	for _, adaptationSet := range mpd.Periods[0].AdaptationSet {
		for _, representation := range adaptationSet.Representation {
			mediaFound := false
			if strings.Contains(representation.MimeType, "audio") {
				moMap = downloader.Audio
				mediaFound = true
				utils.Info("Found Audio")
			} else if strings.Contains(representation.MimeType, "video") {
				moMap = downloader.Video
				mediaFound = true
				utils.Info("Found Video")
			}
			if mediaFound {
				utils.Info("Retrieving Media")
			}
			bandwidth = representation.BandWidth
			bandwidthList = append(bandwidthList, bandwidth)
			moMap[bandwidth] = &entity.MediaObject{}
			moMap[bandwidth].SegmentSize = []float64{}
			moMap[bandwidth].Start = representation.StartWithSap
			moMap[bandwidth].TimeScale = representation.SegmentList.Timescale
			moMap[bandwidth].Duration = representation.SegmentList.Duration
			moMap[bandwidth].BaseURL = mpd.BaseURL
			moMap[bandwidth].Initialization = adaptationSet.SegmentBase.SegmentInitization.SourceURL
			for _, segmentInfo := range representation.SegmentList.SegmentURL {
				moMap[bandwidth].URLList = append(moMap[bandwidth].URLList, segmentInfo.Media)
			}
			segmentDuration = representation.SegmentList.Duration / representation.SegmentList.Timescale
			segmentCount = len(moMap[bandwidth].URLList)
		}
		if config.Limit > config.InitialBufferingCount && config.Limit < len(moMap[bandwidth].URLList) {
			downloader.PlaybackDuration = downloader.MinBufferTime * float64(config.Limit)
			utils.SetJsonHandleMultiValue([]string{"video_metadata", "playback_duration"}, downloader.PlaybackDuration)
			moMap[bandwidth].URLList = moMap[bandwidth].URLList[:config.Limit]
			segmentCount = config.Limit
		}
	}
	utils.SetJsonHandleMultiValue([]string{"video_metadata", "available_bitrates"}, bandwidthList)
	return segmentDuration, segmentCount
}
