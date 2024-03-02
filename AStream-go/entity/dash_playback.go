package entity

type DashPlayback struct {
	MinBufferTime    int
	PlaybackDuration int
	Audio            map[string]interface{}
	Video            map[string]interface{}
}
