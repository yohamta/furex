package furex

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var G graphic

type graphic struct {
	init        bool
	imgOfAPixel *ebiten.Image
}

func (g *graphic) setup() {
	if !g.init {
		g.imgOfAPixel = ebiten.NewImage(1, 1)
	}
}

type FillRectOpts struct {
	Rect  image.Rectangle
	Color color.Color
}

func (g *graphic) FillRect(target *ebiten.Image, opts *FillRectOpts) {
	g.setup()
	r, c := &opts.Rect, &opts.Color
	g.imgOfAPixel.Fill(*c)
	op := &ebiten.DrawImageOptions{}
	w, h := r.Size().X, r.Size().Y
	op.GeoM.Translate(float64(r.Min.X)*(1/float64(w)), float64(r.Min.Y)*(1/float64(h)))
	op.GeoM.Scale(float64(w), float64(h))
	target.DrawImage(g.imgOfAPixel, op)
}

type DrawRectOpts struct {
	Rect        image.Rectangle
	Color       color.Color
	StrokeWidth int
}

func (g *graphic) DrawRect(target *ebiten.Image, opts *DrawRectOpts) {
	g.setup()
	r, c, sw := &opts.Rect, &opts.Color, opts.StrokeWidth
	g.FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Min.Y, r.Min.X+sw, r.Max.Y), Color: *c,
	})
	g.FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Min.Y, r.Min.X+sw, r.Max.Y), Color: *c,
	})
	g.FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Min.Y+sw), Color: *c,
	})
	g.FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Max.Y-sw, r.Max.X, r.Max.Y), Color: *c,
	})
}
