package widgets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/sprites"
	"github.com/yohamta/ganim8/v2"
)

type Sprite struct {
}

var (
	_ furex.Drawer = (*Sprite)(nil)
)

func (t *Sprite) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	sprite := view.Attrs["sprite"]
	spr := sprites.Get(sprite)
	x, y := float64(frame.Min.X)+float64(frame.Dx())/2, float64(frame.Min.Y)+float64(frame.Dy())/2
	ganim8.DrawSprite(screen, spr, 0, x, y, 0, 1, 1, .5, .5)
}
