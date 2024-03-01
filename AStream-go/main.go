package main

import (
	"AStream-go/config"
	"AStream-go/consts"
	"AStream-go/mpd"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	JsonHandle      map[string]interface{}
	JsonHandleMutex = sync.Mutex{}
)

func init() {
	JsonHandle = make(map[string]interface{})
	JsonHandle["playback_info"] = map[string]interface{}{
		"start_time":                 nil,
		"end_time":                   nil,
		"initial_buffering_duration": nil,
		"interruptions": map[string]interface{}{
			"count":          0,
			"events":         new([][]float64),
			"total_duration": 0,
		},
		"up_shifts":   0,
		"down_shifts": 0,
	}
}

func main() {
	mpdPath := flag.String("m", "https://10.0.1.2:6121/BBB.mpd", "MPD download path")
	version := flag.String("vr", "BLFirst", "ABR version")
	flag.Parse()
	config.InitConfig(*version)
	// todo or not (no used?)
	//configure_log_file(playback_type=PLAYBACK.lower(), log_file=None)
	fmt.Println(*mpdPath, *version)

	// todo no used
	JsonHandleMutex.Lock()
	JsonHandle["playback_type"] = "normal"
	JsonHandleMutex.Unlock()

	if mpdPath == nil {
		log.Fatalf("%s MPD url is empty", consts.MainError)
	}

	now := time.Now()
	fmt.Println(now.Format("2006-01-02.15:04:05"))

	mpd.ParseMPD("BBB.mpd")
}
