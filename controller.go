package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// ViewController represents UI controller
type ViewController struct {
	frame    image.Rectangle
	rootView *View
}

// NewViewController creates a new controller
func NewViewController() *ViewController {
	c := &ViewController{}

	return c
}

// SetFrame sets the frame of the view
func (c *ViewController) SetFrame(x, y, w, h int) {
	c.frame = image.Rect(x, y, x+w, y+h)
	if c.rootView != nil {
		c.rootView.SetPosition(x, y)
		c.rootView.SetSize(w, h)
	}
}

// SetRootView sets the frame of the view
func (c *ViewController) SetRootView(v *View) {
	c.rootView = v
	c.rootView.SetPosition(c.frame.Min.X, c.frame.Min.Y)
	c.rootView.SetSize(c.frame.Size().X, c.frame.Size().Y)
	c.rootView.Load()
}

// Update updates the ui
func (c *ViewController) Update() {
	c.rootView.Update()
}

// Layout updates the layout
func (c *ViewController) Layout() {
	c.rootView.Layout()
}

// Draw draws the ui
func (c *ViewController) Draw(screen *ebiten.Image) {
	c.rootView.Draw(screen)
}
