package entity

type DashDownloader struct {
	MinBufferTime    float64 // 2s
	PlaybackDuration float64 // 400s
	Audio            map[int]*MediaObject
	Video            map[int]*MediaObject
}

type MediaObject struct {
	MinBufferTime  float64 // 2s
	Start          int
	TimeScale      int // 24
	Duration       int // 48
	Initialization string
	BaseURL        string
	SegmentSize    []float64
	URLList        []string
}
