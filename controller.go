package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Controller struct {
	layers []*Layer
	frame  image.Rectangle
}

func NewController() *Controller {
	cont := new(Controller)

	return cont
}

func (cont *Controller) AddLayer(l *Layer) {
	cont.layers = append(cont.layers, l)
	f := cont.frame
	l.Layout(f.Min.X, f.Min.Y, f.Max.X, f.Max.Y)
}

func (cont *Controller) Layout(x0, y0, x1, y1 int) {
	cont.frame = image.Rect(x0, y0, x1, y1)
	for l := range cont.layers {
		cont.layers[l].Layout(x0, y0, x1, y1)
	}
}

func (cont *Controller) Update() {
	for l := range cont.layers {
		cont.layers[l].Update()
	}
	cont.handleTouch()
}

func (cont *Controller) Draw(screen *ebiten.Image) {
	for l := range cont.layers {
		cont.layers[l].Draw(screen)
	}
}

func (cont *Controller) handleTouch() {
	justPressedTouchIds := inpututil.JustPressedTouchIDs()

	if justPressedTouchIds != nil {
		for i := 0; i < len(justPressedTouchIds); i++ {
			touchID := justPressedTouchIds[i]
			for j := len(cont.layers) - 1; j >= 0; j-- {
				if cont.layers[j].HandleTouch(touchID) {
					break
				}
			}
		}
	}
}
