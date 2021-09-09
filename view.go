package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ChildLayer struct {
	layer    *Layer
	touchIDs []ebiten.TouchID
}

type View struct {
	layers []*ChildLayer
	frame  image.Rectangle
}

func NewView() *View {
	view := new(View)

	return view
}

func (view *View) AddLayer(l *Layer) {
	child := &ChildLayer{layer: l}
	view.layers = append(view.layers, child)
	f := view.frame
	l.Layout(f.Min.X, f.Min.Y, f.Max.X, f.Max.Y)
}

func (view *View) Layout(x0, y0, x1, y1 int) {
	view.frame = image.Rect(x0, y0, x1, y1)
	for l := range view.layers {
		view.layers[l].layer.Layout(x0, y0, x1, y1)
	}
}

func (view *View) Update() {
	for l := range view.layers {
		view.layers[l].layer.Update()
	}
	view.handleTouch()
	view.handleMouse()
}

func (view *View) Draw(screen *ebiten.Image) {
	for l := range view.layers {
		view.layers[l].layer.Draw(screen)
	}
}

func (view *View) handleTouch() {
	justPressedTouchIds := inpututil.JustPressedTouchIDs()

	if justPressedTouchIds != nil {
		for i := 0; i < len(justPressedTouchIds); i++ {
			touchID := justPressedTouchIds[i]
			recordTouchPosition(touchID)
			for j := len(view.layers) - 1; j >= 0; j-- {
				if view.layers[j].layer.HandleJustPressedTouchID(touchID) {
					view.layers[j].touchIDs = append(view.layers[j].touchIDs, touchID)
					break
				}
			}
		}
	}
	for j := len(view.layers) - 1; j >= 0; j-- {
		touchIDs := view.layers[j].touchIDs
		for t := range touchIDs {
			if inpututil.IsTouchJustReleased(touchIDs[t]) {
				view.layers[j].layer.HandleJustReleasedTouchID(touchIDs[t])
				view.layers[j].touchIDs = append(touchIDs[:t], touchIDs[t+1:]...)
			} else {
				recordTouchPosition(touchIDs[t])
			}
		}
	}
}

func (view *View) handleMouse() {
	x, y := ebiten.CursorPosition()
	for j := len(view.layers) - 1; j >= 0; j-- {
		if view.layers[j].layer.HandleMouse(x, y) {
			break
		}
	}
}
