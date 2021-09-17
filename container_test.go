// Referenced code: https://github.com/golang/exp/blob/master/shiny/widget/flex/flex.go
package furex_test

import (
	"image"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
)

var (
	// testButtonFrame is the expected button frame for test layout
	testButtonFrame = image.Rect(290, 380, 300, 400)
)

func TestContainerButtonTouch(t *testing.T) {
	bf := testButtonFrame
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
			a:    bf.Min,
			b:    bf.Min,
			want: result{true, true, true},
		},
		// {
		// 	name: "press inside left-top edge, release outside",
		// 	a:    bf.Min,
		// 	b:    image.Pt(bf.Min.X, bf.Min.Y-1),
		// 	want: result{true, true, false},
		// },
		// {
		// 	name: "press inside righ-bottom edge, release inside",
		// 	a:    bf.Max,
		// 	b:    bf.Max,
		// 	want: result{true, true, true},
		// },
		// {
		// 	name: "press inside righ-bottom edge, release outside",
		// 	a:    bf.Max,
		// 	b:    image.Pt(bf.Max.X+1, bf.Max.Y),
		// 	want: result{true, true, false},
		// },
		// {
		// 	name: "press outside, release inside",
		// 	a:    image.Pt(bf.Min.X-1, bf.Min.Y),
		// 	b:    image.Pt(bf.Min.X+bf.Dx()/2, bf.Min.Y+bf.Dy()/2),
		// 	want: result{false, false, false},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := _TestContainerButtonTouch(tt.a, tt.b)
			got := result{b.isPressed, b.isReleased, b.isInside}
			if got != tt.want {
				t.Errorf("TestButtonTouch(%s): got %v; want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestContainerButtonMouse(t *testing.T) {
	bf := testButtonFrame
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
			a:    bf.Min,
			b:    bf.Min,
			want: result{true, true, true},
		},
		{
			name: "press inside left-top edge, release outside",
			a:    bf.Min,
			b:    image.Pt(bf.Min.X, bf.Min.Y-1),
			want: result{true, true, false},
		},
		{
			name: "press inside righ-bottom edge, release inside",
			a:    bf.Max,
			b:    bf.Max,
			want: result{true, true, true},
		},
		{
			name: "press inside righ-bottom edge, release outside",
			a:    bf.Max,
			b:    image.Pt(bf.Max.X+1, bf.Max.Y),
			want: result{true, true, false},
		},
		{
			name: "press outside, release inside",
			a:    image.Pt(bf.Min.X-1, bf.Min.Y),
			b:    image.Pt(bf.Min.X+bf.Dx()/2, bf.Min.Y+bf.Dy()/2),
			want: result{false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := _TestContainerButtonMouse(tt.a, tt.b)
			got := result{b.isPressed, b.isReleased, b.isInside}
			if got != tt.want {
				t.Errorf("TestButtonTouch(%s): got %v; want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestContainerMouseMove(t *testing.T) {
	bf := testButtonFrame
	type result struct {
		isMouseMoved bool
		mousePoint   image.Point
	}
	var tests = []struct {
		name string
		a    image.Point
		want result
	}{
		{
			name: "move mouse left-top inside",
			a:    image.Point{bf.Min.X, bf.Min.Y},
			want: result{isMouseMoved: true, mousePoint: image.Point{bf.Min.X, bf.Min.Y}},
		},
		{
			name: "move mouse right-bottom inside",
			a:    image.Point{bf.Max.X, bf.Max.Y},
			want: result{isMouseMoved: true, mousePoint: image.Point{bf.Max.X, bf.Max.Y}},
		},
		{
			name: "move mouse left outside",
			a:    image.Point{bf.Min.X - 1, bf.Min.Y},
			want: result{isMouseMoved: false, mousePoint: image.Point{-1, -1}},
		},
		{
			name: "move mouse right outside",
			a:    image.Point{bf.Max.X + 1, bf.Min.Y},
			want: result{isMouseMoved: false, mousePoint: image.Point{-1, -1}},
		},
		{
			name: "move mouse top outside",
			a:    image.Point{bf.Min.X, bf.Min.Y - 1},
			want: result{isMouseMoved: false, mousePoint: image.Point{-1, -1}},
		},
		{
			name: "move mouse bottom outside",
			a:    image.Point{bf.Min.X, bf.Max.Y + 1},
			want: result{isMouseMoved: false, mousePoint: image.Point{-1, -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := _TestContainerMouseMove(tt.a)
			got := result{b.isMouseMoved, b.mousePoint}
			if got != tt.want {
				t.Errorf("TestMouseMove(%s): got %v; want %v", tt.name, got, tt.want)
			}
		})
	}
}

func _TestContainerButtonTouch(pressedPos image.Point, releasedPos image.Point) *MockButton {
	flex, button := testLayout()
	flex.HandleJustPressedTouchID(0, pressedPos.X, pressedPos.Y)
	flex.HandleJustReleasedTouchID(0, releasedPos.X, releasedPos.Y)
	return button
}

func _TestContainerButtonMouse(pressedPos image.Point, releasedPos image.Point) *MockButton {
	flex, button := testLayout()
	flex.HandleJustPressedMouseButtonLeft(pressedPos.X, pressedPos.Y)
	flex.HandleJustReleasedMouseButtonLeft(releasedPos.X, releasedPos.Y)
	return button
}

func _TestContainerMouseMove(mousePoint image.Point) *MockButton {
	flex, button := testLayout()
	flex.HandleMouse(mousePoint.X, mousePoint.Y)
	return button
}

func testLayout() (*furex.Flex, *MockButton) {
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
	buttonSize := image.Pt(10, 20)
	button := NewMockButton(buttonSize.X, buttonSize.Y)
	inner1.AddChild(button)

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
	// │      │     (100,150)              │
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
	return flex, button
}

type MockButton struct {
	size         image.Point
	frame        image.Rectangle
	isPressed    bool
	isReleased   bool
	isInside     bool
	isMouseMoved bool
	mousePoint   image.Point
}

func NewMockButton(w, h int) *MockButton {
	m := new(MockButton)
	m.size = image.Pt(w, h)
	m.isInside = false
	m.isPressed = false
	m.isReleased = false
	m.isMouseMoved = false
	m.mousePoint = image.Pt(-1, -1)
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

func (m *MockButton) HandleMouse(x, y int) bool {
	m.isMouseMoved = true
	m.mousePoint = image.Pt(x, y)
	return true
}
