package furex_test

import (
	"image"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miyahoyo/furex"
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
				t.Errorf("TestFlexAlignments(%s): got %v; want %v", tt.name, got, tt.want)
			}

		})
	}
}

func TestFlexWrap(t *testing.T) {
	flexSize := image.Pt(200, 200)
	itemSize := image.Pt(100, 100)

	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Row
	flex.Justify = furex.JustifyStart
	flex.AlignItems = furex.AlignItemStart
	flex.Wrap = furex.Wrap

	item1 := NewMockItem(itemSize.X, itemSize.Y)
	flex.AddChild(item1)

	item2 := NewMockItem(itemSize.X, itemSize.Y)
	flex.AddChild(item2)

	item3 := NewMockItem(itemSize.X, itemSize.Y)
	flex.AddChild(item3)

	flex.Update()
	flex.Draw(nil)

	// (0,0)
	// ┌───────────────(100,0)───────────┐
	// │box1            │box2            │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// (0,100)──────────┼────────────(200,100)
	// │box3            │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// │                │                │
	// └──────────────(100,200)──────────┘
	// 															 (200,200)

	want := image.Rect(0, 100, 100, 200)
	got := item3.frame
	if got != want {
		t.Errorf("TestFlexWrap: got %v; want %v", got, want)
	}
}

func TestFlexFrameChange(t *testing.T) {
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
	flex.Draw(nil)

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
		t.Errorf("TestFlexFrameChange: got %v; want %v", got, want)
	}
}

func TestFlexNesting(t *testing.T) {
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
	flex.AddChildContainer(inner1)

	// add item into the child flex
	itemSize := image.Pt(30, 40)
	item := NewMockItem(itemSize.X, itemSize.Y)
	inner1.AddChild(item)

	// execute layout & draw
	flex.Update()
	flex.Draw(nil)

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

func TestMarginedItem(t *testing.T) {
	flexSize := image.Pt(300, 300)
	itemSize := image.Pt(100, 100)

	var tests = []struct {
		name string
		a    image.Point
		b    image.Point
		c    []int
		d    furex.Direction
		e    furex.Justify
		f    furex.AlignItem
		want image.Rectangle
	}{
		{
			name: "Row, Center, Margin Top",
			a:    flexSize,
			b:    itemSize,
			c:    []int{100, 0, 0, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 150, 200, 250),
		},
		{
			name: "Row, Center, Margin Top = Margin Bottom",
			a:    flexSize,
			b:    itemSize,
			c:    []int{100, 0, 100, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 100, 200, 200),
		},
		{
			name: "Row, Center, Margin Bottom",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 100, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 50, 200, 150),
		},
		{
			name: "Row, Start, Margin Top",
			a:    flexSize,
			b:    itemSize,
			c:    []int{50, 0, 0, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemStart,
			want: image.Rect(100, 50, 200, 150),
		},
		{
			name: "Row, End, Margin Bottom",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 50, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemEnd,
			want: image.Rect(100, 150, 200, 250),
		},
		{
			name: "Col, Center, Margin Top",
			a:    flexSize,
			b:    itemSize,
			c:    []int{100, 0, 0, 0},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 150, 200, 250),
		},
		{
			name: "Col, Center, Margin Top = Margin Bottom",
			a:    flexSize,
			b:    itemSize,
			c:    []int{100, 0, 100, 0},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 100, 200, 200),
		},
		{
			name: "Col, Center, Margin Bottom",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 100, 0},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 50, 200, 150),
		},
		{
			name: "Col, Start, Margin Top",
			a:    flexSize,
			b:    itemSize,
			c:    []int{50, 0, 0, 0},
			d:    furex.Column,
			e:    furex.JustifyStart,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 50, 200, 150),
		},
		{
			name: "Col, End, Margin Bottom",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 50, 0},
			d:    furex.Column,
			e:    furex.JustifyEnd,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 150, 200, 250),
		},
		{
			name: "Row, Center, Margin Right",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 100, 0, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(50, 100, 150, 200),
		},
		{
			name: "Row, Center, Margin Left",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 0, 100},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(150, 100, 250, 200),
		},
		{
			name: "Row, Center, Margin Right = Margin Left",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 100, 0, 100},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 100, 200, 200),
		},
		{
			name: "Row, Left, Margin Left",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 0, 100},
			d:    furex.Row,
			e:    furex.JustifyStart,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 100, 200, 200),
		},
		{
			name: "Row, Right, Margin Right",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 100, 0, 0},
			d:    furex.Row,
			e:    furex.JustifyEnd,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 100, 200, 200),
		},
		{
			name: "Row, Center, Margin Bottom",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 100, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 50, 200, 150),
		},
		{
			name: "Row, Start, Margin Top",
			a:    flexSize,
			b:    itemSize,
			c:    []int{50, 0, 0, 0},
			d:    furex.Row,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemStart,
			want: image.Rect(100, 50, 200, 150),
		},
		{
			name: "Col, Center, Margin Right",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 100, 0, 0},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(50, 100, 150, 200),
		},
		{
			name: "Col, Center, Margin Left",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 0, 100},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(150, 100, 250, 200),
		},
		{
			name: "Col, Center, Margin Right = Margin Left",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 100, 0, 100},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemCenter,
			want: image.Rect(100, 100, 200, 200),
		},
		{
			name: "Col, Left, Margin Left",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 0, 0, 100},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemStart,
			want: image.Rect(100, 100, 200, 200),
		},
		{
			name: "Col, Right, Margin Right",
			a:    flexSize,
			b:    itemSize,
			c:    []int{0, 100, 0, 0},
			d:    furex.Column,
			e:    furex.JustifyCenter,
			f:    furex.AlignItemEnd,
			want: image.Rect(100, 100, 200, 200),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flexMarginedBounds(
				tt.a, tt.b, tt.c, flex.Direction(tt.d), flex.Justify(tt.e), flex.AlignItem(tt.f),
			)
			if got != tt.want {
				t.Errorf("TestMarginedItem(%s): got %v; want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestMultiMarginedItems(t *testing.T) {
	flexSize := image.Pt(300, 300)
	itemSize := image.Pt(50, 50)

	var tests = []struct {
		name  string
		a     image.Point
		b     image.Point
		c     []int
		d     []int
		e     furex.Direction
		f     furex.Justify
		g     furex.AlignItem
		want1 image.Rectangle
		want2 image.Rectangle
	}{
		{
			name:  "Row, Center, A{}, B{Margin Left}",
			a:     flexSize,
			b:     itemSize,
			c:     []int{0, 0, 0, 0},
			d:     []int{0, 0, 0, 50},
			e:     furex.Row,
			f:     furex.JustifyCenter,
			g:     furex.AlignItemCenter,
			want1: image.Rect(75, 125, 125, 175),
			want2: image.Rect(175, 125, 225, 175),
		},
		{
			name:  "Col, Center, A{}, B{Margin Top}",
			a:     flexSize,
			b:     itemSize,
			c:     []int{0, 0, 0, 0},
			d:     []int{50, 0, 0, 0},
			e:     furex.Column,
			f:     furex.JustifyCenter,
			g:     furex.AlignItemCenter,
			want1: image.Rect(125, 75, 175, 125),
			want2: image.Rect(125, 175, 175, 225),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := flexMultiMarginedBounds(
				tt.a, tt.b, tt.c, tt.d, flex.Direction(tt.e), flex.Justify(tt.f), flex.AlignItem(tt.g),
			)
			if got1 != tt.want1 {
				t.Errorf("TestMarginedItem(%s) A: got %v; want %v", tt.name, got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("TestMarginedItem(%s) B: got %v; want %v", tt.name, got2, tt.want2)
			}
		})
	}
}

func TestMultiMarginedWrapRowItems(t *testing.T) {
	flexSize := image.Pt(200, 200)
	itemSize := image.Pt(85, 85)

	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Row
	flex.Justify = furex.JustifyStart
	flex.AlignItems = furex.AlignItemCenter
	flex.AlignContent = furex.AlignContentCenter
	flex.Wrap = furex.Wrap

	margin := []int{10, 0, 0, 10}

	item1 := NewMockMarginedItem(itemSize.X, itemSize.Y, margin)
	item2 := NewMockMarginedItem(itemSize.X, itemSize.Y, margin)
	item3 := NewMockMarginedItem(itemSize.X, itemSize.Y, margin)
	item4 := NewMockMarginedItem(itemSize.X, itemSize.Y, margin)
	flex.AddChild(item1)
	flex.AddChild(item2)
	flex.AddChild(item3)
	flex.AddChild(item4)

	flex.Update()
	flex.Draw(nil)

	want1 := image.Rect(10, 15, 10+85, 15+85)
	want2 := image.Rect(105, 15, 105+85, 15+85)
	want3 := image.Rect(10, 110, 10+85, 110+85)
	want4 := image.Rect(105, 110, 105+85, 110+85)

	got1 := item1.frame
	got2 := item2.frame
	got3 := item3.frame
	got4 := item4.frame

	if got1 != want1 {
		t.Errorf("TestMultiMarginedWrapRowItems box1: got %v; want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("TestMultiMarginedWrapRowItems box2: got %v; want %v", got2, want2)
	}
	if got3 != want3 {
		t.Errorf("TestMultiMarginedWrapRowItems box3: got %v; want %v", got3, want3)
	}
	if got4 != want4 {
		t.Errorf("TestMultiMarginedWrapRowItems box4: got %v; want %v", got4, want4)
	}
}

func flexMarginedBounds(flexSize image.Point, itemSize image.Point, margin []int, direction flex.Direction, justify flex.Justify, alignItem flex.AlignItem) image.Rectangle {
	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Direction(direction)
	flex.Justify = furex.Justify(justify)
	flex.AlignItems = furex.AlignItem(alignItem)

	item := NewMockMarginedItem(itemSize.X, itemSize.Y, margin)
	flex.AddChild(item)
	flex.Update()
	flex.Draw(nil)

	return item.frame
}

func flexMultiMarginedBounds(flexSize image.Point, itemSize image.Point, margin1 []int, margin2 []int, direction flex.Direction, justify flex.Justify, alignItem flex.AlignItem) (image.Rectangle, image.Rectangle) {
	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Direction(direction)
	flex.Justify = furex.Justify(justify)
	flex.AlignItems = furex.AlignItem(alignItem)

	item1 := NewMockMarginedItem(itemSize.X, itemSize.Y, margin1)
	item2 := NewMockMarginedItem(itemSize.X, itemSize.Y, margin2)
	flex.AddChild(item1)
	flex.AddChild(item2)
	flex.Update()
	flex.Draw(nil)

	return item1.frame, item2.frame
}

func flexItemBounds(flexSize image.Point, itemSize image.Point, direction flex.Direction, justify flex.Justify, alignItem flex.AlignItem) image.Rectangle {
	flex := furex.NewFlex(flexSize.X, flexSize.Y)
	flex.Direction = furex.Direction(direction)
	flex.Justify = furex.Justify(justify)
	flex.AlignItems = furex.AlignItem(alignItem)

	item := NewMockItem(itemSize.X, itemSize.Y)
	flex.AddChild(item)
	flex.Update()
	flex.Draw(nil)

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

func (m *MockItem) Size() (int, int) {
	return m.size.X, m.size.Y
}

func (m *MockItem) Draw(screen *ebiten.Image, frame image.Rectangle) {
	m.frame = frame
}

type MockMarginedItem struct {
	size   image.Point
	frame  image.Rectangle
	margin []int
}

func NewMockMarginedItem(w, h int, margin []int) *MockMarginedItem {
	m := new(MockMarginedItem)
	m.size = image.Pt(w, h)
	m.margin = margin
	return m
}

func (m *MockMarginedItem) Margin() []int {
	return m.margin
}

func (m *MockMarginedItem) Size() (int, int) {
	return m.size.X, m.size.Y
}

func (m *MockMarginedItem) Draw(screen *ebiten.Image, frame image.Rectangle) {
	m.frame = frame
}
