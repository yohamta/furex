package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/furex/internal/touch"
)

// Container represents a container that can have child components
type Container interface {
	// SetFrame sets the location (x,y) and size (width,height) relative to the window (0,0).
	SetFrame(image.Rectangle)

	// AddChild adds a child component
	AddChild(child Component)

	// AddChild adds a child component
	AddChildContainer(child Container)

	// Draw draws its children component
	Draw(screen *ebiten.Image)

	// Size returns the size(x,y) of the container.
	Size() image.Point

	// Update updates the container
	Update()

	setParent(parent Container)
}

type Child struct {
	bounds image.Rectangle
	item   interface{}

	component Component
	container Container

	isButtonPressed         bool
	isMouseLeftClickHandler bool
	handledTouchID          ebiten.TouchID
}

func NewChild(item interface{}, component Component, container Container) *Child {
	return &Child{
		item:           item,
		component:      component,
		container:      container,
		handledTouchID: -1,
	}
}

type ContainerEmbed struct {
	children []*Child
	isDirty  bool
	frame    image.Rectangle
	parent   Container
	touchIDs []ebiten.TouchID
}

func (cont *ContainerEmbed) processEvent() {
	if cont.IsRoot() {
		cont.handleTouch()
		cont.handleMouse()
	}
}

func (cont *ContainerEmbed) IsRoot() bool {
	return cont.parent == nil
}

// Draw draws it's children
func (cont *ContainerEmbed) Draw(screen *ebiten.Image) {
	for c := range cont.children {
		child := cont.children[c]
		if child.component != nil {
			child.component.Draw(screen, child.bounds.Add(cont.frame.Min))
		} else {
			child.container.Draw(screen)
		}
	}
}

// SetFrame sets the location (x,y) and size (width,height) relative to the window (0,0).
func (cont *ContainerEmbed) SetFrame(frame image.Rectangle) {
	cont.frame = frame
	cont.isDirty = true
}

// SetFramePosition sets the location (x,y) relative to the window (0,0).
func (cont *ContainerEmbed) SetFramePosition(pos image.Point) {
	cont.SetFrame(image.Rect(pos.X, pos.Y, pos.X+cont.frame.Dx(), pos.Y+cont.frame.Dy()))
}

// AddChild adds child component
func (cont *ContainerEmbed) AddChild(child Component) {
	c := NewChild(child, child, nil)
	cont.children = append(cont.children, c)
	cont.isDirty = true
}

// AddChildContainer adds child container
func (cont *ContainerEmbed) AddChildContainer(child Container) {
	c := NewChild(child, nil, child)
	cont.children = append(cont.children, c)
	cont.isDirty = true
	child.setParent(cont)
}

// Update updates the contaienr
func (cont *ContainerEmbed) Update() {}

// SetSize sets the size of the flex container.
func (cont *ContainerEmbed) SetSize(size image.Point) {
	cont.frame = image.Rect(
		cont.frame.Min.X,
		cont.frame.Min.Y,
		cont.frame.Min.X+size.X,
		cont.frame.Min.Y+size.Y,
	)
}

// Size returns the size of the contaienr
func (cont *ContainerEmbed) Size() image.Point {
	return cont.frame.Size()
}

func (cont *ContainerEmbed) setParent(parent Container) {
	cont.parent = parent
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
		touchHandler, ok := child.item.(TouchHandler)
		if ok && touchHandler != nil {
			if result == false && isInside(childFrame, x, y) {
				if touchHandler.HandleJustPressedTouchID(touchID, x, y) {
					child.handledTouchID = touchID
					result = true
					break
				}
			}
		}

		button, ok := child.item.(Button)
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
		touchHandler, ok := child.item.(TouchHandler)
		if ok && touchHandler != nil {
			if child.handledTouchID == touchID {
				touchHandler.HandleJustReleasedTouchID(touchID, x, y)
				child.handledTouchID = -1
			}
		}

		button, ok := child.item.(Button)
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
		mouseHandler, ok := child.item.(MouseHandler)
		if ok && mouseHandler != nil {
			if result == false && isInside(childFrame, x, y) {
				if mouseHandler.HandleMouse(x, y) {
					result = true
				}
			}
		}
	}
	return result
}

func (cont *ContainerEmbed) HandleJustPressedMouseButtonLeft(x, y int) bool {
	result := false

	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		childFrame := cont.childFrame(child)
		mouseLeftClickHandler, ok := child.item.(MouseLeftClickHandler)
		if ok && mouseLeftClickHandler != nil {
			if result == false && isInside(childFrame, x, y) {
				if mouseLeftClickHandler.HandleJustPressedMouseButtonLeft(x, y) {
					result = true
					child.isMouseLeftClickHandler = true
				}
			}
		}

		button, ok := child.item.(Button)
		if ok && button != nil {
			if result == false && isInside(childFrame, x, y) {
				if child.isButtonPressed == false {
					child.isButtonPressed = true
					child.isMouseLeftClickHandler = true
					result = true
					button.HandlePress()
				}
			}
		}
	}
	return result
}

func (cont *ContainerEmbed) HandleJustReleasedMouseButtonLeft(x, y int) {
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		mouseLeftClickHandler, ok := child.item.(MouseLeftClickHandler)
		if ok && mouseLeftClickHandler != nil {
			if child.isMouseLeftClickHandler {
				child.isMouseLeftClickHandler = false
				mouseLeftClickHandler.HandleJustReleasedMouseButtonLeft(x, y)
			}
		}

		button, ok := child.item.(Button)
		if ok && button != nil {
			if child.isButtonPressed == true && child.isMouseLeftClickHandler {
				child.isButtonPressed = false
				child.isMouseLeftClickHandler = false
				if x == 0 && y == 0 {
					button.HandleRelease(true)
				} else {
					button.HandleRelease(isInside(cont.childFrame(child), x, y))
				}
			}
		}
	}
}

func isInside(r *image.Rectangle, x, y int) bool {
	return r.Min.X <= x && x <= r.Max.X && r.Min.Y <= y && y <= r.Max.Y
}

func (cont *ContainerEmbed) handleTouch() {
	justPressedTouchIds := inpututil.JustPressedTouchIDs()

	if justPressedTouchIds != nil {
		for i := 0; i < len(justPressedTouchIds); i++ {
			touchID := justPressedTouchIds[i]
			x, y := ebiten.TouchPosition(touchID)
			touch.RecordTouchPosition(touchID, x, y)

			cont.HandleJustPressedTouchID(touchID, x, y)
			cont.touchIDs = append(cont.touchIDs, touchID)
		}
	}

	touchIDs := cont.touchIDs
	for t := range touchIDs {
		if inpututil.IsTouchJustReleased(touchIDs[t]) {
			pos := touch.LastTouchPosition(touchIDs[t])
			cont.HandleJustReleasedTouchID(touchIDs[t], pos.X, pos.Y)
		} else {
			x, y := ebiten.TouchPosition(touchIDs[t])
			touch.RecordTouchPosition(touchIDs[t], x, y)
		}
	}
}

func (cont *ContainerEmbed) handleMouse() {
	x, y := ebiten.CursorPosition()
	cont.HandleMouse(x, y)
	if inpututil.IsMouseButtonJustPressed((ebiten.MouseButtonLeft)) {
		cont.HandleJustPressedMouseButtonLeft(x, y)
	}
	if inpututil.IsMouseButtonJustReleased((ebiten.MouseButtonLeft)) {
		cont.HandleJustReleasedMouseButtonLeft(x, y)
	}
}
