package furex

// TouchHandler represents a touch handler
type TouchHandler interface {
	Component

	HandleJustPressedTouchID(touchID int) bool
	HandleJustReleasedTouchID(touchID int) bool
}

// MouseHandler represents a mouse handler
type MouseHandler interface {
	Component

	HandleMouse(x, y int) bool
}
