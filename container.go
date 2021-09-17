package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Container represents a container that can have child components
type Container interface {
	Component

	SetFrame(image.Rectangle)
	AddChild(child Component)
}

type Child struct {
	bounds          image.Rectangle
	component       Component
	IsButtonPressed bool
	handledTouchID  ebiten.TouchID
}

type ContainerEmbed struct {
	children []*Child
	isDirty  bool
	frame    image.Rectangle
}

// SetFrame sets the location (x,y) and size (width,height) relative to the window (0,0).
func (cont *ContainerEmbed) SetFrame(frame image.Rectangle) {
	cont.frame = frame
	cont.isDirty = true
}

func (cont *ContainerEmbed) AddChild(child Component) {
	c := &Child{component: child, handledTouchID: -1}
	cont.children = append(cont.children, c)
	cont.isDirty = true
}

func (cont *ContainerEmbed) ChildBounds(child Component) *image.Rectangle {
	for i := 0; i < len(cont.children); i++ {
		c := cont.children[i]
		if c.component == child {
			return &c.bounds
		}
	}
	return nil
}

func (cont *ContainerEmbed) HandleJustPressedTouchID(touchID ebiten.TouchID) bool {
	result := false
	x, y := ebiten.TouchPosition(touchID)
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		handler, ok := child.component.(TouchHandler)
		if ok && handler != nil {
			if result == false && isInside(child.bounds, x, y) {
				if handler.HandleJustPressedTouchID(touchID) {
					child.handledTouchID = touchID
					result = true
					break
				}
			}
		}

		button, ok := child.component.(ButtonComponent)
		if ok && button != nil {
			if result == false && isInside(child.bounds, x, y) {
				if child.IsButtonPressed == false {
					child.IsButtonPressed = true
					child.handledTouchID = touchID
					button.HandlePress(touchID)
				}
				result = true
			} else if child.handledTouchID == touchID {
				child.handledTouchID = -1
			}
		}
	}
	return result
}

func (cont *ContainerEmbed) HandleJustReleasedTouchID(touchID ebiten.TouchID) {
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		handler, ok := child.component.(TouchHandler)
		if ok && handler != nil {
			if child.handledTouchID == touchID {
				handler.HandleJustReleasedTouchID(touchID)
				child.handledTouchID = -1
			}
		}

		button, ok := child.component.(ButtonComponent)
		if ok && button != nil {
			if child.handledTouchID == touchID {
				if child.IsButtonPressed == true {
					child.IsButtonPressed = false
					child.handledTouchID = -1
					x, y := lastTouchPosition(touchID)
					if x == 0 && y == 0 {
						button.HandleRelease(touchID, true)
					} else {
						button.HandleRelease(touchID, isInside(child.bounds, x, y))
					}
				}
			}
		}
	}
}

func (cont *ContainerEmbed) HandleMouse(x, y int) bool {
	result := false
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		handler, ok := child.component.(MouseHandler)
		if ok && handler != nil {
			if result == false && isInside(child.bounds, x, y) {
				if handler.HandleMouse(x, y) {
					result = true
				}
			}
		}

		button, ok := child.component.(ButtonComponent)
		if ok && button != nil && child.handledTouchID == -1 {
			if result == false && isInside(child.bounds, x, y) {
				if child.IsButtonPressed {
					if inpututil.IsMouseButtonJustReleased((ebiten.MouseButtonLeft)) {
						button.HandleRelease(-1, isInside(child.bounds, x, y))
						child.IsButtonPressed = false
					}
				} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					button.HandlePress(-1)
					child.IsButtonPressed = true
				}
				result = true
			} else {
				if child.IsButtonPressed {
					if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
						button.HandleRelease(-1, isInside(child.bounds, x, y))
						child.IsButtonPressed = false
					}
				}
			}
		}
	}
	return result
}

func isInside(r image.Rectangle, x, y int) bool {
	return r.Min.X <= x && x <= r.Max.X && r.Min.Y <= y && y <= r.Max.Y
}
