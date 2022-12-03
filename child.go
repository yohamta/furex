package furex

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type child struct {
	absolute                 bool
	item                     *View
	bounds                   image.Rectangle
	isButtonPressed          bool
	isMouseLeftButtonHandler bool
	isMouseEntered           bool
	handledTouchID           ebiten.TouchID
	swipe
}

type swipe struct {
	downX, downY int
	upX, upY     int
	downTime     time.Time
	upTime       time.Time
	swipeDir     SwipeDirection
	swipeTouchID ebiten.TouchID
}

func (c *child) HandleJustPressedTouchID(
	frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	var result = false
	if c.checkButtonHandlerStart(frame, touchID, x, y) {
		result = true
	}
	if !result && c.checkTouchHandlerStart(frame, touchID, x, y) {
		result = true
	}
	c.checkSwipeHandlerStart(frame, touchID, x, y)
	return result
}

func (c *child) HandleJustReleasedTouchID(
	frame *image.Rectangle, touchID ebiten.TouchID, x, y int) {
	c.checkTouchHandlerEnd(frame, touchID, x, y)
	c.checkButtonHandlerEnd(frame, touchID, x, y)
	c.checkSwipeHandlerEnd(frame, touchID, x, y)
}

func (c *child) checkTouchHandlerStart(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	touchHandler, ok := c.item.Handler.(TouchHandler)
	if ok {
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
	if ok {
		if c.handledTouchID == touchID {
			touchHandler.HandleJustReleasedTouchID(touchID, x, y)
			c.handledTouchID = -1
		}
	}
}

func (c *child) checkSwipeHandlerStart(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	_, ok := c.item.Handler.(SwipeHandler)
	if ok {
		if isInside(frame, x, y) {
			c.swipeTouchID = touchID
			c.swipe.downTime = time.Now()
			c.swipe.downX, c.swipe.downY = x, y
			return true
		}
	}
	return false
}

func (c *child) checkSwipeHandlerEnd(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	swipeHandler, ok := c.item.Handler.(SwipeHandler)
	if ok {
		if c.swipeTouchID != touchID {
			return false
		}
		c.swipeTouchID = -1
		c.upTime = time.Now()
		c.upX, c.upY = x, y
		if c.checkSwipe() {
			swipeHandler.HandleSwipe(c.swipeDir)
			return true
		}
	}
	return false
}

const swipeThresholdDist = 50.
const swipeThresholdTime = time.Millisecond * 300

func (c *child) checkSwipe() bool {
	dur := c.upTime.Sub(c.downTime)
	if dur > swipeThresholdTime {
		return false
	}

	deltaX := float64(c.downX - c.upX)
	if math.Abs(deltaX) >= swipeThresholdDist {
		if deltaX > 0 {
			c.swipeDir = SwipeDirectionLeft
		} else {
			c.swipeDir = SwipeDirectionRight
		}
		return true
	}

	deltaY := float64(c.downY - c.upY)
	if math.Abs(deltaY) >= swipeThresholdDist {
		if deltaY > 0 {
			c.swipeDir = SwipeDirectionUp
		} else {
			c.swipeDir = SwipeDirectionDown
		}
		return true
	}

	return false
}

func (c *child) checkButtonHandlerStart(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) bool {
	button, ok := c.item.Handler.(ButtonHandler)
	if ok {
		for {
			if button, ok := c.item.Handler.(NotButton); ok {
				if !button.IsButton() {
					break
				}
			}
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
			break
		}
	}
	return false
}

func (c *child) checkButtonHandlerEnd(frame *image.Rectangle, touchID ebiten.TouchID, x, y int) {
	button, ok := c.item.Handler.(ButtonHandler)
	if ok {
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
