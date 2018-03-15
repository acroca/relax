package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/acroca/relax"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/gobuffalo/packr"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AppKit -framework IOKit
#import <Foundation/Foundation.h>


#import "media.h"

*/
import "C"

//export play_pause
func play_pause() {
	playPauseChn <- true
}

var quitChn = make(chan bool)
var playPauseChn = make(chan bool)

func main() {
	runtime.LockOSThread()
	go run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			handleQuit()
		}
	}()
	C.StartApp()
}

func handleQuit() {
	quitChn <- true
	os.Exit(0)
}

func run() {
	box := packr.NewBox("../tracks")

	fmt.Println("Loading audio tracks...")
	trackFiles := box.List()
	tracks := make([]*relax.Track, len(trackFiles))
	sliders := make(relax.VolumeSliders, len(trackFiles))
	for i, trackFile := range trackFiles {
		tracks[i] = relax.TrackFromOGGData(box.Bytes(trackFile))
		sliders[i] = tracks[i].Slider
	}
	go sliders.Start()
	sampleRate := beep.SampleRate(tracks[0].SampleRate)
	speaker.Init(sampleRate, sampleRate.N(time.Second/10))

	for _, track := range tracks {
		speaker.Play(track)
	}
	fmt.Println("Playing")
	for {
		select {
		case <-playPauseChn:
			for _, track := range tracks {
				track.PlayPause()
			}
		case <-quitChn:
			fmt.Println("Bye")
			return
		}
	}
}
