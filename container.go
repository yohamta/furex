package furex

import (
	"image"
	"image/color"

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
	ct.handleTouch()
	ct.handleMouse()
}

// Draw draws it's children
func (ct *containerEmbed) Draw(screen *ebiten.Image) {
	if Debug {
		G.DrawRect(screen, &DrawRectOpts{
			Rect:        ct.frame,
			Color:       color.RGBA{0xff, 0xff, 0, 0xff},
			StrokeWidth: 2,
		})
	}
	for _, c := range ct.children {
		b := c.bounds.Add(ct.frame.Min)
		if c.item.Handler != nil {
			c.item.Handler.HandleDraw(screen, b)
		}
		c.item.Draw(screen)
		if Debug {
			G.DrawRect(screen, &DrawRectOpts{
				Rect:        b,
				Color:       color.RGBA{0xff, 0, 0xff, 0xff},
				StrokeWidth: 2,
			})
		}
	}
}

func (ct *containerEmbed) HandleJustPressedTouchID(
	touchID ebiten.TouchID, x, y int,
) bool {
	result := false
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		touchHandler, ok := child.item.Handler.(TouchHandler)
		if ok && touchHandler != nil {
			if !result && isInside(childFrame, x, y) {
				if touchHandler.HandleJustPressedTouchID(touchID, x, y) {
					child.handledTouchID = touchID
					result = true
					break
				}
			}
		}

		button, ok := child.item.Handler.(ButtonHandler)
		if ok && button != nil {
			if !result && isInside(childFrame, x, y) {
				if !child.isButtonPressed {
					child.isButtonPressed = true
					child.handledTouchID = touchID
					button.HandlePress(x, y, touchID)
				}
				result = true
			} else if child.handledTouchID == touchID {
				child.handledTouchID = -1
			}
		}

		if !result && child.item.HandleJustPressedTouchID(touchID, x, y) {
			result = true
			break
		}
	}
	return result
}

func (ct *containerEmbed) HandleJustReleasedTouchID(touchID ebiten.TouchID, x, y int) {
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		touchHandler, ok := child.item.Handler.(TouchHandler)
		if ok && touchHandler != nil {
			if child.handledTouchID == touchID {
				touchHandler.HandleJustReleasedTouchID(touchID, x, y)
				child.handledTouchID = -1
			}
		}

		button, ok := child.item.Handler.(ButtonHandler)
		if ok && button != nil {
			if child.handledTouchID == touchID {
				if child.isButtonPressed {
					child.isButtonPressed = false
					child.handledTouchID = -1
					if x == 0 && y == 0 {
						button.HandleRelease(x, y, false)
					} else {
						button.HandleRelease(x, y,
							!isInside(ct.childFrame(child), x, y))
					}
				}
			}
		}

		child.item.HandleJustReleasedTouchID(touchID, x, y)
	}
}

func (ct *containerEmbed) HandleMouse(x, y int) bool {
	result := false
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		mouseHandler, ok := child.item.Handler.(MouseHandler)
		if ok && mouseHandler != nil {
			if !result && isInside(childFrame, x, y) {
				if mouseHandler.HandleMouse(x, y) {
					result = true
				}
			}
		}

		if !result && child.item.HandleMouse(x, y) {
			result = true
		}
	}
	return result
}

func (ct *containerEmbed) HandleJustPressedMouseButtonLeft(x, y int) bool {
	result := false

	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		childFrame := ct.childFrame(child)
		mouseLeftClickHandler, ok := child.item.Handler.(MouseLeftButtonHandler)
		if ok && mouseLeftClickHandler != nil {
			if !result && isInside(childFrame, x, y) {
				if mouseLeftClickHandler.HandleJustPressedMouseButtonLeft(x, y) {
					result = true
					child.isMouseLeftButtonHandler = true
				}
			}
		}

		button, ok := child.item.Handler.(ButtonHandler)
		if ok && button != nil {
			if !result && isInside(childFrame, x, y) {
				if !child.isButtonPressed {
					child.isButtonPressed = true
					child.isMouseLeftButtonHandler = true
					result = true
					button.HandlePress(x, y, -1)
				}
			}
		}

		if !result && child.item.HandleJustPressedMouseButtonLeft(x, y) {
			result = true
		}
	}
	return result
}

func (ct *containerEmbed) HandleJustReleasedMouseButtonLeft(x, y int) {
	for c := len(ct.children) - 1; c >= 0; c-- {
		child := ct.children[c]
		mouseLeftClickHandler, ok := child.item.Handler.(MouseLeftButtonHandler)
		if ok && mouseLeftClickHandler != nil {
			if child.isMouseLeftButtonHandler {
				child.isMouseLeftButtonHandler = false
				mouseLeftClickHandler.HandleJustReleasedMouseButtonLeft(x, y)
			}
		}

		button, ok := child.item.Handler.(ButtonHandler)
		if ok && button != nil {
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

		child.item.HandleJustReleasedMouseButtonLeft(x, y)
	}
}

func isInside(r *image.Rectangle, x, y int) bool {
	return r.Min.X <= x && x <= r.Max.X && r.Min.Y <= y && y <= r.Max.Y
}

func (ct *containerEmbed) handleTouch() {
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

func (ct *containerEmbed) handleMouse() {
	x, y := ebiten.CursorPosition()
	ct.HandleMouse(x, y)
	if inpututil.IsMouseButtonJustPressed((ebiten.MouseButtonLeft)) {
		ct.HandleJustPressedMouseButtonLeft(x, y)
	}
	if inpututil.IsMouseButtonJustReleased((ebiten.MouseButtonLeft)) {
		ct.HandleJustReleasedMouseButtonLeft(x, y)
	}
}

func (ct *containerEmbed) setFrame(frame image.Rectangle) {
	ct.frame = frame
	ct.isDirty = true
}

func (ct *containerEmbed) childFrame(c *child) *image.Rectangle {
	r := c.bounds.Add(ct.frame.Min)
	return &r
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
