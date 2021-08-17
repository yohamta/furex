package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Container represents a container that can have child components
type Container interface {
	Component

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
}

func (f *Flex) Draw(screen *ebiten.Image, frame image.Rectangle) {
	for c := range f.children {
		child := f.children[c]
		child.component.Draw(screen, child.bounds.Add(frame.Min))
	}
}

func (cont *ContainerEmbed) AddChild(child Component) {
	c := &Child{component: child, handledTouchID: -1}
	cont.children = append(cont.children, c)
	cont.isDirty = true
}

func (cont *ContainerEmbed) HandleJustPressedTouchID(touchID ebiten.TouchID) bool {
	result := false
	x, y := ebiten.TouchPosition(touchID)
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		handler, ok := child.component.(TouchHandler)
		if ok && handler != nil {
			if result == false && IsInside(child.bounds, x, y) {
				if handler.HandleJustPressedTouchID(touchID) {
					child.handledTouchID = touchID
					result = true
					break
				}
			}
		}

		button, ok := child.component.(ButtonComponent)
		if ok && button != nil {
			if result == false && IsInside(child.bounds, x, y) {
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
					x, y := ebiten.TouchPosition(touchID)
					child.IsButtonPressed = false
					child.handledTouchID = -1
					button.HandleRelease(touchID, IsInside(child.bounds, x, y))
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
			if result == false && IsInside(child.bounds, x, y) {
				if handler.HandleMouse(x, y) {
					result = true
				}
			}
		}

		button, ok := child.component.(ButtonComponent)
		if ok && button != nil && child.handledTouchID == -1 {
			if result == false && IsInside(child.bounds, x, y) {
				if child.IsButtonPressed {
					if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) == false {
						button.HandleRelease(-1, true)
						child.IsButtonPressed = false
					}
				} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					button.HandlePress(-1)
					child.IsButtonPressed = true
				}
				result = true
			} else {
				if child.IsButtonPressed {
					if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) == false {
						button.HandleRelease(-1, false)
						child.IsButtonPressed = false
					}
				}
			}
		}
	}
	return result
}
