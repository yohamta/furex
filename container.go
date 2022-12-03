package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type containerEmbed struct {
	children []*child
	isDirty  bool
	frame    image.Rectangle
	touchIDs []ebiten.TouchID
}

func (ct *containerEmbed) processEvent() {
	ct.handleTouchEvents()
	ct.handleMouseEvents()
}

// Draw draws it's children
func (ct *containerEmbed) Draw(screen *ebiten.Image) {
	for _, c := range ct.children {
		b := c.bounds
		if !c.absolute {
			b = c.bounds.Add(ct.frame.Min)
		}
		if c.item.Handler != nil {
			c.item.Handler.HandleDraw(screen, b)
		}
		c.item.Draw(screen)
	}
}

func (ct *containerEmbed) HandleJustPressedTouchID(touchID ebiten.TouchID, x, y int) bool {
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		if child.HandleJustPressedTouchID(childFrame, touchID, x, y) {
			return true
		}
		if child.item.HandleJustPressedTouchID(touchID, x, y) {
			return true
		}
	}
	return false
}

func (ct *containerEmbed) HandleJustReleasedTouchID(touchID ebiten.TouchID, x, y int) {
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		child.HandleJustReleasedTouchID(childFrame, touchID, x, y)
		child.item.HandleJustReleasedTouchID(touchID, x, y)
	}
}

func (ct *containerEmbed) handleMouse(x, y int) bool {
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		mouseHandler, ok := child.item.Handler.(MouseHandler)
		if ok && mouseHandler != nil {
			if isInside(childFrame, x, y) {
				if mouseHandler.HandleMouse(x, y) {
					return true
				}
			}
		}
		if child.item.handleMouse(x, y) {
			return true
		}
	}
	return false
}

func (ct *containerEmbed) handleMouseEnterLeave(x, y int) bool {
	result := false
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		mouseHandler, ok := child.item.Handler.(MouseEnterLeaveHandler)
		if ok {
			if !result && !child.isMouseEntered && isInside(childFrame, x, y) {
				if mouseHandler.HandleMouseEnter(x, y) {
					result = true
					child.isMouseEntered = true
				}
			}

			if child.isMouseEntered && !isInside(childFrame, x, y) {
				child.isMouseEntered = false
				mouseHandler.HandleMouseLeave()
			}
		}

		if child.item.handleMouseEnterLeave(x, y) {
			result = true
		}
	}
	return result
}

func (ct *containerEmbed) handleMouseButtonLeftPressed(x, y int) bool {
	result := false

	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		mouseLeftClickHandler, ok := child.item.Handler.(MouseLeftButtonHandler)
		if ok {
			if !result && isInside(childFrame, x, y) {
				if mouseLeftClickHandler.HandleJustPressedMouseButtonLeft(x, y) {
					result = true
					child.isMouseLeftButtonHandler = true
				}
			}
		}

		button, ok := child.item.Handler.(ButtonHandler)
		if ok {
			if !result && isInside(childFrame, x, y) {
				if !child.isButtonPressed {
					child.isButtonPressed = true
					child.isMouseLeftButtonHandler = true
					result = true
					button.HandlePress(x, y, -1)
				}
			}
		}

		if !result && child.item.handleMouseButtonLeftPressed(x, y) {
			result = true
		}
	}
	return result
}

func (ct *containerEmbed) handleMouseButtonLeftReleased(x, y int) {
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		mouseLeftClickHandler, ok := child.item.Handler.(MouseLeftButtonHandler)
		if ok {
			if child.isMouseLeftButtonHandler {
				child.isMouseLeftButtonHandler = false
				mouseLeftClickHandler.HandleJustReleasedMouseButtonLeft(x, y)
			}
		}

		button, ok := child.item.Handler.(ButtonHandler)
		if ok {
			if child.isButtonPressed && child.isMouseLeftButtonHandler {
				child.isButtonPressed = false
				child.isMouseLeftButtonHandler = false
				if x == 0 && y == 0 {
					button.HandleRelease(x, y, true)
				} else {
					button.HandleRelease(x, y, !isInside(ct.childFrame(child), x, y))
				}
			}
		}

		child.item.handleMouseButtonLeftReleased(x, y)
	}
}

func isInside(r *image.Rectangle, x, y int) bool {
	return r.Min.X <= x && x <= r.Max.X && r.Min.Y <= y && y <= r.Max.Y
}

func (ct *containerEmbed) handleTouchEvents() {
	justPressedTouchIds := inpututil.AppendJustPressedTouchIDs(nil)

	if justPressedTouchIds != nil {
		for i := 0; i < len(justPressedTouchIds); i++ {
			touchID := justPressedTouchIds[i]
			x, y := ebiten.TouchPosition(touchID)
			recordTouchPosition(touchID, x, y)

			ct.HandleJustPressedTouchID(touchID, x, y)
			ct.touchIDs = append(ct.touchIDs, touchID)
		}
	}

	touchIDs := ct.touchIDs
	for t := range touchIDs {
		if inpututil.IsTouchJustReleased(touchIDs[t]) {
			pos := lastTouchPosition(touchIDs[t])
			ct.HandleJustReleasedTouchID(touchIDs[t], pos.X, pos.Y)
		} else {
			x, y := ebiten.TouchPosition(touchIDs[t])
			recordTouchPosition(touchIDs[t], x, y)
		}
	}
}

func (ct *containerEmbed) handleMouseEvents() {
	x, y := ebiten.CursorPosition()
	ct.handleMouse(x, y)
	ct.handleMouseEnterLeave(x, y)
	if inpututil.IsMouseButtonJustPressed((ebiten.MouseButtonLeft)) {
		ct.handleMouseButtonLeftPressed(x, y)
	}
	if inpututil.IsMouseButtonJustReleased((ebiten.MouseButtonLeft)) {
		ct.handleMouseButtonLeftReleased(x, y)
	}
}

func (ct *containerEmbed) setFrame(frame image.Rectangle) {
	ct.frame = frame
	ct.isDirty = true
}

func (ct *containerEmbed) childFrame(c *child) *image.Rectangle {
	if !c.absolute {
		r := c.bounds.Add(ct.frame.Min)
		return &r
	}
	return &c.bounds
}

type touchPosition struct {
	X, Y int
}

var (
	touchPositions = make(map[ebiten.TouchID]touchPosition)
)

func recordTouchPosition(t ebiten.TouchID, x, y int) {
	touchPositions[t] = touchPosition{x, y}
}

func lastTouchPosition(t ebiten.TouchID) *touchPosition {
	s, ok := touchPositions[t]
	if ok {
		return &s
	}
	return &touchPosition{0, 0}
}
