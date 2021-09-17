package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/furex/internal/container"
	"github.com/yohamta/furex/internal/touch"
)

// Container represents a container that can have child components.
type Container interface {
	// SetFrame sets the location (x,y) and size (width,height) relative to the window (0,0).
	SetFrame(image.Rectangle)

	// AddChild adds a child component.
	AddChild(child Component)

	// AddChild adds a child component.
	AddChildContainer(child Container)

	// Draw draws its children component.
	Draw(screen *ebiten.Image)

	// Size returns the size(x,y) of the container.
	Size() image.Point

	// Update updates the container.
	Update()

	setParent(parent Container)
}

type ContainerEmbed struct {
	children []*container.Child
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
		container, ok := child.Item.(Container)
		if ok && container != nil {
			container.Draw(screen)
			continue
		}
		component, ok := child.Item.(Component)
		if ok && component != nil {
			component.Draw(screen, child.Bounds.Add(cont.frame.Min))
			continue
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
	c := container.NewChild(child)
	cont.children = append(cont.children, c)
	cont.isDirty = true
}

// AddChildContainer adds child container
func (cont *ContainerEmbed) AddChildContainer(child Container) {
	c := container.NewChild(child)
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

func (cont *ContainerEmbed) childFrame(c *container.Child) *image.Rectangle {
	r := c.Bounds.Add(cont.frame.Min)
	return &r
}

func (cont *ContainerEmbed) HandleJustPressedTouchID(touchID ebiten.TouchID, x, y int) bool {
	result := false
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		childFrame := cont.childFrame(child)
		touchHandler, ok := child.Item.(TouchHandler)
		if ok && touchHandler != nil {
			if result == false && isInside(childFrame, x, y) {
				if touchHandler.HandleJustPressedTouchID(touchID, x, y) {
					child.HandledTouchID = touchID
					result = true
					break
				}
			}
		}

		button, ok := child.Item.(Button)
		if ok && button != nil {
			if result == false && isInside(childFrame, x, y) {
				if child.IsButtonPressed == false {
					child.IsButtonPressed = true
					child.HandledTouchID = touchID
					button.HandlePress()
				}
				result = true
			} else if child.HandledTouchID == touchID {
				child.HandledTouchID = -1
			}
		}
	}
	return result
}

func (cont *ContainerEmbed) HandleJustReleasedTouchID(touchID ebiten.TouchID, x, y int) {
	for c := len(cont.children) - 1; c >= 0; c-- {
		child := cont.children[c]
		touchHandler, ok := child.Item.(TouchHandler)
		if ok && touchHandler != nil {
			if child.HandledTouchID == touchID {
				touchHandler.HandleJustReleasedTouchID(touchID, x, y)
				child.HandledTouchID = -1
			}
		}

		button, ok := child.Item.(Button)
		if ok && button != nil {
			if child.HandledTouchID == touchID {
				if child.IsButtonPressed == true {
					child.IsButtonPressed = false
					child.HandledTouchID = -1
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
		mouseHandler, ok := child.Item.(MouseHandler)
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
		mouseLeftClickHandler, ok := child.Item.(MouseLeftClickHandler)
		if ok && mouseLeftClickHandler != nil {
			if result == false && isInside(childFrame, x, y) {
				if mouseLeftClickHandler.HandleJustPressedMouseButtonLeft(x, y) {
					result = true
					child.IsMouseLeftClickHandler = true
				}
			}
		}

		button, ok := child.Item.(Button)
		if ok && button != nil {
			if result == false && isInside(childFrame, x, y) {
				if child.IsButtonPressed == false {
					child.IsButtonPressed = true
					child.IsMouseLeftClickHandler = true
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
		mouseLeftClickHandler, ok := child.Item.(MouseLeftClickHandler)
		if ok && mouseLeftClickHandler != nil {
			if child.IsMouseLeftClickHandler {
				child.IsMouseLeftClickHandler = false
				mouseLeftClickHandler.HandleJustReleasedMouseButtonLeft(x, y)
			}
		}

		button, ok := child.Item.(Button)
		if ok && button != nil {
			if child.IsButtonPressed == true && child.IsMouseLeftClickHandler {
				child.IsButtonPressed = false
				child.IsMouseLeftClickHandler = false
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
