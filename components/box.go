package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
)

type Box struct {
	Color color.Color
}

var _ furex.DrawHandler = (*Box)(nil)

func (b *Box) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	furex.G.FillRect(screen, &furex.FillRectOpts{
		Rect: frame, Color: b.Color,
	})
}
