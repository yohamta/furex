package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Component represents a component of the UI
type Component interface {
	Update()

	Draw(screen *ebiten.Image, frame image.Rectangle)
}
