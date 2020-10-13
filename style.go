package furex

import "image"

// FixedSizeComponent represents a component with fixed size
// This interface should be implemented for flex child
type FixedSizeComponent interface {
	Component

	GetSize() image.Point
}
