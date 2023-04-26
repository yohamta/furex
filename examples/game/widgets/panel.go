package widgets

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/sprites"
	"github.com/yohamta/furex/v2/examples/game/text"
	"github.com/yohamta/ganim8/v2"
)

type Panel struct {
	Color   color.Color
	OnClick func()

	mouseover bool
	pressed   bool
}

var (
	_ furex.ButtonHandler          = (*Panel)(nil)
	_ furex.NotButton              = (*Panel)(nil)
	_ furex.DrawHandlerWithView    = (*Panel)(nil)
	_ furex.MouseEnterLeaveHandler = (*Panel)(nil)
)

func (p *Panel) HandleDraw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	// This code is just for demo.
	// It's dirty and not optimized.

	PanelName := view.Attrs["sprite"]
	border := sprites.Get(fmt.Sprintf("%s_top_left", PanelName)).Width()
	top := sprites.Get(fmt.Sprintf("%s_top", PanelName)).Height()
	fborder := float64(border)

	spr := sprites.Get(fmt.Sprintf("%s_center", PanelName))
	x := float64(frame.Min.X) + fborder
	for x < float64(frame.Max.X)-fborder {
		y := float64(frame.Min.Y) + float64(top)
		for y < float64(frame.Max.Y)-fborder {
			opts := ganim8.DrawOpts(x, y, 0, 1, 1, 0, 0)
			p.drawSprite(screen, spr, opts)
			y += float64(spr.H())
		}
		x += float64(spr.W())
	}

	// top_left
	spr = sprites.Get(fmt.Sprintf("%s_top_left", PanelName))
	opts := ganim8.DrawOpts(float64(frame.Min.X), float64(frame.Min.Y), 0, 1, 1, 0, 0)
	p.drawSprite(screen, spr, opts)

	// top
	spr = sprites.Get(fmt.Sprintf("%s_top", PanelName))
	for x := float64(frame.Min.X + border); x < float64(frame.Max.X-border); x += float64(spr.W()) {
		opts := ganim8.DrawOpts(x, float64(frame.Min.Y), 0, 1, 1, 0, 0)
		p.drawSprite(screen, spr, opts)
	}
	// top_right
	spr = sprites.Get(fmt.Sprintf("%s_top_right", PanelName))
	opts = ganim8.DrawOpts(float64(frame.Max.X-border), float64(frame.Min.Y), 0, 1, 1, 0, 0)
	p.drawSprite(screen, spr, opts)
	// left
	spr = sprites.Get(fmt.Sprintf("%s_left", PanelName))
	for y := float64(frame.Min.Y + border); y < float64(frame.Max.Y-border); y += float64(spr.H()) {
		opts = ganim8.DrawOpts(float64(frame.Min.X), y, 0, 1, 1, 0, 0)
		p.drawSprite(screen, spr, opts)
	}
	// right
	spr = sprites.Get(fmt.Sprintf("%s_right", PanelName))
	for y := float64(frame.Min.Y + border); y < float64(frame.Max.Y-border); y += float64(spr.H()) {
		opts = ganim8.DrawOpts(float64(frame.Max.X-spr.W()), y, 0, 1, 1, 0, 0)
		p.drawSprite(screen, spr, opts)
	}
	// bottom_left
	spr = sprites.Get(fmt.Sprintf("%s_bottom_left", PanelName))
	opts = ganim8.DrawOpts(float64(frame.Min.X), float64(frame.Max.Y-border), 0, 1, 1, 0, 0)
	p.drawSprite(screen, spr, opts)
	// bottom
	spr = sprites.Get(fmt.Sprintf("%s_bottom", PanelName))
	for x := float64(frame.Min.X + border); x < float64(frame.Max.X-border); x += float64(spr.W()) {
		opts = ganim8.DrawOpts(x, float64(frame.Max.Y-spr.H()), 0, 1, 1, 0, 0)
		p.drawSprite(screen, spr, opts)
	}
	// bottom_right
	spr = sprites.Get(fmt.Sprintf("%s_bottom_right", PanelName))
	opts = ganim8.DrawOpts(float64(frame.Max.X-border), float64(frame.Max.Y-border), 0, 1, 1, 0, 0)
	p.drawSprite(screen, spr, opts)

	if view.Text != "" {
		x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)
		text.R.SetAlign(etxt.YCenter, etxt.XCenter)
		text.R.SetTarget(screen)
		if p.Color != nil {
			text.R.SetColor(p.Color)
		} else {
			text.R.SetColor(color.White)
		}
		text.R.Draw(view.Text, int(x), int(y))
	}
}

func (p *Panel) drawSprite(screen *ebiten.Image, spr *ganim8.Sprite, opts *ganim8.DrawOptions) {
	if p.IsButton() {
		if p.pressed {
			opts.ColorM.Scale(0.9, 0.9, 0.9, 1)
		} else if p.mouseover {
			opts.ColorM.Scale(1.1, 1.1, 1.1, 1)
		}
	}
	ganim8.DrawSpriteWithOpts(screen, spr, 0, opts, nil)
}

func (p *Panel) IsButton() bool {
	return p.OnClick != nil
}

func (p *Panel) HandleMouseEnter(x, y int) bool {
	p.mouseover = true
	return true
}

func (p *Panel) HandleMouseLeave() {
	p.mouseover = false
}

func (p *Panel) HandlePress(x, y int, t ebiten.TouchID) {
	p.pressed = true
}

func (p *Panel) HandleRelease(x, y int, isCancel bool) {
	p.pressed = false
	if !isCancel {
		if p.OnClick != nil {
			p.OnClick()
		}
	}
}
