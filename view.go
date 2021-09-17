package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// View manages a root flex container
type View struct {
	frame    image.Rectangle
	touchIDs []ebiten.TouchID
	flex     *Flex
}

type tpos struct {
	x, y int
}

var (
	touchPositions = make(map[ebiten.TouchID]tpos)
)

// NewView makes a View
func NewView(frame image.Rectangle, flex *Flex) *View {
	v := new(View)
	v.frame = frame
	v.flex = flex

	s := v.flex.Size()
	v.flex.SetFrame(image.Rect(
		frame.Min.X, frame.Min.Y, frame.Min.X+s.X, frame.Min.Y+s.Y,
	))

	return v
}

// Update updates the flex container and handle UI interactions.
func (v *View) Update() {
	v.flex.Update()
	v.handleTouch()
	v.handleMouse()
}

// Draw draws the all flex container tree into the screen
func (v *View) Draw(screen *ebiten.Image) {
	v.flex.Draw(screen, v.frame)
}

func (v *View) handleTouch() {
	justPressedTouchIds := inpututil.JustPressedTouchIDs()

	if justPressedTouchIds != nil {
		for i := 0; i < len(justPressedTouchIds); i++ {
			touchID := justPressedTouchIds[i]
			recordTouchPosition(touchID)

			v.flex.HandleJustPressedTouchID(touchID)
			v.touchIDs = append(v.touchIDs, touchID)
		}
	}
	touchIDs := v.touchIDs
	for t := range touchIDs {
		if inpututil.IsTouchJustReleased(touchIDs[t]) {
			v.flex.HandleJustReleasedTouchID(touchIDs[t])
		} else {
			recordTouchPosition(touchIDs[t])
		}
	}
}

func (v *View) handleMouse() {
	x, y := ebiten.CursorPosition()
	v.flex.HandleMouse(x, y)
}

func recordTouchPosition(t ebiten.TouchID) {
	x, y := ebiten.TouchPosition(t)
	touchPositions[t] = tpos{x, y}
}
