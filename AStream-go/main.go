package main

import (
	"AStream-go/config"
	"AStream-go/consts"
	"AStream-go/mpd"
	"AStream-go/proxy"
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
	"sync"
	"time"
)

var (
	jsonHandle      map[string]interface{}
	jsonHandleMutex = sync.Mutex{}
	logger          *logrus.Logger
)

func init() {
	jsonHandle = make(map[string]interface{})
	jsonHandle["playback_info"] = map[string]interface{}{
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

/*
def get_mpd(url):
    """ 下载mpd文件并设置初始延迟
        Module to download the MPD from the URL and save it to file
    """
    start = time.time()
    glueConnection.syn_download(url)
    mpd_file = config_dash.SVC_FILE_PATH + url.split('/')[-1]
    config_dash.LOG.info("Downloaded the MPD file {}".format(mpd_file))
    config_dash.JSON_HANDLE['mpd_delay'] = time.time() - start
    return mpd_file
*/

func getMPD(mpdURL string) *mpd.MPD {
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
	logger.Infof("Downloaded the MPD file %s", mpdURL)
	jsonHandleMutex.Lock()
	jsonHandle["mpd_delay"] = time.Now().Sub(start)
	jsonHandleMutex.Unlock()
	return mpd.ParseMPD(config.DownloadPath + filepath.Base(mpdURL))
}

func getDomain(mpdURL string) string {
	u, err := url.Parse(mpdURL)
	if err != nil {
		log.Fatalf("%s MPD url parse error", consts.MainError)
	}
	return fmt.Sprintf("%s://%s/", u.Scheme, u.Host)
}

func configureLogFile(logFilePath string) {
	// 设置日志路径
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	// todo set formatter better
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:03:04",
	})

	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Fatalf("%s open log file failed: ", err)
	}
	//defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		logger.Fatalf("%s truncate log file failed: ", err)
	}

	logger.SetOutput(file)
}

/*
def main():
 	// finished
    parser = ArgumentParser(description='Process Client parameters')
    create_arguments(parser)
    args = parser.parse_args()
    globals().update(vars(args))

	// finished
    configure_log_file(playback_type=PLAYBACK.lower(), log_file=None)
    clean_files(config_dash.SVC_FILE_PATH)

	// finished
    if not MPD:
        print("ERROR: Please provide the URL to the MPD file. Try Again..")
        return None

	// finished
    test.version = args.VERSION
    print('Selected scheduler is : {}'.format(test.version))
    glueConnection.setupPM()
    config_dash.LOG.info('Downloading MPD file %s' % MPD)

    # Retrieve the MPD files for the video
    mpd_file = get_mpd(MPD)
    domain = get_domain_name(MPD)
    config_dash.SERVER_DOMAIN = domain
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

	configureLogFile(config.LogFilename)
	utils.CleanFiles(config.DownloadPath)

	if mpdURL == nil {
		log.Fatalf("%s MPD url is empty", consts.MainError)
	}

	config.InitConfig(*version)
	logger.Infof("Selected scheduler is : %s", config.Version)
	proxy.ClientSetup()
	logger.Infof("Downloading MPD file %s", *mpdURL)

	_ = getMPD(*mpdURL)
	domain := getDomain(*mpdURL)
	config.ServerDomain = domain
}
