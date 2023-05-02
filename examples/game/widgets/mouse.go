package widgets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/examples/game/sprites"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
)

type Mouse struct {
	isMouseActive bool
	x, y          int
}

var (
	_ furex.DrawHandler            = (*Mouse)(nil)
	_ furex.MouseHandler           = (*Mouse)(nil)
	_ furex.MouseEnterLeaveHandler = (*Mouse)(nil)
)

func (m *Mouse) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	if m.isMouseActive {
		spr := sprites.Get("cursorSword_gold.png")
		// set origin to .1, .1 to make the cursor point to the mouse position
		ganim8.DrawSprite(screen, spr, 0, float64(m.x), float64(m.y), 0, 1, 1, .1, .1)
	}
}

func (m *Mouse) HandleMouse(x, y int) bool {
	m.x = x
	m.y = y
	return true
}

func (m *Mouse) HandleMouseEnter(x int, y int) bool {
	m.isMouseActive = true
	m.x = x
	m.y = y
	return true
}

func (m *Mouse) HandleMouseLeave() {
	m.isMouseActive = false
}
