package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Component represents a UI component that can be added to a Flex container.
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
	// Size returns the size(x,y) of the component.
	Size() image.Point

	// Position returns the position(x,y) relative to it's parent container.
	Position() image.Point
}

// ButtonComponent represents a button component.
type Button interface {
	// HandlePress handle the event when user just started pressing the button
	// The parameter (x, y) is the location relative to the window (0,0).
	HandlePress(x, y int)

	// HandleRelease handle the event when user just released the button.
	// The parameter (x, y) is the location relative to the window (0,0).
	// The parameter isCancel is true when the touch/left click is released outside of the button.
	HandleRelease(x, y int, isCancel bool)
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
	// The parameter (x, y) is the location relative to the window (0,0).
	HandleMouse(x, y int) bool
}

// MouseLeftClickHandler represents a component that handle mouse button left click.
type MouseLeftClickHandler interface {
	// HandleJustPressedMouseButtonLeft handle left mouse button click just pressed.
	// The parameter (x, y) is the location relative to the window (0,0).
	// It returns true if it handles the mouse move.
	HandleJustPressedMouseButtonLeft(x, y int) bool
	// HandleJustReleasedTouchID handles the touchID just released.
	// The parameter (x, y) is the location relative to the window (0,0).
	HandleJustReleasedMouseButtonLeft(x, y int)
}
