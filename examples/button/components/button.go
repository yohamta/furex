package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/examples/shared"
)

type SampleButton struct {
	size       image.Point
	isPressing bool
}

func NewSampleButton(w, h int) *SampleButton {
	button := new(SampleButton)
	button.size = image.Pt(w, h)
	return button
}

func (button *SampleButton) GetSize() image.Point {
	return button.size
}

func (button *SampleButton) OnPressButton() {
	button.isPressing = true
}

func (button *SampleButton) OnReleaseButton() {
	button.isPressing = false
}

func (button *SampleButton) Update() {}

func (button *SampleButton) Draw(screen *ebiten.Image, frame image.Rectangle) {
	if button.isPressing {
		shared.FillRect(screen, frame, color.RGBA{0xff, 0, 0, 0xff})
	} else {
		shared.FillRect(screen, frame, color.RGBA{0, 0xff, 0, 0xff})
	}
	shared.DrawRect(screen, frame, color.RGBA{0xff, 0xff, 0xff, 0xff}, 2)
}
