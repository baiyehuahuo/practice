package main

import (
	"AStream-go/entity"
	"fmt"
	"sort"
)

func startPlayback(downloader *entity.DashDownloader, domain string, segmentDuration, segmentCount int) {
	bitrates := make([]int, 0, len(downloader.Video))
	for bitrate := range downloader.Video {
		bitrates = append(bitrates, bitrate)
	}
	sort.Ints(bitrates)

	player := NewDashPlayer(downloader.PlaybackDuration, segmentDuration, bitrates)
	go play(player)
	dpList := make([][]string, segmentCount)
	for i := 0; i < segmentCount; i++ {
		for _, bitrate := range bitrates {
			dpList[i] = append(dpList[i], downloader.Video[bitrate].URLList[i])
		}
	}

	var previousDownloadTime, previousDownloadSize []float64
	for {
		fmt.Println("hahaha", previousDownloadTime, previousDownloadSize)
	}
}
