package shared

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Box struct {
	color color.Color
	size  image.Point
}

func NewBox(w, h int, clr color.Color) *Box {
	box := new(Box)
	box.size = image.Pt(w, h)
	box.color = clr
	return box
}

func (box *Box) Size() image.Point {
	return box.size
}

func (box *Box) Update() {}

func (box *Box) Draw(screen *ebiten.Image, frame image.Rectangle) {
	FillRect(screen, frame, box.color)
}
