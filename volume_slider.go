package relax

import (
	"math"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	rnd = rand.New(rand.NewSource(time.Now().Unix()))
)

type VolumeSliders []*VolumeSlider

func (a VolumeSliders) Len() int           { return len(a) }
func (a VolumeSliders) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a VolumeSliders) Less(i, j int) bool { return rnd.Int()%2 == 0 }

type VolumeSlider struct {
	stop chan bool
	mtx  sync.Mutex

	target float64
	val    float64
	speed  float64
}

func NewVolumeSlider() *VolumeSlider {

	return &VolumeSlider{
		stop:   make(chan bool),
		speed:  0.000001,
		val:    0,
		target: 0,
	}
}

func (vs *VolumeSlider) Val() float64 {
	vs.mtx.Lock()
	defer vs.mtx.Unlock()
	if vs.foundTarget() {
		return vs.val
	}
	vs.moveTowardsTarget()
	return vs.val
}

func (vs *VolumeSlider) foundTarget() bool {
	return math.Abs(vs.val-vs.target) < 0.01
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

func (vss VolumeSliders) Start() {
	t := time.NewTicker(3 * time.Second)

	select {
	case <-t.C:
		vss.SetTargets()
	}
}
func (vss VolumeSliders) SetTargets() {
	if len(vss) == 1 {
		vss[0].target = 1
		return
	}
	sorted := make(VolumeSliders, len(vss))
	for i, vs := range vss {
		sorted[i] = vs
	}
	sort.Sort(sorted)

	if len(sorted) == 2 {
		vss[0].target = 1
		vss[2].target = 0.3
		return
	}

	// Assigns targets based on a sinus function between (Pi/2) and 0
	for idx, vs := range sorted {
		vs.target = math.Sin((math.Pi / 2) * (float64(len(sorted)-1-idx) / float64(len(sorted)-1)))
	}
}
