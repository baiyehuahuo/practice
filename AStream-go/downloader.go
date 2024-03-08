package main

import (
	"AStream-go/entity"
	"AStream-go/proxy"
	"fmt"
	"sort"
	"time"
)

func startPlayback(downloader *entity.DashDownloader, domain string, segmentDuration, segmentCount int, abrType string) {
	bitrates := make([]int, 0, len(downloader.Video))
	for bitrate := range downloader.Video {
		bitrates = append(bitrates, bitrate)
	}
	sort.Ints(bitrates)

	player := NewDashPlayer(downloader.PlaybackDuration, segmentDuration, bitrates)
	go player.Start()
	dpList := make([][]string, segmentCount)
	for i := 0; i < segmentCount; i++ {
		for _, bitrate := range bitrates {
			dpList[i] = append(dpList[i], downloader.Video[bitrate].URLList[i])
		}
	}

	proxy.SynDownload(domain + downloader.Video[bitrates[0]].Initialization)
	player.Write(-1, 0)
	var previousDownloadTime []float64
	var previousDownloadSize []int64
	var segmentNumber, layer int
	var state float64
	for {
		segmentNumber, layer, state = ABRSelect(player, abrType)
		fmt.Println(segmentNumber, layer, state)
		if state == -1 {
			break
		} else if state > 0 {
			time.Sleep(time.Duration(float64(time.Second) * state))
			continue
		}
		startTime := time.Now()
		url := domain + dpList[segmentNumber][layer]
		downloadSize := proxy.SynDownload(url)

		previousDownloadTime = append(previousDownloadTime, time.Now().Sub(startTime).Seconds())
		previousDownloadSize = append(previousDownloadSize, downloadSize)
		player.Write(segmentNumber, layer)
	}
	sleepTime := player.TotalRemain(segmentNumber)
	fmt.Printf("remain time: %v, and sleep 10s\n", sleepTime)
	time.Sleep(sleepTime + time.Second*10)
	proxy.CloseConnection()
}
