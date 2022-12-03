package widgets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/text"
)

type Text struct {
	Color     color.Color
	Value     string
	Shadow    bool
	HorzAlign etxt.HorzAlign
	VertAlign etxt.VertAlign
}

var (
	_ furex.DrawHandler = (*Text)(nil)
)

func (t *Text) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	if t.Shadow {
		ebitenutil.DrawRect(
			screen, float64(frame.Min.X), float64(frame.Min.Y), float64(len(t.Value)*6+4), float64(frame.Dy()), color.RGBA{0, 0, 0, 50})
	}
	x, y := frame.Min.X+frame.Dx()/2, frame.Min.Y+frame.Dy()/2
	if t.HorzAlign == etxt.Left {
		x = frame.Min.X
	}
	if t.VertAlign == etxt.Top {
		y = frame.Min.Y
	}
	if t.Color != nil {
		text.R.SetColor(t.Color)
	} else {
		text.R.SetColor(color.White)
	}
	text.R.SetAlign(t.VertAlign, t.HorzAlign)
	text.R.SetTarget(screen)
	text.R.Draw(t.Value, x, y)
}
