package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// View represents a view with certain bounds
type View interface {
	// OnLayout layouts the content
	OnLayout()

	// OnUpdate updates the content
	OnUpdate()

	// OnDraw renders the view to the screen
	OnDraw(screen *ebiten.Image, frame image.Rectangle)

	// GetStyle returns style data
	GetStyle() *Style

	// Layout layouts the content
	Layout()

	// Update updates the content
	Update()

	// Draw renders the view to the screen
	Draw(screen *ebiten.Image, frame image.Rectangle)
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

func (v *ViewEmbed) Layout() {}

func (v *ViewEmbed) Update() {
	if v.isDirty {
		v.OnLayout()
	}
	v.OnUpdate()
	for i := 0; i < len(v.children); i++ {
		v.children[i].Update()
	}
}

func (v *ViewEmbed) Draw(screen *ebiten.Image, frame image.Rectangle) {
	v.OnDraw(screen, frame)
	for i := 0; i < len(v.children); i++ {
		child := v.children[i]
		child.Draw(screen, child.GetStyle().Bounds.Add(v.GetStyle().Bounds.Min))
	}
}

func (v *ViewEmbed) AddChild(child View) {
	v.children = append(v.children, child)
	v.isDirty = true
}

func (v *ViewEmbed) OnLayout() {}

func (v *ViewEmbed) OnUpdate() {}

func (v *ViewEmbed) OnDraw(screen *ebiten.Image, frame image.Rectangle) {}
