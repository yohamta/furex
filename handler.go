package furex

import "github.com/hajimehoshi/ebiten/v2"

// TouchHandler represents a touch handler
type TouchHandler interface {
	HandleJustPressedTouchID(touchID ebiten.TouchID) bool
	HandleJustReleasedTouchID(touchID ebiten.TouchID)
}

// MouseHandler represents a mouse handler
type MouseHandler interface {
	HandleMouse(x, y int) bool
}
