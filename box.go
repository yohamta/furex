package furex

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Box struct {
	ViewEventHandlerFuncs

	w, h        int
	color       color.Color
	offsetImage *ebiten.Image
}

func NewBox(w, h int, clr color.Color) *Box {
	r := new(Box)

	r.w = w
	r.h = h
	r.color = clr

	return r
}

func (r *Box) OnLoad(v *View) {
	v.SetRect(0, 0, r.w, r.h)
	size := v.Rect().Size()
	offsetImage, _ := ebiten.NewImage(size.X, size.Y, ebiten.FilterDefault)
	FillRect(offsetImage, Rect{0, 0, size.X, size.Y}, r.color)
	r.offsetImage = offsetImage
}

func (r *Box) OnDraw(v *View, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(v.Rect().Min.X), float64(v.Rect().Min.Y))
	screen.DrawImage(r.offsetImage, op)
}
