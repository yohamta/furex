package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
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
	buttonTouchID   int
	handledTouchID  int
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
	c := &Child{component: child}
	cont.children = append(cont.children, c)
	cont.isDirty = true
}

func (cont *ContainerEmbed) HandleJustPressedTouchID(touchID int) bool {
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
				}
			}
		}

		button, ok := child.component.(ButtonComponent)
		if ok && button != nil {
			if result == false && IsInside(child.bounds, x, y) {
				if child.IsButtonPressed == false {
					child.IsButtonPressed = true
					child.buttonTouchID = touchID
					button.OnPressButton()
				}
				result = true
			} else if child.buttonTouchID == touchID {
				child.buttonTouchID = -1
			}
		}
	}
	return result
}

func (cont *ContainerEmbed) HandleJustReleasedTouchID(touchID int) {
	x, y := ebiten.TouchPosition(touchID)
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		handler, ok := child.component.(TouchHandler)
		if ok && handler != nil {
			if child.handledTouchID == touchID && IsInside(child.bounds, x, y) {
				handler.HandleJustReleasedTouchID(touchID)
			}
		}

		button, ok := child.component.(ButtonComponent)
		if ok && button != nil {
			if child.buttonTouchID == touchID && IsInside(child.bounds, x, y) {
				if child.IsButtonPressed == true {
					child.IsButtonPressed = false
					child.buttonTouchID = -1
					button.OnPressButton()
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
		if ok && button != nil {
			if result == false && IsInside(child.bounds, x, y) {
				if child.IsButtonPressed {
					if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) == false {
						button.OnReleaseButton()
						child.IsButtonPressed = false
					}
				} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					button.OnPressButton()
					child.IsButtonPressed = true
				}
				result = true
			} else {
				if child.IsButtonPressed {
					if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) == false {
						button.OnReleaseButton()
						child.IsButtonPressed = false
					}
				}
			}
		}
	}
	return result
}
