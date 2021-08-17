package furex

import "github.com/hajimehoshi/ebiten/v2"

// ButtonComponent represents a button
type ButtonComponent interface {
	// OnPressButton will be called when the button is pressed
	HandlePress(t ebiten.TouchID)
	// OnReleaseButton will be called when the button is released
	HandleRelease(t ebiten.TouchID, isInside bool)
}
