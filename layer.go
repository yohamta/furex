package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
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
	updatable, ok := l.cont.(UpdatableComponent)
	if ok && updatable != nil {
		updatable.Update()
	}
}

func (l *Layer) HandleJustPressedTouchID(touchID ebiten.TouchID) bool {
	touchable, ok := l.cont.(TouchHandler)
	if ok == false {
		return false
	}
	return touchable.HandleJustPressedTouchID(touchID)
}

func (l *Layer) HandleJustReleasedTouchID(touchID ebiten.TouchID) {
	touchable, ok := l.cont.(TouchHandler)
	if ok == false {
		return
	}
	touchable.HandleJustReleasedTouchID(touchID)
}

func (l *Layer) HandleMouse(x, y int) bool {
	clickable, ok := l.cont.(MouseHandler)
	if ok == false {
		return false
	}
	return clickable.HandleMouse(x, y)
}

func (l *Layer) Draw(screen *ebiten.Image) {
	l.cont.Draw(screen, l.frame)
}
