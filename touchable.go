package furex

type TouchableComponent interface {
	Component

	HandleTouch(touchID int) bool
}
