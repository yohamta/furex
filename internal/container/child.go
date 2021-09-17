package container

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Child struct {
	Item                    interface{}
	Bounds                  image.Rectangle
	IsButtonPressed         bool
	IsMouseLeftClickHandler bool
	HandledTouchID          ebiten.TouchID
}

func NewChild(item interface{}) *Child {
	return &Child{
		Item:                    item,
		HandledTouchID:          -1,
		IsButtonPressed:         false,
		IsMouseLeftClickHandler: false,
	}
}
