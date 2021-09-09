package furex

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type tpos struct {
	x, y int
}

var (
	touchPositions = make(map[ebiten.TouchID]tpos)
)

func recordTouchPosition(t ebiten.TouchID) {
	x, y := ebiten.TouchPosition(t)
	touchPositions[t] = tpos{x, y}
}

func lastTouchPosition(t ebiten.TouchID) (int, int) {
	s, ok := touchPositions[t]
	if ok {
		return s.x, s.y
	}
	return 0, 0
}
