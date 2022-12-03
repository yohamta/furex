package widgets

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/common/graphic"
)

type Box struct {
	Color color.Color
}

var _ furex.DrawHandler = (*Box)(nil)

func (b *Box) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	graphic.FillRect(screen, &graphic.FillRectOpts{
		Rect: frame, Color: b.Color,
	})
}
