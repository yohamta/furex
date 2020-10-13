package furex

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Button displays a color filled button
type Button struct {
	size       image.Point
	isPressing bool
}

func NewButton(w, h int) *Button {
	button := new(Button)
	button.size = image.Pt(w, h)
	return button
}

func (button *Button) GetSize() image.Point {
	return button.size
}

func (button *Button) Update() {}

func (button *Button) OnPressButton() {
	button.isPressing = true
}

func (button *Button) OnReleaseButton() {
	button.isPressing = false
}

func (button *Button) Draw(screen *ebiten.Image, frame image.Rectangle) {
	if button.isPressing {
		FillRect(screen, frame, color.RGBA{0xff, 0, 0, 0xff})
	} else {
		FillRect(screen, frame, color.RGBA{0, 0xff, 0, 0xff})
	}
	DrawRect(screen, frame, color.RGBA{0xff, 0xff, 0xff, 0xff}, 2)
}
