package furex

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Box display a box with filled color
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

func (box *Box) Update() {}

func (box *Box) Draw(screen *ebiten.Image, frame image.Rectangle) {
	p := frame.Min
	s := frame.Size()
	FillRect(screen, Rect{p.X, p.Y, s.X, s.Y}, box.color)
}
