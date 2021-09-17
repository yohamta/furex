// Referenced code: https://github.com/golang/exp/blob/master/shiny/widget/flex/flex.go
package furex_test

import (
	"image"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
	"golang.org/x/exp/shiny/widget/flex"
)

func TestFlexAlignments(t *testing.T) {
	flexSize := image.Pt(100, 100)
	itemSize := image.Pt(50, 50)

	var tests = []struct {
		name string
		a    image.Point
		b    image.Point
		c    furex.Direction
		d    furex.Justify
		e    furex.AlignItem
		want image.Rectangle
	}{
		{
			name: "Column - Center, Center",
			a:    flexSize,
			b:    itemSize,
			c:    furex.Column,
			d:    furex.JustifyCenter,
			e:    furex.AlignItemCenter,
			want: image.Rect(25, 25, 75, 75),
		},
		{
			name: "Column - Start, End",
			a:    flexSize,
			b:    itemSize,
			c:    furex.Column,
			d:    furex.JustifyStart,
			e:    furex.AlignItemEnd,
			want: image.Rect(50, 0, 100, 50),
		},
		{
			name: "Row - Center, Center",
			a:    flexSize,
			b:    itemSize,
			c:    furex.Row,
			d:    furex.JustifyCenter,
			e:    furex.AlignItemCenter,
			want: image.Rect(25, 25, 75, 75),
		},
		{
			name: "Row - End, Start",
			a:    flexSize,
			b:    itemSize,
			c:    furex.Row,
			d:    furex.JustifyEnd,
			e:    furex.AlignItemStart,
			want: image.Rect(50, 0, 100, 50),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flexItemBounds(
				tt.a, tt.b, flex.Direction(tt.c), flex.Justify(tt.d), flex.AlignItem(tt.e),
			)
			if got != tt.want {
				t.Errorf("flexItemBounds(%s): got %v; want %v", tt.name, got, tt.want)
			}

		})
	}
}

func TestChangeFrame(t *testing.T) {
	flexSize := image.Pt(100, 150)
	itemSize := image.Pt(30, 40)

	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Column
	flex.Justify = furex.JustifyCenter
	flex.AlignItems = furex.AlignItemCenter
	flex.SetFrame(image.Rect(100, 50, 100+flexSize.X, 50+flexSize.Y))

	item := NewMockItem(itemSize.X, itemSize.Y)
	flex.AddChild(item)
	flex.Update()
	flex.Draw(nil, image.Rect(0, 0, 0, 0))

	//  (0,0)
	//  ┌───────────────────────────────────┐
	//  │                                   │
	//  │         (100,50)                  │
	//  │           ┌───────────────────────┤
	//  │           │ flex                  │
	//  │           │                       │
	//  │           │                       │
	//  │           │                       │
	//  │           │                       │
	//  │           │           item(30x40) │
	//  │           │      ┌─────────┐      │
	//  │           │      │         │      │
	//  │           │      │         │      │
	//  │           │      │  item   │      │
	//  │           │      │         │      │
	//  │           │      │         │      │
	//  │           │      └─────────┘      │
	//  │           │                       │
	//  │           │                       │
	//  │           │                       │
	//  │           │                       │
	//  │           │                       │
	//  └───────────┴───────────────────────┘
	//                                  (150,200)
	//  expected item frame:
	//  x = (100-30)/2 + 100(container's frame) = 135
	//  y = (150-40)/2 + 50 (container's frame) = 105

	want := image.Rect(135, 105, 135+30, 105+40)
	got := item.frame
	if got != want {
		t.Errorf("TestChangeFrame: got %v; want %v", got, want)
	}
}

func TestNesting(t *testing.T) {
	// parent
	flexSize := image.Pt(300, 500)
	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Column
	flex.Justify = furex.JustifyCenter
	flex.AlignItems = furex.AlignItemCenter
	flex.SetFrame(image.Rect(100, 50, 100+flexSize.X, 50+flexSize.Y))

	// child
	flexSize2 := image.Pt(100, 200)
	inner1 := furex.NewFlex(flexSize2.X, flexSize2.Y)
	inner1.Direction = furex.Column
	inner1.Justify = furex.JustifyEnd
	inner1.AlignItems = furex.AlignItemEnd
	flex.AddChild(inner1)

	// add item into the child flex
	itemSize := image.Pt(30, 40)
	item := NewMockItem(itemSize.X, itemSize.Y)
	inner1.AddChild(item)

	// execute layout & draw
	flex.Update()
	flex.Draw(nil, image.Rect(0, 0, 0, 0))

	// 	(0,0)
	// ┌───────────────────────────────────┐
	// │ view                              │
	// │      (100,50)                     │
	// │      ┌────────────────────────────┤
	// │      │flex(300x500)               │
	// │      │                            │
	// │      │                            │
	// │      │     (200,200)              │
	// │      │     ┌─────────────────┐    │
	// │      │     │flex2(100x200)   │    │
	// │      │     │                 │    │
	// │      │     │    ┌────────────┤    │
	// │      │     │    │item(30x40) │    │
	// │      │     │    │            │    │
	// │      │     │    │            │    │
	// │      │     │    │            │    │
	// │      │     │    │            │    │
	// │      │     │    │            │    │
	// │      │     └────┴────────────┘    │
	// │      │                  (300,400) │
	// │      │                            │
	// │      │                            │
	// └──────┴────────────────────────────┘
	//                                 (400,550)
	// expected item frame:
	// x = 300-30 = 270 to 300
	// y = 400-40 = 360 to 400

	want := image.Rect(270, 360, 300, 400)
	got := item.frame
	if got != want {
		t.Errorf("TestNesting: got %v; want %v", got, want)
	}
}

func TestButtonTouch(t *testing.T) {
	// expected button frame:
	// x = 300-10 = 290 to 300
	// y = 400-20 = 380 to 400
	type result struct {
		isPressed  bool
		isReleased bool
		isInside   bool
	}
	var tests = []struct {
		name string
		a    image.Point
		b    image.Point
		want result
	}{
		{
			name: "press inside left-top edge, release inside",
			a:    image.Pt(290, 380),
			b:    image.Pt(290, 380),
			want: result{true, true, true},
		},
		{
			name: "press inside left-top edge, release outside",
			a:    image.Pt(290, 380),
			b:    image.Pt(290, 379),
			want: result{true, true, false},
		},
		{
			name: "press inside righ-bottom edge, release inside",
			a:    image.Pt(300, 400),
			b:    image.Pt(300, 400),
			want: result{true, true, true},
		},
		{
			name: "press inside righ-bottom edge, release outside",
			a:    image.Pt(300, 400),
			b:    image.Pt(301, 400),
			want: result{true, true, false},
		},
		{
			name: "press outside, release inside",
			a:    image.Pt(289, 390),
			b:    image.Pt(295, 390),
			want: result{false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := _TestButtonTouch(tt.a, tt.b)
			got := result{b.isPressed, b.isReleased, b.isInside}
			if got != tt.want {
				t.Errorf("TestButtonTouch(%s): got %v; want %v", tt.name, got, tt.want)
			}

		})
	}
}

func _TestButtonTouch(pressedPos image.Point, releasedPos image.Point) *MockButton {
	// parent
	flexSize := image.Pt(300, 500)
	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Column
	flex.Justify = furex.JustifyCenter
	flex.AlignItems = furex.AlignItemCenter
	flex.SetFrame(image.Rect(100, 50, 100+flexSize.X, 50+flexSize.Y))

	// child
	flexSize2 := image.Pt(100, 200)
	inner1 := furex.NewFlex(flexSize2.X, flexSize2.Y)
	inner1.Direction = furex.Column
	inner1.Justify = furex.JustifyEnd
	inner1.AlignItems = furex.AlignItemEnd
	flex.AddChild(inner1)

	// add item into the child flex
	buttonSize := image.Pt(10, 20)
	button := NewMockButton(buttonSize.X, buttonSize.Y)
	inner1.AddChild(button)

	// execute layout & draw
	flex.Update()
	flex.Draw(nil, image.Rect(0, 0, 0, 0))

	// 	(0,0)
	// ┌───────────────────────────────────┐
	// │ view                              │
	// │      (100,50)                     │
	// │      ┌────────────────────────────┤
	// │      │flex(300x500)               │
	// │      │                            │
	// │      │                            │
	// │      │     (200,200)              │
	// │      │     ┌─────────────────┐    │
	// │      │     │flex2(100x200)   │    │
	// │      │     │                 │    │
	// │      │     │   ┌──────-──────┤    │
	// │      │     │   │button(10x20)│    │
	// │      │     │   │             │    │
	// │      │     │   │             │    │
	// │      │     │   │             │    │
	// │      │     │   │             │    │
	// │      │     │   │             │    │
	// │      │     └───┴──────────-──┘    │
	// │      │                  (300,400) │
	// │      │                            │
	// │      │                            │
	// └──────┴────────────────────────────┘
	//                                 (400,550)
	// expected button frame:
	// x = 300-10 = 290 to 300
	// y = 400-20 = 380 to 400

	flex.HandleJustPressedTouchID(0, pressedPos.X, pressedPos.Y)
	flex.HandleJustReleasedTouchID(0, releasedPos.X, releasedPos.Y)

	return button
}

func flexItemBounds(flexSize image.Point, itemSize image.Point, direction flex.Direction, justify flex.Justify, alignItem flex.AlignItem) image.Rectangle {
	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Direction(direction)
	flex.Justify = furex.Justify(justify)
	flex.AlignItems = furex.AlignItem(alignItem)

	item := NewMockItem(itemSize.X, itemSize.Y)
	flex.AddChild(item)
	flex.Update()
	flex.Draw(nil, image.Rect(0, 0, 0, 0))

	return item.frame
}

type MockItem struct {
	size  image.Point
	frame image.Rectangle
}

func NewMockItem(w, h int) *MockItem {
	m := new(MockItem)
	m.size = image.Pt(w, h)
	return m
}

func (m *MockItem) Size() image.Point {
	return m.size
}

func (m *MockItem) Draw(screen *ebiten.Image, frame image.Rectangle) {
	m.frame = frame
}

type MockButton struct {
	size       image.Point
	frame      image.Rectangle
	isPressed  bool
	isReleased bool
	isInside   bool
}

func NewMockButton(w, h int) *MockButton {
	m := new(MockButton)
	m.size = image.Pt(w, h)
	m.isInside = false
	m.isPressed = false
	m.isReleased = false
	return m
}

func (m *MockButton) Size() image.Point {
	return m.size
}

func (m *MockButton) Draw(screen *ebiten.Image, frame image.Rectangle) {
	m.frame = frame
}

func (m *MockButton) HandlePress() {
	m.isPressed = true
}

func (m *MockButton) HandleRelease(isInside bool) {
	m.isReleased = true
	m.isInside = isInside
}
