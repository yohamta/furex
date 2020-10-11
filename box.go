package furex

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Box display a box with filled color
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

func (box *Box) GetSize() image.Point {
	return box.size
}

func (box *Box) Update() {}

func (box *Box) Draw(screen *ebiten.Image, frame image.Rectangle) {
	FillRect(screen, frame, box.color)
}
