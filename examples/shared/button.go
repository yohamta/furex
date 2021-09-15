package shared

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Button struct {
	size       image.Point
	isPressing bool
}

func NewButton(w, h int) *Button {
	b := new(Button)
	b.size = image.Pt(w, h)
	return b
}

func (b *Button) Size() image.Point {
	return b.size
}

func (b *Button) HandlePress(t ebiten.TouchID) {
	b.isPressing = true
}

func (b *Button) HandleRelease(t ebiten.TouchID, isInside bool) {
	b.isPressing = false
	if isInside {
		println("clicked!")
	} else {
		println("The click is cancelled!")
	}
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle) {
	if b.isPressing {
		FillRect(screen, frame, color.RGBA{0xaa, 0, 0, 0xff})
	} else {
		FillRect(screen, frame, color.RGBA{0, 0xaa, 0, 0xff})
	}
	DrawRect(screen, frame, color.RGBA{0xff, 0xff, 0xff, 0xff}, 2)
	ebitenutil.DebugPrintAt(screen, "Button",
		frame.Min.X+((frame.Dx()-36)/2), frame.Min.Y+b.size.Y/2-8)
}