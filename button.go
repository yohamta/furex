package furex

// ButtonComponent represents a button
type ButtonComponent interface {
	// OnPressButton will be called when the button is pressed
	OnPressButton()
	// OnReleaseButton will be called when the button is released
	OnReleaseButton()
}
