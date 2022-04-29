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
	Grow         float64
	Shrink       float64

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
		v.layout(v.frame.Dx(), v.frame.Dy(), &v.containerEmbed)
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

func (v *View) AddChild(cv *View) *View {
	child := &child{item: cv, handledTouchID: -1}
	v.children = append(v.children, child)
	v.isDirty = true
	cv.hasParent = true
	return v
}

func (v *View) AddTo(pv *View) *View {
	if v.hasParent {
		panic("this view has been already added to a parent")
	}
	pv.AddChild(v)
	return v
}

func (v *View) AddChildren(views ...*View) *View {
	for _, vv := range views {
		v.AddChild(vv)
	}
	return v
}

func (v *View) RemoveChild(cv *View) bool {
	for i, child := range v.children {
		if child.item == cv {
			v.children = append(v.children[:i], v.children[i+1:]...)
			v.isDirty = true
			cv.hasParent = false
			return true
		}
	}
	return false
}

func (v *View) RemoveAll() {
	v.isDirty = true
	for _, child := range v.children {
		child.item.hasParent = false
	}
	v.children = []*child{}
}

func (v *View) PopChild() *View {
	if len(v.children) == 0 {
		return nil
	}
	c := v.children[len(v.children)-1]
	v.children = v.children[:len(v.children)-1]
	v.isDirty = true
	return c.item
}

type child struct {
	item                     *View
	bounds                   image.Rectangle
	isButtonPressed          bool
	isMouseLeftButtonHandler bool
	handledTouchID           ebiten.TouchID
}
