package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type Layer struct {
	cont  Container
	frame image.Rectangle
}

func NewLayerWithContainer(cont Container) *Layer {
	l := &Layer{cont: cont}
	return l
}

func (l *Layer) Layout(x0, y0, x1, y1 int) {
	l.frame = image.Rect(x0, y0, x1, y1)
}

func (l *Layer) Update() {
	l.cont.Update()
}

func (l *Layer) HandleTouch(touchID int) bool {
	touchable, ok := l.cont.(TouchableComponent)
	if ok == false {
		return false
	}
	return touchable.HandleTouch(touchID)
}

func (l *Layer) Draw(screen *ebiten.Image) {
	l.cont.Draw(screen, l.frame)
}
