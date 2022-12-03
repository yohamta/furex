package graphic

import (
	"image"
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	g    graphic
	once sync.Once
)

type graphic struct {
	imgOfAPixel *ebiten.Image
}

func (g *graphic) setup() {
	once.Do(func() {
		g.imgOfAPixel = ebiten.NewImage(1, 1)
	})
}

type FillRectOpts struct {
	Rect  image.Rectangle
	Color color.Color
}

func FillRect(target *ebiten.Image, opts *FillRectOpts) {
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

func DrawRect(target *ebiten.Image, opts *DrawRectOpts) {
	g.setup()
	r, c, sw := &opts.Rect, &opts.Color, opts.StrokeWidth
	FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Min.Y, r.Min.X+sw, r.Max.Y), Color: *c,
	})
	FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Min.Y, r.Min.X+sw, r.Max.Y), Color: *c,
	})
	FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Min.Y+sw), Color: *c,
	})
	FillRect(target, &FillRectOpts{
		Rect: image.Rect(r.Min.X, r.Max.Y-sw, r.Max.X, r.Max.Y), Color: *c,
	})
}
