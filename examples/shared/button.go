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
	margin     []int
}

func NewButton(w, h int) *Button {
	b := new(Button)
	b.size = image.Pt(w, h)
	b.margin = []int{0, 0, 0, 0}
	return b
}

func (b *Button) Size() (int, int) {
	return b.size.X, b.size.Y
}

func (b *Button) Margin() []int {
	return b.margin
}

func (b *Button) SetMargin(m []int) {
	b.margin = m
}

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.isPressing = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.isPressing = false
	if isCancel {
		println("The click is cancelled!")
	} else {
		println("clicked!")
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
