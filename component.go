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

// FixedSizeComponent represents a component with fixed size.
type FixedSizeComponent interface {
	// Size returns the size(x,y) of the component.
	Size() image.Point
}

// AbsolutePositionComponent represents a component with fixed size.
type AbsolutePositionComponent interface {
	FixedSizeComponent

	// Position returns the position(x,y) relative to it's parent container.
	Position() image.Point
}

// ButtonComponent represents a button component.
type Button interface {
	Component

	// OnPressButton will be called when the button is pressed.
	HandlePress()

	// OnReleaseButton will be called when the button is released.
	HandleRelease(isInside bool)
}

// TouchHandler represents a component that handle touches.
type TouchHandler interface {
	// HandleJustPressedTouchID handles the touchID just pressed and returns true if it handles the TouchID
	HandleJustPressedTouchID(touch ebiten.TouchID, x, y int) bool
	// HandleJustReleasedTouchID handles the touchID just released
	// Should be called only when it handled the TouchID when pressed
	HandleJustReleasedTouchID(touch ebiten.TouchID, x, y int)
}

// MouseHandler represents a component that handle mouse move.
type MouseHandler interface {
	// HandleMouse handles the mouch move and returns true if it handle the mouse move.
	HandleMouse(x, y int) bool
}

// MouseLeftClickHandler represents a component that handle mouse button left click.
type MouseLeftClickHandler interface {
	// HandleJustPressedTouchID handles the touchID just pressed and returns true if it handles the TouchID.
	HandleJustPressedMouseButtonLeft(x, y int) bool
	// HandleJustReleasedTouchID handles the touchID just released.
	// Should be called only when it handled the TouchID when pressed.
	HandleJustReleasedMouseButtonLeft(x, y int)
}
