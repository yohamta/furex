package furex

import "image"

// FixedSizeComponent represents a component with fixed size
type FixedSizeComponent interface {
	Component

	Size() image.Point
}

// AbsolutePositionComponent represents a component with fixed size
type AbsolutePositionComponent interface {
	FixedSizeComponent

	Position() image.Point
}
