package main

import (
	"AStream-go/config"
	"AStream-go/consts"
	"AStream-go/entity"
	"AStream-go/proxy"
	"AStream-go/read_mpd"
	"AStream-go/utils"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"
)

func getMPD(mpdURL string) *entity.MPD {
	// 下载mpd文件并设置初始延迟
	start := time.Now()

	// todo download really!
	// proxy.SynDownload(mpdURL)

	source, _ := os.Open("BBB.mpd")
	defer source.Close()
	destination, _ := os.Create(path.Join(config.DownloadPath, filepath.Base(mpdURL)))
	defer destination.Close()
	_, _ = io.Copy(destination, source)

	// download log info
	utils.Infof("Downloaded the MPD file %s", mpdURL)

	utils.SetJsonHandleValue("mpd_delay", time.Now().Sub(start))
	mpdPath := config.DownloadPath + filepath.Base(mpdURL)
	utils.SetJsonHandleSecondValue("video_metadata", "mpd_file", mpdPath)

	return read_mpd.ParseMPD(mpdPath)
}

func getDomain(mpdURL string) string {
	u, err := url.Parse(mpdURL)
	if err != nil {
		log.Fatalf("%s MPD url parse error", consts.MainError)
	}
	return fmt.Sprintf("%s://%s/", u.Scheme, u.Host)
}

/*
def main():
    dp_object = DashPlayback()

    # Reading the MPD file created
    dp_object, video_segment_duration = read_mpd.read_mpd(mpd_file, dp_object)
    config_dash.LOG.info("The DASH media has %d video representations" % len(dp_object.video))

    if "NORMAL" in PLAYBACK.upper():
        config_dash.LOG.critical("Started Normal-DASH Playback")
        start_playback_smart(dp_object, domain, "NORMAL", video_segment_duration)
    else:
        config_dash.LOG.error("Unknown Playback parameter {}".format(PLAYBACK))
*/

func main() {
	mpdURL := flag.String("m", "https://10.0.1.2:6121/BBB.mpd", "MPD download path")
	version := flag.String("vr", "BLFirst", "ABR version")
	flag.Parse()

	utils.ConfigureLogFile(config.LogFilename)
	utils.CleanFiles(config.DownloadPath)

	if mpdURL == nil {
		log.Fatalf("%s MPD url is empty", consts.MainError)
	}

	config.InitConfig(*version)
	utils.Infof("Selected scheduler is : %s", config.Version)
	proxy.ClientSetup()
	utils.Infof("Downloading MPD file %s", *mpdURL)

	mpd := getMPD(*mpdURL)
	domain := getDomain(*mpdURL)
	config.ServerDomain = domain

	dashPlayback := &entity.DashPlayback{Audio: make(map[int]*entity.MediaObject), Video: make(map[int]*entity.MediaObject)}
	segmentDuration := read_mpd.ReadMPD(dashPlayback, mpd)
	logrus.New().Infof("The DASH media has %d video representations, segment duration is %v", len(dashPlayback.Video), segmentDuration)
	utils.Infof("The DASH media has %d video representations, segment duration is %v", len(dashPlayback.Video), segmentDuration)

	utils.Warn("Started Normal-DASH Playback")
	utils.SaveJsonHandle(config.JsonABRLogPath)
}
