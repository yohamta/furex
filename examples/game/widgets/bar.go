package widgets

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/sprites"
	"github.com/yohamta/ganim8/v2"
)

type Bar struct {
	Value float64
	Color string

	mouseover bool
	pressed   bool
}

var (
	_ furex.DrawHandler = (*Bar)(nil)
)

func (b *Bar) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	b.drawBlackBar(screen, frame)
	b.drawBar(screen, frame)
}

func (b *Bar) drawBlackBar(screen *ebiten.Image, frame image.Rectangle) {
	x, y := float64(frame.Min.X), float64(frame.Min.Y)
	spr := sprites.Get("barBack_horizontalLeft.png")
	ganim8.DrawSprite(screen, spr, 0, x, y, 0, 1, 1, 0, 0)
	x += float64(spr.W())

	spr = sprites.Get("barBack_horizontalMid.png")
	for x < float64(frame.Max.X)-float64(spr.W()) {
		ganim8.DrawSprite(screen, spr, 0, x, y, 0, 1, 1, 0, 0)
		x += float64(spr.W())
	}

	spr = sprites.Get("barBack_horizontalRight.png")
	ganim8.DrawSprite(screen, spr, 0, float64(frame.Max.X), y, 0, 1, 1, 1, 0)
}

func (b *Bar) drawBar(screen *ebiten.Image, frame image.Rectangle) {
	maxX := frame.Min.X + int(b.Value*float64(frame.Dx()))

	x, y := float64(frame.Min.X), float64(frame.Min.Y)
	spr := sprites.Get(fmt.Sprintf("bar%s_horizontalLeft.png", b.Color))
	ganim8.DrawSprite(screen, spr, 0, x, y, 0, 1, 1, 0, 0)
	x += float64(spr.W())

	spr = sprites.Get(fmt.Sprintf("bar%s_horizontalMid.png", b.Color))
	for x < float64(maxX)-float64(spr.W()) {
		ganim8.DrawSprite(screen, spr, 0, x, y, 0, 1, 1, 0, 0)
		x += float64(spr.W())
	}

	spr = sprites.Get(fmt.Sprintf("bar%s_horizontalRight.png", b.Color))
	ganim8.DrawSprite(screen, spr, 0, float64(maxX), y, 0, 1, 1, 1, 0)
}
