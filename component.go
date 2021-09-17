package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Component represents a component of the UI.
type Component interface {
	// Draw function draws the content of the component inside the frame.
	// The frame parameter represents the location (x,y) and size (width,height) relative to the window (0,0).
	Draw(screen *ebiten.Image, frame image.Rectangle)
}

// UpdatableComponent represents a component that updates by one tick.
type UpdatableComponent interface {
	// Update updates the state of the component by one tick.
	Update()
}
