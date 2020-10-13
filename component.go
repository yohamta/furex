package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Component represents a component of the UI
type Component interface {
	Update()

	Draw(screen *ebiten.Image, frame image.Rectangle)
}

// FixedSizeComponent represents a component with fixed size
// This interface should be implemented for flex child
type FixedSizeComponent interface {
	Component

	GetSize() image.Point
}

// ButtonComponent represents a button
type ButtonComponent interface {
	OnPressButton()
	OnReleaseButton()
}
