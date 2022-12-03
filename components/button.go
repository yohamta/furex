package components

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/furex/v2"
)

type Button struct {
	mouseover bool
	pressed   bool
	Text      string
	OnClick   func()
}

var _ furex.ButtonHandler = (*Button)(nil)
var _ furex.DrawHandler = (*Button)(nil)
var _ furex.MouseEnterLeaveHandler = (*Button)(nil)

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
	if b.pressed {
		furex.G.FillRect(screen, &furex.FillRectOpts{
			Rect: frame, Color: color.RGBA{0xaa, 0, 0, 0xff},
		})
	} else if b.mouseover {
		furex.G.FillRect(screen, &furex.FillRectOpts{
			Rect: frame, Color: color.RGBA{0xaa, 0xaa, 0, 0xff},
		})
	} else {
		furex.G.FillRect(screen, &furex.FillRectOpts{
			Rect: frame, Color: color.RGBA{0, 0xaa, 0, 0xff},
		})
	}
	ebitenutil.DebugPrintAt(screen, b.Text,
		frame.Min.X+((frame.Dx()-36)/2), frame.Min.Y+frame.Dy()/2-8)
}

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.mouseover = true
	return true
}

func (b *Button) HandleMouseLeave() {
	b.mouseover = false
}
