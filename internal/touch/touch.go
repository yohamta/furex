package touch

import "github.com/hajimehoshi/ebiten/v2"

type TouchPosition struct {
	X, Y int
}

var (
	touchPositions = make(map[ebiten.TouchID]TouchPosition)
)

func RecordTouchPosition(t ebiten.TouchID, x, y int) {
	touchPositions[t] = TouchPosition{x, y}
}

func LastTouchPosition(t ebiten.TouchID) *TouchPosition {
	s, ok := touchPositions[t]
	if ok {
		return &s
	}
	return &TouchPosition{0, 0}
}
