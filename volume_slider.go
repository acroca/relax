package relax

import (
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	rnd = rand.New(rand.NewSource(time.Now().Unix()))
)

type VolumeSlider struct {
	stop   chan bool
	ticker *time.Ticker
	mtx    sync.Mutex

	target float64
	val    float64
	speed  float64
}

func NewVolumeSlider() *VolumeSlider {
	t := time.NewTicker(10 * time.Millisecond)

	return &VolumeSlider{
		stop:   make(chan bool),
		ticker: t,
		speed:  0.000001,
		val:    0,
		target: 0,
	}
}

func (vs *VolumeSlider) Val() float64 {
	vs.mtx.Lock()
	if vs.foundTarget() {
		vs.setTarget()
	}
	vs.moveTowardsTarget()
	vs.mtx.Unlock()
	return vs.val
}

func (vs *VolumeSlider) foundTarget() bool {
	return math.Abs(vs.val-vs.target) < 0.01
}

func (vs *VolumeSlider) setTarget() {
	vs.target = rnd.Float64()
}

func (vs *VolumeSlider) moveTowardsTarget() {
	if vs.val > vs.target {
		vs.val = vs.val - vs.speed
	} else {
		vs.val = vs.val + vs.speed
	}
}

func (vs *VolumeSlider) String() string {
	str := "-----------------------------------------------------------------------------------------"
	idx := float64(len(str)) * vs.val
	return strings.Replace(str, "-", "*", int(idx))
}
