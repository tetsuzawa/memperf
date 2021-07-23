package main

import (
	"encoding/json"
	"github.com/tetsuzawa/memperf"
	"log"
	"math/rand"
	"os"
)

const (
	samples      = 50000
	maxId        = 65536
	maxNumBlocks = 100
)

func main() {
	rand.Seed(1)
	frames1 := memperf.Frames{Items: make([]memperf.Frame, samples)}
	for i := 0; i < samples; i++ {
		bcc := make([]int64, rand.Int63n(maxNumBlocks))
		for j := range bcc {
			bcc[j] = rand.Int63n(maxId)
		}
		bct := make([]int64, rand.Int63n(maxNumBlocks))
		for j := range bct {
			bct[j] = rand.Int63n(maxId)
		}
		frame := memperf.Frame{
			Id:                      int64(i + 1),
			MediumId:                rand.Int63n(maxId),
			BlockCampaignCategories: bcc,
			BlockCreativeTags:       bct,
		}
		frames1.Items[i] = frame
	}

	frames2 := memperf.Frames{}
	frames2.Items = make([]memperf.Frame, len(frames1.Items))
	copy(frames2.Items, frames1.Items)

	for i := 0; i < samples; i += 4 {
		bcc := make([]int64, rand.Int63n(maxNumBlocks))
		for j := range bcc {
			bcc[j] = rand.Int63n(maxId)
		}
		bct := make([]int64, rand.Int63n(maxNumBlocks))
		for j := range bct {
			bct[j] = rand.Int63n(maxId)
		}
		frame := memperf.Frame{
			Id:                      int64(i + 1),
			MediumId:                rand.Int63n(maxId),
			BlockCampaignCategories: bcc,
			BlockCreativeTags:       bct,
		}
		frames2.Items[i] = frame
	}
	if err := saveFrames("frames_1.json", &frames1); err != nil {
		log.Fatalln(err)
	}
	if err := saveFrames("frames_2.json", &frames2); err != nil {
		log.Fatalln(err)
	}
}

func saveFrames(fname string, frames *memperf.Frames) error {
	f, err := os.Create(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	b, err := json.Marshal(frames)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = f.Write(b)
	return err
}
