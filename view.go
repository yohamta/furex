package furex

import (
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// View represents a UI element.
// You can set flex options, size, position and so on.
// Handlers can be set to create custom component such as button or list.
type View struct {
	// TODO: Remove these fields in the future.
	ID           string
	Left         int
	Top          int
	Width        int
	Height       int
	MarginLeft   int
	MarginTop    int
	MarginRight  int
	MarginBottom int
	Position     Position
	Direction    Direction
	Wrap         FlexWrap
	Justify      Justify
	AlignItems   AlignItem
	AlignContent AlignContent
	Grow         float64
	Shrink       float64
	Raw          string

	Handler DrawHandler

	containerEmbed
	flexEmbed
	hasParent bool
	lock      sync.Mutex
}

// Update updates the view
func (v *View) Update() {
	if v.isDirty {
		v.startLayout()
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

func (v *View) startLayout() {
	v.lock.Lock()
	defer v.lock.Unlock()
	if !v.hasParent {
		v.frame = image.Rect(v.Left, v.Top, v.Left+v.Width, v.Top+v.Height)
	}
	v.flexEmbed.View = v

	for _, child := range v.children {
		if child.item.Position == PositionStatic {
			child.item.startLayout()
		}
	}

	v.layout(v.frame.Dx(), v.frame.Dy(), &v.containerEmbed)
	v.isDirty = false
}

// UpdateWithSize the view with modified height and width
func (v *View) UpdateWithSize(width, height int) {
	if !v.hasParent && (v.Width != width || v.Height != height) {
		v.Height = height
		v.Width = width
		v.isDirty = true
	}
	v.Update()
}

// Draw draws the view
func (v *View) Draw(screen *ebiten.Image) {
	if v.isDirty {
		v.startLayout()
	}
	v.containerEmbed.Draw(screen)
	if Debug && !v.hasParent {
		debugBorders(screen, v.containerEmbed)
	}
}

// AddTo add itself to a parent view
func (v *View) AddTo(parent *View) *View {
	if v.hasParent {
		panic("this view has been already added to a parent")
	}
	parent.AddChild(v)
	return v
}

// AddChild adds one or multiple child views
func (v *View) AddChild(views ...*View) *View {
	for _, vv := range views {
		v.addChild(vv)
	}
	return v
}

// RemoveChild removes a specified view
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

// RemoveAll removes all children view
func (v *View) RemoveAll() {
	v.isDirty = true
	for _, child := range v.children {
		child.item.hasParent = false
	}
	v.children = []*child{}
}

// PopChild remove the last child view add to this view
func (v *View) PopChild() *View {
	if len(v.children) == 0 {
		return nil
	}
	c := v.children[len(v.children)-1]
	v.children = v.children[:len(v.children)-1]
	v.isDirty = true
	c.item.hasParent = false
	return c.item
}

func (v *View) addChild(cv *View) *View {
	child := &child{item: cv, handledTouchID: -1}
	v.children = append(v.children, child)
	v.isDirty = true
	cv.hasParent = true
	return v
}

func (v *View) width() int {
	if v.Width == 0 {
		return v.calculatedWidth
	}
	return v.Width
}

func (v *View) height() int {
	if v.Height == 0 {
		return v.calculatedHeight
	}
	return v.Height
}

func (v *View) getChildren() []*View {
	if v == nil || v.children == nil {
		return nil
	}
	ret := make([]*View, len(v.children))
	for i, child := range v.children {
		ret[i] = child.item
	}
	return ret
}

// This is for debugging and testing.
type ViewConfig struct {
	Left         int
	Top          int
	Width        int
	Height       int
	MarginLeft   int
	MarginTop    int
	MarginRight  int
	MarginBottom int
	Position     Position
	Direction    Direction
	Wrap         FlexWrap
	Justify      Justify
	AlignItems   AlignItem
	AlignContent AlignContent
	Grow         float64
	Shrink       float64
	children     []*ViewConfig
}

func (v *View) Config() *ViewConfig {
	cfg := &ViewConfig{
		Left:         v.Left,
		Top:          v.Top,
		Width:        v.Width,
		Height:       v.Height,
		MarginLeft:   v.MarginLeft,
		MarginTop:    v.MarginTop,
		MarginRight:  v.MarginRight,
		MarginBottom: v.MarginBottom,
		Position:     v.Position,
		Direction:    v.Direction,
		Wrap:         v.Wrap,
		Justify:      v.Justify,
		AlignItems:   v.AlignItems,
		AlignContent: v.AlignContent,
		Grow:         v.Grow,
		Shrink:       v.Shrink,
		children:     []*ViewConfig{},
	}
	for _, child := range v.getChildren() {
		cfg.children = append(cfg.children, child.Config())
	}
	return cfg
}
