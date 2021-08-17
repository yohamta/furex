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

func (button *SampleButton) Size() image.Point {
	return button.size
}

func (button *SampleButton) HandlePress(t ebiten.TouchID) {
	button.isPressing = true
}

func (button *SampleButton) HandleRelease(t ebiten.TouchID, isInside bool) {
	println("isInside: ", isInside)
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
