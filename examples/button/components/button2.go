package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/furex"
)

type Button2 struct {
	size       image.Point
	isPressing bool
}

func PrintRect(rect image.Rectangle) {
	println(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Min.Y)
}

func NewSampleButton2(w, h int) *Button2 {
	button := new(Button2)
	button.size = image.Pt(w, h)
	return button
}

func (button *Button2) GetPosition() image.Point {
	return image.Pt(0, 0)
}

func (button *Button2) GetSize() image.Point {
	return button.size
}

func (button *Button2) HandleJustPressedTouchID(touchID int) bool {
	println("HandleJustPressedTouchID")
	return true
}

func (button *Button2) HandleJustReleasedTouchID(touchID int) {
	println("HandleJustReleasedTouchID")
}

func (button *Button2) Update() {}

func (button *Button2) Draw(screen *ebiten.Image, frame image.Rectangle) {
	// PrintRect(frame)
	if button.isPressing {
		furex.FillRect(screen, frame, color.RGBA{0xff, 0, 0, 0xff})
	} else {
		furex.FillRect(screen, frame, color.RGBA{0, 0xff, 0, 0xff})
	}
	furex.DrawRect(screen, frame, color.RGBA{0xff, 0xff, 0xff, 0xff}, 2)
}
