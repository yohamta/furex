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
	Component

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

// TouchHandler represents a touch handler
type TouchHandler interface {
	Component

	// HandleJustPressedTouchID handles the touchID just pressed and returns true if it handles the TouchID
	HandleJustPressedTouchID(touchID ebiten.TouchID) bool
	// HandleJustReleasedTouchID handles the touchID just released
	// Should be called only when it handled the TouchID when pressed
	HandleJustReleasedTouchID(touchID ebiten.TouchID)
}

// MouseHandler represents a mouse handler
type MouseHandler interface {
	Component

	// HandleMouse handles the TouchID when it is pressed
	HandleMouse(x, y int) bool
}

// ButtonComponent represents a button
type Button interface {
	Component

	// OnPressButton will be called when the button is pressed
	HandlePress(t ebiten.TouchID)

	// OnReleaseButton will be called when the button is released
	HandleRelease(t ebiten.TouchID, isInside bool)
}
