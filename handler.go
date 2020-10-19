package furex

// TouchHandler represents a touch handler
type TouchHandler interface {
	HandleJustPressedTouchID(touchID int) bool
	HandleJustReleasedTouchID(touchID int)
}

// MouseHandler represents a mouse handler
type MouseHandler interface {
	HandleMouse(x, y int) bool
}
