package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type ViewController struct {
	frame    image.Rectangle
	rootView View
}

func NewViewController(x, y, w, h int) *ViewController {
	c := &ViewController{}

	c.frame = image.Rect(x, y, x+w, y+h)
	c.SetRootView(NewFlex(w, h))

	return c
}

func (c *ViewController) SetRootView(view View) {
	c.rootView = view
	c.rootView.SetPosition(c.frame.Min.X, c.frame.Min.Y)
	c.rootView.SetSize(c.frame.Size().X, c.frame.Size().Y)
	c.rootView.OnLoad()
}

func (c *ViewController) RootView() View {
	return c.rootView
}

func (c *ViewController) Load() {
	var f func(v View)
	f = func(v View) {
		if !v.IsLoaded() {
			v.OnLoad()
			v.SetLoaded(true)
		}
		for i := 0; i < len(v.Children()); i++ {
			child := v.Children()[i]
			if !child.IsLoaded() {
				child.Children()[i].OnLoad()
				child.SetLoaded(true)
			}
			f(child)
		}
	}
	f(c.rootView)
}

func (c *ViewController) Layout() {
	var f func(v View)
	f = func(v View) {
		v.OnLayout()
		for i := 0; i < len(v.Children()); i++ {
			child := v.Children()[i]
			child.OnLayout()
			f(child)
		}
	}
	f(c.rootView)
}

func (c *ViewController) Update() {
	var f func(v View)
	f = func(v View) {
		v.OnUpdate()
		for i := 0; i < len(v.Children()); i++ {
			child := v.Children()[i]
			child.OnUpdate()
			f(child)
		}
	}
	f(c.rootView)
}

func (c *ViewController) Draw(screen *ebiten.Image) {
	var f func(v View)
	f = func(v View) {
		v.OnDraw(screen)
		for i := 0; i < len(v.Children()); i++ {
			child := v.Children()[i]
			child.OnDraw(screen)
			f(child)
		}
	}
	f(c.rootView)
}
