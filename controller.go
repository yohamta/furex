package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type ViewController struct {
	root View
}

func NewViewController() *ViewController {
	vc := new(ViewController)
	return vc
}

func (vc *ViewController) SetRootView(v View) {
	vc.root = v
}

func (vc *ViewController) Update() {
	var f func(v View)
	f = func(v View) {
		// TODO: handle touch if it is a button
		v.OnUpdate()
		for c := range v.Children() {
			f(v.Children()[c])
		}
	}
	f(vc.root)
}

func (vc *ViewController) Draw(screen *ebiten.Image, frame image.Rectangle) {
	var f func(v View, offset image.Point)
	f = func(v View, offset image.Point) {
		v.OnDraw(screen, v.GetStyle().Bounds.Add(offset))
		for c := range v.Children() {
			f(v.Children()[c], v.GetStyle().Bounds.Min)
		}
	}
	f(vc.root, image.Pt(0, 0))
}
