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
	isButtonPressed bool
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

func (cont *ContainerEmbed) childFrame(c *Child) *image.Rectangle {
	r := c.bounds.Add(cont.frame.Min)
	return &r
}

func (cont *ContainerEmbed) HandleJustPressedTouchID(touchID ebiten.TouchID, x, y int) bool {
	result := false
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		childFrame := cont.childFrame(child)
		handler, ok := child.component.(TouchHandler)
		if ok && handler != nil {
			if result == false && isInside(childFrame, x, y) {
				if handler.HandleJustPressedTouchID(touchID, x, y) {
					child.handledTouchID = touchID
					result = true
					break
				}
			}
		}

		button, ok := child.component.(Button)
		if ok && button != nil {
			if result == false && isInside(childFrame, x, y) {
				if child.isButtonPressed == false {
					child.isButtonPressed = true
					child.handledTouchID = touchID
					button.HandlePress()
				}
				result = true
			} else if child.handledTouchID == touchID {
				child.handledTouchID = -1
			}
		}
	}
	return result
}

func (cont *ContainerEmbed) HandleJustReleasedTouchID(touchID ebiten.TouchID, x, y int) {
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		handler, ok := child.component.(TouchHandler)
		if ok && handler != nil {
			if child.handledTouchID == touchID {
				handler.HandleJustReleasedTouchID(touchID, x, y)
				child.handledTouchID = -1
			}
		}

		button, ok := child.component.(Button)
		if ok && button != nil {
			if child.handledTouchID == touchID {
				if child.isButtonPressed == true {
					child.isButtonPressed = false
					child.handledTouchID = -1
					if x == 0 && y == 0 {
						button.HandleRelease(true)
					} else {
						button.HandleRelease(isInside(cont.childFrame(child), x, y))
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
		childFrame := cont.childFrame(child)
		handler, ok := child.component.(MouseHandler)
		if ok && handler != nil {
			if result == false && isInside(childFrame, x, y) {
				if handler.HandleMouse(x, y) {
					result = true
				}
			}
		}

		button, ok := child.component.(Button)
		if ok && button != nil && child.handledTouchID == -1 {
			if result == false && isInside(childFrame, x, y) {
				if child.isButtonPressed {
					if inpututil.IsMouseButtonJustReleased((ebiten.MouseButtonLeft)) {
						button.HandleRelease(isInside(childFrame, x, y))
						child.isButtonPressed = false
					}
				} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
					button.HandlePress()
					child.isButtonPressed = true
				}
				result = true
			} else {
				if child.isButtonPressed {
					if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
						button.HandleRelease(isInside(childFrame, x, y))
						child.isButtonPressed = false
					}
				}
			}
		}
	}
	return result
}

func isInside(r *image.Rectangle, x, y int) bool {
	return r.Min.X <= x && x <= r.Max.X && r.Min.Y <= y && y <= r.Max.Y
}
