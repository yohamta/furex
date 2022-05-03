package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type child struct {
	item                     *View
	bounds                   image.Rectangle
	isButtonPressed          bool
	isMouseLeftButtonHandler bool
	handledTouchID           ebiten.TouchID
}

func (c *child) HandleJustPressedTouchID(
	frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	if c.checkTouchHandlerStart(frame, touchID, x, y) {
		return true
	}
	if c.checkButtonHandlerStart(frame, touchID, x, y) {
		return true
	}
	return false
}

func (c *child) HandleJustReleasedTouchID(
	frame *image.Rectangle, touchID ebiten.TouchID, x, y int) {
	c.checkTouchHandlerEnd(frame, touchID, x, y)
	c.checkButtonHandlerEnd(frame, touchID, x, y)
}

func (c *child) checkTouchHandlerStart(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	touchHandler, ok := c.item.Handler.(TouchHandler)
	if ok && touchHandler != nil {
		if isInside(frame, x, y) {
			if touchHandler.HandleJustPressedTouchID(touchID, x, y) {
				c.handledTouchID = touchID
				return true
			}
		}
	}
	return false
}

func (c *child) checkTouchHandlerEnd(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) {
	touchHandler, ok := c.item.Handler.(TouchHandler)
	if ok && touchHandler != nil {
		if c.handledTouchID == touchID {
			touchHandler.HandleJustReleasedTouchID(touchID, x, y)
			c.handledTouchID = -1
		}
	}
}

func (c *child) checkButtonHandlerStart(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	button, ok := c.item.Handler.(ButtonHandler)
	if ok && button != nil {
		if isInside(frame, x, y) {
			if !c.isButtonPressed {
				c.isButtonPressed = true
				c.handledTouchID = touchID
				button.HandlePress(x, y, touchID)
			}
			return true
		} else if c.handledTouchID == touchID {
			c.handledTouchID = -1
		}
	}
	return false
}

func (c *child) checkButtonHandlerEnd(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) {
	button, ok := c.item.Handler.(ButtonHandler)
	if ok && button != nil {
		if c.handledTouchID == touchID {
			if c.isButtonPressed {
				c.isButtonPressed = false
				c.handledTouchID = -1
				if x == 0 && y == 0 {
					button.HandleRelease(x, y, false)
				} else {
					button.HandleRelease(x, y, !isInside(frame, x, y))
				}
			}
		}
	}
}
