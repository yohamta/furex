package widgets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/sprites"
	"github.com/yohamta/ganim8/v2"
)

type Button struct {
	Text          string
	OnClick       func()
	Sprite        string
	SpritePressed string

	mouseover bool
	pressed   bool
}

var (
	_ furex.ButtonHandler          = (*Button)(nil)
	_ furex.DrawHandler            = (*Button)(nil)
	_ furex.MouseEnterLeaveHandler = (*Button)(nil)
)

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.pressed = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.pressed = false
	if !isCancel {
		if b.OnClick != nil {
			b.OnClick()
		}
	}
}

func (b *Button) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	x, y := float64(frame.Min.X+frame.Dx()/2), float64(frame.Min.Y+frame.Dy()/2)

	opts := ganim8.DrawOpts(x, y, 0, 1, 1, .5, .5)
	if b.mouseover {
		opts.ColorM.Scale(1.1, 1.1, 1.1, 1)
	}
	if b.pressed && b.SpritePressed != "" {
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(b.SpritePressed), 0, opts, nil)
	} else if b.Sprite != "" {
		ganim8.DrawSpriteWithOpts(screen, sprites.Get(b.Sprite), 0, opts, nil)
	}

	ebitenutil.DebugPrintAt(screen, b.Text,
		frame.Min.X+((frame.Dx()-len(b.Text)*6)/2), frame.Min.Y+frame.Dy()/2-8)
}

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.mouseover = true
	return true
}

func (b *Button) HandleMouseLeave() {
	b.mouseover = false
}
