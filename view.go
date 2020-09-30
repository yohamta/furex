package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type View interface {
	Children() []View
	AddChild(child View)

	SetBounds(x0, y0, x1, y1 int)
	SetPosition(x, y int)
	SetSize(w, h int)

	Bounds() image.Rectangle
	Position() image.Point
	Size() image.Point

	IsLoaded() bool
	SetLoaded(loaded bool)

	OnLoad()
	OnLayout()
	OnUpdate()
	OnDraw(screen *ebiten.Image)
}

type ViewEmbed struct {
	viewLoaded bool
	bounds     image.Rectangle
	children   []View
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

func (v *ViewEmbed) Bounds() image.Rectangle {
	return v.bounds
}

func (v *ViewEmbed) SetBounds(x0, y0, x1, y1 int) {
	v.bounds = image.Rect(x0, y0, x1, y1)
}

func (v *ViewEmbed) IsLoaded() bool {
	return v.viewLoaded
}

func (v *ViewEmbed) SetLoaded(loaded bool) {
	v.viewLoaded = loaded
}

func (v *ViewEmbed) AddChild(child View) {
	v.children = append(v.children, child)
}

func (v *ViewEmbed) Children() []View {
	return v.children
}

func (v *ViewEmbed) OnLayout() {}

func (v *ViewEmbed) OnLoad() {}

func (v *ViewEmbed) OnUpdate() {}

func (v *ViewEmbed) OnDraw(screen *ebiten.Image) {}
