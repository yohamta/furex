package furex

type Touchable interface {
	HandleTouch(touchID int) bool
}
