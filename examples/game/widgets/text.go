package widgets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
)

type Text struct {
	Value string
}

var (
	_ furex.DrawHandler = (*Text)(nil)
)

func (t *Text) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	ebitenutil.DrawRect(
		screen, float64(frame.Min.X), float64(frame.Min.Y), float64(len(t.Value)*6+4), float64(frame.Dy()), color.RGBA{0, 0, 0, 50})
	ebitenutil.DebugPrintAt(screen, t.Value,
		frame.Min.X+2, frame.Min.Y+frame.Dy()/2-8)
}
