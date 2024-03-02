package read_mpd

import (
	"AStream-go/consts"
	"AStream-go/entity"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func ParseMPD(path string) *entity.MPD {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%s read failed: %s", consts.MPDErrorTag, err.Error())
	}

	mpd := new(entity.MPD)
	err = xml.Unmarshal(content, mpd)
	if err != nil {
		log.Fatalf("%s parse failed: %s", consts.MPDErrorTag, err.Error())
	}

	fmt.Println(len(mpd.Periods[0].AdaptationSet[0].Representation))
	for _, representation := range mpd.Periods[0].AdaptationSet[0].Representation {
		fmt.Println(len(representation.SegmentList.SegmentURL),
			representation.SegmentList.SegmentURL[0].Media)
	}
	return mpd
}

func ReadMPD(playback *entity.DashPlayback, mpd *entity.MPD) {

}
