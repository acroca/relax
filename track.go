package relax

import (
	"io/ioutil"
	"math"

	"github.com/mccoyst/vorbis"
)

type Track struct {
	SampleRate int
	Slider     *VolumeSlider
	data       []int16
	curIdx     int
	playing    bool
}

func TrackFromOGGData(raw []byte) *Track {
	data, _, sr, err := vorbis.Decode(raw)
	if err != nil {
		panic(err)
	}

	track := &Track{
		data:       data,
		Slider:     NewVolumeSlider(),
		SampleRate: sr,
		playing:    true,
	}
	return track
}
func TrackFromOGG(filename string) *Track {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return TrackFromOGGData(raw)
}

func (track *Track) PlayPause() {
	track.playing = !track.playing
}
func (track *Track) Stream(samples [][2]float64) (n int, ok bool) {
	if !track.playing {
		for i := range samples {
			samples[i] = [2]float64{}
		}
		return len(samples), true
	}
	var myIdx int
	var val float64
	for i := range samples {
		myIdx = track.curIdx + (i * 2)
		if myIdx >= len(track.data) {
			// loop
			myIdx -= len(track.data)
		}
		val = track.Slider.Val()
		samples[i][0] = val * (float64(track.data[myIdx]) / float64(math.MaxInt16))
		samples[i][1] = val * (float64(track.data[(myIdx)+1]) / float64(math.MaxInt16))
	}
	track.curIdx += len(samples) * 2
	if track.curIdx > len(track.data) {
		track.curIdx -= len(track.data)
	}
	return len(samples), true
}
func (track *Track) Err() error {
	return nil
}
