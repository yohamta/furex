package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/furex/internal/touch"
)

// View manages a root flex container
type View struct {
	frame    image.Rectangle
	touchIDs []ebiten.TouchID
	flex     *Flex
}

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
			x, y := ebiten.TouchPosition(touchID)
			touch.RecordTouchPosition(touchID, x, y)

			v.flex.HandleJustPressedTouchID(touchID, x, y)
			v.touchIDs = append(v.touchIDs, touchID)
		}
	}
	touchIDs := v.touchIDs
	for t := range touchIDs {
		if inpututil.IsTouchJustReleased(touchIDs[t]) {
			pos := touch.LastTouchPosition(touchIDs[t])
			v.flex.HandleJustReleasedTouchID(touchIDs[t], pos.X, pos.Y)
		} else {
			x, y := ebiten.TouchPosition(touchIDs[t])
			touch.RecordTouchPosition(touchIDs[t], x, y)
		}
	}
}

func (v *View) handleMouse() {
	x, y := ebiten.CursorPosition()
	v.flex.HandleMouse(x, y)
	if inpututil.IsMouseButtonJustPressed((ebiten.MouseButtonLeft)) {
		v.flex.HandleJustPressedMouseButtonLeft(x, y)
	}
	if inpututil.IsMouseButtonJustReleased((ebiten.MouseButtonLeft)) {
		v.flex.HandleJustReleasedMouseButtonLeft(x, y)
	}
}
