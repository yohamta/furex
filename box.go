package furex

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Box struct {
	ViewEmbed
	color color.Color
}

func NewBox(w, h int, clr color.Color) *Box {
	box := new(Box)
	box.SetSize(w, h)
	box.color = clr
	return box
}

func (box *Box) OnDraw(screen *ebiten.Image) {
	pos := box.Position()
	size := box.Size()
	FillRect(screen, Rect{pos.X, pos.Y, size.X, size.Y}, box.color)
}
