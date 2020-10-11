package furex

// Container represents a container that can have child components
type Container interface {
	Component

	AddChild(child Component)
}
