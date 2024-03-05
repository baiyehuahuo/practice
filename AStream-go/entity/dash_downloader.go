package entity

type DashDownloader struct {
	MinBufferTime    float64
	PlaybackDuration float64
	Audio            map[int]*MediaObject
	Video            map[int]*MediaObject
}

type MediaObject struct {
	MinBufferTime  float64
	Start          int
	TimeScale      int
	Duration       int
	Initialization string
	BaseURL        string
	SegmentSize    []float64
	URLList        []string
}
