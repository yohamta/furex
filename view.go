package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type ChildLayer struct {
	layer    *Layer
	touchIDs []int
}

type View struct {
	layers []*ChildLayer
	frame  image.Rectangle
}

func NewView() *View {
	cont := new(View)

	return cont
}

func (cont *View) AddLayer(l *Layer) {
	child := &ChildLayer{layer: l}
	cont.layers = append(cont.layers, child)
	f := cont.frame
	l.Layout(f.Min.X, f.Min.Y, f.Max.X, f.Max.Y)
}

func (cont *View) Layout(x0, y0, x1, y1 int) {
	cont.frame = image.Rect(x0, y0, x1, y1)
	for l := range cont.layers {
		cont.layers[l].layer.Layout(x0, y0, x1, y1)
	}
}

func (cont *View) Update() {
	for l := range cont.layers {
		cont.layers[l].layer.Update()
	}
	cont.handleTouch()
	cont.handleMouse()
}

func (cont *View) Draw(screen *ebiten.Image) {
	for l := range cont.layers {
		cont.layers[l].layer.Draw(screen)
	}
}

func (cont *View) handleTouch() {
	justPressedTouchIds := inpututil.JustPressedTouchIDs()

	if justPressedTouchIds != nil {
		for i := 0; i < len(justPressedTouchIds); i++ {
			touchID := justPressedTouchIds[i]
			for j := len(cont.layers) - 1; j >= 0; j-- {
				if cont.layers[j].layer.HandleJustPressedTouchID(touchID) {
					cont.layers[j].touchIDs = append(cont.layers[j].touchIDs, touchID)
					break
				}
			}
		}
		for j := len(cont.layers) - 1; j >= 0; j-- {
			touchIDs := cont.layers[j].touchIDs
			for t := range touchIDs {
				if inpututil.IsTouchJustReleased(touchIDs[t]) {
					cont.layers[j].layer.HandleJustReleasedTouchID(touchIDs[t])
					cont.layers[j].touchIDs = append(touchIDs[:t], touchIDs[t+1:]...)
				}
			}
		}
	}
}

func (cont *View) handleMouse() {
	x, y := ebiten.CursorPosition()
	for j := len(cont.layers) - 1; j >= 0; j-- {
		if cont.layers[j].layer.HandleMouse(x, y) {
			break
		}
	}
}
