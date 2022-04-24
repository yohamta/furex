package furex

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type View struct {
	Left         int
	Top          int
	Width        int
	Height       int
	MarginLeft   int
	MarginTop    int
	MarginRight  int
	MarginBottom int
	Position     Position
	Handler      DrawHandler
	Direction    Direction
	Wrap         FlexWrap
	Justify      Justify
	AlignItems   AlignItem
	AlignContent AlignContent

	containerEmbed
	flexEmbed
	hasParent bool
}

func (v *View) Update() {
	if v.isDirty {
		if !v.hasParent {
			v.frame = image.Rect(v.Left, v.Top, v.Left+v.Width, v.Top+v.Height)
		}
		v.flexEmbed.View = v
		v.layout(v.Width, v.Height, &v.containerEmbed)
		v.isDirty = false
	}
	for _, v := range v.children {
		v.item.Update()
		u, ok := v.item.Handler.(UpdateHandler)
		if ok && u != nil {
			u.HandleUpdate()
		}
	}
	if !v.hasParent {
		v.processEvent()
	}
}

func (v *View) Draw(screen *ebiten.Image) {
	v.containerEmbed.Draw(screen)
}

func (v *View) AddChild(cv *View) {
	child := &child{item: cv, handledTouchID: -1}
	v.children = append(v.children, child)
	v.isDirty = true
	cv.hasParent = true
}

type child struct {
	item                     *View
	bounds                   image.Rectangle
	isButtonPressed          bool
	isMouseLeftButtonHandler bool
	handledTouchID           ebiten.TouchID
}
