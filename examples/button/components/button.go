package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/furex"
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
		furex.FillRect(screen, frame, color.RGBA{0xff, 0, 0, 0xff})
	} else {
		furex.FillRect(screen, frame, color.RGBA{0, 0xff, 0, 0xff})
	}
	furex.DrawRect(screen, frame, color.RGBA{0xff, 0xff, 0xff, 0xff}, 2)
}
