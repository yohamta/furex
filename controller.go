package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type Controller struct {
	root  Container
	frame image.Rectangle
}

func NewController() *Controller {
	cont := new(Controller)
	return cont
}

func (cont *Controller) SetRootContaienr(c Container) {
	cont.root = c
}

func (cont *Controller) Layout(x0, y0, x1, y1 int) {
	cont.frame = image.Rect(x0, y0, x1, y1)
}

func (cont *Controller) Update() {
	cont.root.Update()
}

func (cont *Controller) HandleTouch(touchID int) bool {
	touchable, ok := cont.root.(Touchable)
	if ok == false {
		return false
	}
	return touchable.HandleTouch(touchID)
}

func (cont *Controller) Draw(screen *ebiten.Image) {
	cont.root.Draw(screen, cont.frame)
}
