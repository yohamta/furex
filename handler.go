package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// DrawHandler represents a component that can be added to a container.
type DrawHandler interface {
	// HandleDraw function draws the content of the component inside the frame.
	// The frame parameter represents the location (x,y) and size (width,height) relative to the window (0,0).
	HandleDraw(screen *ebiten.Image, frame image.Rectangle)
}

// UpdateHandler represents a component that updates by one tick.
type UpdateHandler interface {
	// Updater updates the state of the component by one tick.
	HandleUpdate()
}

// ButtonHandler represents a button component.
type ButtonHandler interface {
	// HandlePress handle the event when user just started pressing the button
	// The parameter (x, y) is the location relative to the window (0,0).
	// touchID is the unique ID of the touch.
	// If the button is pressed by a mouse, touchID is -1.
	HandlePress(x, y int, t ebiten.TouchID)

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

// MouseLeftButtonHandler represents a component that handle mouse button left click.
type MouseLeftButtonHandler interface {
	// HandleJustPressedMouseButtonLeft handle left mouse button click just pressed.
	// The parameter (x, y) is the location relative to the window (0,0).
	// It returns true if it handles the mouse move.
	HandleJustPressedMouseButtonLeft(x, y int) bool
	// HandleJustReleasedTouchID handles the touchID just released.
	// The parameter (x, y) is the location relative to the window (0,0).
	HandleJustReleasedMouseButtonLeft(x, y int)
}

// SwipeHandler represents different swipe directions.
type SwipeDirection int

const (
	SwipeDirectionLeft SwipeDirection = iota
	SwipeDirectionRight
	SwipeDirectionUp
	SwipeDirectionDown
)

// SwipeHandler represents a component that handle swipe.
type SwipeHandler interface {
	// HandleSwipe handle swipe. It returns true
	// if it handles the swipe event
	HandleSwipe(dir SwipeDirection)
}
