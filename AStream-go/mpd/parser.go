package mpd

import (
	"AStream-go/consts"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func ParseMPD(path string) *MPD {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%s read failed: %s", consts.MPDErrorTag, err.Error())
	}

	mpd := new(MPD)
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
