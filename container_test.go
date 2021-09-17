// Referenced code: https://github.com/golang/exp/blob/master/shiny/widget/flex/flex.go
package furex_test

import (
	"image"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
)

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
