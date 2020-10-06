package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// View represents a view with certain bounds
type View interface {
	// OnUpdate updates the content
	OnUpdate()

	// OnDraw renders the view to the screen
	OnDraw(screen *ebiten.Image, frame image.Rectangle)

	// GetStyle returns style data
	GetStyle() *Style

	// AddChild adds a child
	AddChild(v View)

	// GetChild returns children
	Children() []View
}

// ViewEmbed is a common implementation of a View
type ViewEmbed struct {
	Style

	children []View

	isDirty bool
}

func (v *ViewEmbed) GetStyle() *Style {
	return &v.Style
}

func (v *ViewEmbed) AddChild(child View) {
	v.children = append(v.children, child)
	v.isDirty = true
}

func (v *ViewEmbed) Children() []View {
	return v.children
}

func (v *ViewEmbed) OnLayout() {}

func (v *ViewEmbed) OnUpdate() {}

func (v *ViewEmbed) OnDraw(screen *ebiten.Image, frame image.Rectangle) {}
