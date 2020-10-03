package furex

import (
	"github.com/hajimehoshi/ebiten"
)

// View represents a view with certain bounds
type View interface {
	// Layout layouts the content
	Layout()

	// Update updates the content
	Update()

	// Draw renders the view to the screen
	Draw(screen *ebiten.Image)

	// GetStyle returns style data
	GetStyle() *Style
}

// ViewEmbed is a common implementation of a View
type ViewEmbed struct {
	Style
}

func (v *ViewEmbed) GetStyle() *Style {
	return &v.Style
}

func (v *ViewEmbed) Layout() {}

func (v *ViewEmbed) Update() {}

func (v *ViewEmbed) Draw(screen *ebiten.Image) {}
