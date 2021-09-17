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

// FixedSizeComponent represents a component with fixed size
type FixedSizeComponent interface {
	Component

	// Size returns the size(x,y) of the container.
	Size() image.Point
}

// AbsolutePositionComponent represents a component with fixed size
type AbsolutePositionComponent interface {
	FixedSizeComponent

	// Position returns the position(x,y) relative to it's parent container.
	Position() image.Point
}
