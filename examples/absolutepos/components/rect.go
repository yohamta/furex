package components

import (
	"image"
	"image/color"

	"github.com/yotahamada/furex"
)

type Rect struct {
	furex.Box
	pos image.Point
}

func NewRect(x, y, w, h int, clr color.Color) *Rect {
	r := &Rect{Box: *furex.NewBox(w, h, clr)}
	r.pos = image.Pt(x, y)
	return r
}

func (r *Rect) GetPosition() image.Point {
	return r.pos
}
