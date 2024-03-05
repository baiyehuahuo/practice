package entity

type DashPlayback struct {
	MinBufferTime    float64
	PlaybackDuration float64
	Audio            map[int]*MediaObject
	Video            map[int]*MediaObject
}

type MediaObject struct {
	MinBufferTime float64
	Start         int
	TimeScale     int
	//SegmentDuration int
	Initialization string
	BaseURL        string
	SegmentSize    []float64
	URLList        []string
}

//class MediaObject(object):
//""" 处理音频和视频流的对象
//Object to handle audio and video stream
//"""
//def __init__(self):
//self.min_buffer_time = None
//self.start = None
//self.timescale = None
//self.segment_duration = None
//self.initialization = None
//self.base_url = None
//self.url_list = list()
