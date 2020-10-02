package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// View represents a view with certain bounds
type View interface {
	// SetBounds sets bounds for this view
	SetBounds(x0, y0, x1, y1 int)

	// GetBounds returns bounds for this view
	GetBounds() image.Rectangle

	// Update updates the content
	Update()

	// Draw renders the view to the screen
	Draw(screen *ebiten.Image)
}

// ViewEmbed is a common implementation of a View
type ViewEmbed struct {
	bounds image.Rectangle
}

func (v *ViewEmbed) SetPosition(x, y int) {
	v.bounds.Add(image.Point{
		x - v.bounds.Min.X,
		y - v.bounds.Min.Y,
	})
}

func (v *ViewEmbed) SetSize(w, h int) {
	v.bounds.Max.X = w + v.bounds.Min.X
	v.bounds.Max.Y = h + v.bounds.Min.Y
}

func (v *ViewEmbed) Size() image.Point {
	return v.bounds.Size()
}

func (v *ViewEmbed) Position() image.Point {
	return v.bounds.Min
}

func (v *ViewEmbed) GetBounds() image.Rectangle {
	return v.bounds
}

func (v *ViewEmbed) SetBounds(x0, y0, x1, y1 int) {
	v.bounds = image.Rect(x0, y0, x1, y1)
}
