package components

import (
	"image"
	"image/color"

	"github.com/yohamta/furex/examples/shared"
)

type Rect struct {
	shared.Box
	pos image.Point
}

func NewRect(x, y, w, h int, clr color.Color) *Rect {
	r := &Rect{Box: *shared.NewBox(w, h, clr)}
	r.pos = image.Pt(x, y)
	return r
}

func (r *Rect) Position() image.Point {
	return r.pos
}
