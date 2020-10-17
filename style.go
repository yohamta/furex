package furex

import "image"

// FixedSizeComponent represents a component with fixed size
type FixedSizeComponent interface {
	Component

	GetSize() image.Point
}

// AbsolutePositionComponent represents a component with fixed size
type AbsolutePositionComponent interface {
	FixedSizeComponent

	GetPosition() image.Point
}
