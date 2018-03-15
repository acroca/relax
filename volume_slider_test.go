package relax

import (
	"fmt"
	"testing"
)

func TestVolumeSliders(t *testing.T) {
	original := VolumeSliders{
		NewVolumeSlider(),
		NewVolumeSlider(),
		NewVolumeSlider(),
		NewVolumeSlider(),
		NewVolumeSlider(),
		NewVolumeSlider(),
	}
	original.SetTargets()
	for _, vs := range original {
		fmt.Printf("===> %[2]v: %[1]v\n", vs.target, `vs.target`)
	}

}
