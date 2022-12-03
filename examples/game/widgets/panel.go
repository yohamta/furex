package widgets

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/sprites"
	"github.com/yohamta/ganim8/v2"
)

type Panel struct {
	Sprite string
}

var _ furex.DrawHandler = (*Panel)(nil)

func (p *Panel) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	// This code is just for demo.
	// It's super dirty and not optimized.

	PanelName := p.Sprite
	border := sprites.Get(fmt.Sprintf("%s_top_left", PanelName)).Width()
	fborder := float64(border)

	spr := sprites.Get(fmt.Sprintf("%s_center", PanelName))
	x := float64(frame.Min.X) + fborder
	for x < float64(frame.Max.X)-fborder-2 {
		y := float64(frame.Min.Y) + fborder - 2
		for y < float64(frame.Max.Y)-fborder {
			ganim8.DrawSprite(screen, spr, 0, x, y, 0, 1, 1, 0, 0)
			y += float64(spr.H())
		}
		x += float64(spr.W())
	}

	// top_left
	spr = sprites.Get(fmt.Sprintf("%s_top_left", PanelName))
	ganim8.DrawSprite(screen, spr, 0, float64(frame.Min.X), float64(frame.Min.Y), 0, 1, 1, 0, 0)
	// top
	spr = sprites.Get(fmt.Sprintf("%s_top", PanelName))
	for x := float64(frame.Min.X + border); x < float64(frame.Max.X-border); x += float64(spr.W()) {
		ganim8.DrawSprite(screen, spr, 0, x, float64(frame.Min.Y), 0, 1, 1, 0, 0)
	}
	// top_right
	spr = sprites.Get(fmt.Sprintf("%s_top_right", PanelName))
	ganim8.DrawSprite(screen, spr, 0, float64(frame.Max.X-border), float64(frame.Min.Y), 0, 1, 1, 0, 0)
	// left
	spr = sprites.Get(fmt.Sprintf("%s_left", PanelName))
	for y := float64(frame.Min.Y + border); y < float64(frame.Max.Y-border); y += float64(spr.H()) {
		ganim8.DrawSprite(screen, spr, 0, float64(frame.Min.X), y, 0, 1, 1, 0, 0)
	}
	// right
	spr = sprites.Get(fmt.Sprintf("%s_right", PanelName))
	for y := float64(frame.Min.Y + border); y < float64(frame.Max.Y-border); y += float64(spr.H()) {
		ganim8.DrawSprite(screen, spr, 0, float64(frame.Max.X-spr.W()), y, 0, 1, 1, 0, 0)
	}
	// bottom_left
	spr = sprites.Get(fmt.Sprintf("%s_bottom_left", PanelName))
	ganim8.DrawSprite(screen, spr, 0, float64(frame.Min.X), float64(frame.Max.Y-border), 0, 1, 1, 0, 0)
	// bottom
	spr = sprites.Get(fmt.Sprintf("%s_bottom", PanelName))
	for x := float64(frame.Min.X + border); x < float64(frame.Max.X-border); x += float64(spr.W()) {
		ganim8.DrawSprite(screen, spr, 0, x, float64(frame.Max.Y-spr.H()), 0, 1, 1, 0, 0)
	}
	// bottom_right
	spr = sprites.Get(fmt.Sprintf("%s_bottom_right", PanelName))
	ganim8.DrawSprite(screen, spr, 0, float64(frame.Max.X-border), float64(frame.Max.Y-border), 0, 1, 1, 0, 0)
}
