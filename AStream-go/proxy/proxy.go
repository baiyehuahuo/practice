package proxy

import (
	"AStream-go/config"
	"io"
	_ "net/http/pprof"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"

	"net/http"
	"path/filepath"

	"time"

	"fmt"
	//quic "github.com/lucas-clemente/quic-go"
	//"github.com/lucas-clemente/quic-go/h2quic"
)

const (
	// Prefix for PROXY specific messages
	logTag      = "PROXY MODULE:"
	svcFilePath = "DownloadedSegment"
)

var (
	hcClientMutex sync.Mutex
	hcClient      *http.Client
)

func ClientSetup() {
	// f, _ := os.OpenFile("golang.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0755)
	// log.SetOutput(f)
	// log.SetFlags(0)
	// log.SetPrefix(logTag)
	// log.SetOutput(os.Stdout)

	// Accept any offered certificate chain
	// Use a HTTP/2.0 connection via QUIC
	hcClientMutex.Lock()
	hcClient = &http.Client{
		//Transport: &h2quic.RoundTripper{
		//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//	QuicConfig:      &quic.Config{CreatePaths: true},
		//},
	}
	hcClientMutex.Unlock()
}

func CloseConnection() {
	hcClientMutex.Lock()
	if hcClient != nil {
		hcClient.CloseIdleConnections()
		time.Sleep(time.Second * 10)
		hcClient = nil
	}
	hcClientMutex.Unlock()
}

func SynDownload(url string) int64 {
	filename := filepath.Base(url)
	source, _ := os.Open(path.Join("dataset/BBB", filename))
	defer func(source *os.File) { _ = source.Close() }(source)
	destination, _ := os.Create(path.Join(config.DownloadPath, filename))
	defer func(destination *os.File) { _ = destination.Close() }(destination)
	size, _ := io.Copy(destination, source)
	time.Sleep(time.Duration(float64(size) / 338743 * float64(time.Second)))
	return size
}

func SynDownloadOri(url string) int64 {
	segmentNo, layer := getSegmentInfo(url)

	// fmt.Printf(logTag+"go moudle GET %s, ddl %d\n", url, priority.Weight)
	hcClientMutex.Lock()
	rsp, err := hcClient.Get(url)
	hcClientMutex.Unlock()

	if err != nil {
		fmt.Printf(logTag+"seg%d-L%d download error: %s\n", segmentNo, layer, err)
		return 0
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(rsp.Body)

	segmentName := strings.Split(url, "/")[len(strings.Split(url, "/"))-1]
	f, err := os.Create(filepath.Join(svcFilePath, segmentName))
	if err != nil {
		fmt.Printf(logTag+"seg%d-L%d create file fail: %s\n", segmentNo, layer, err)
	}
	defer func(f *os.File) { _ = f.Close() }(f)

	received, err := io.Copy(f, rsp.Body)
	if err != nil {
		fmt.Printf(logTag+"seg%d-L%d io segment file copy error : %s\n", segmentNo, layer, err)
	}
	fmt.Printf(logTag+"seg%d-L%d body received: %d\n", segmentNo, layer, received)
	return received
}

func getSegmentInfo(segmentURL string) (int, int) {
	if strings.Contains(segmentURL, "init.svc") || strings.Contains(segmentURL, ".mpd") {
		return 0, 0
	}
	splitURL := strings.Split(segmentURL, "/")
	segmentURL = splitURL[len(splitURL)-1]
	Info := strings.Split(segmentURL, ".")[1]

	SegInfo := strings.Split(Info, "-")[0]
	segmentNo, _ := strconv.Atoi(strings.Trim(SegInfo, "seg"))

	LayInfo := strings.Split(Info, "-")[1]
	layerNo, _ := strconv.Atoi(strings.Trim(LayInfo, "L"))

	return segmentNo, layerNo
}

func getDep(segmentNo, layer int) uint32 {
	return uint32(segmentNo<<4 | layer)
}
