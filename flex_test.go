package furex

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlexAlignments(t *testing.T) {
	w, h := 100, 100
	child := &View{
		Width:  50,
		Height: 50,
	}

	var tests = []struct {
		name string
		flex *View
		want image.Rectangle
	}{
		{
			name: "Column - Center, Center",
			flex: &View{
				Width:      w,
				Height:     h,
				Direction:  Column,
				Justify:    JustifyCenter,
				AlignItems: AlignItemCenter,
			},
			want: image.Rect(25, 25, 75, 75),
		},
		{
			name: "Column - Start, End",
			flex: &View{
				Width:      w,
				Height:     h,
				Direction:  Column,
				Justify:    JustifyStart,
				AlignItems: AlignItemEnd,
			},
			want: image.Rect(50, 0, 100, 50),
		},
		{
			name: "Row - Center, Center",
			flex: &View{
				Width:      w,
				Height:     h,
				Direction:  Row,
				Justify:    JustifyCenter,
				AlignItems: AlignItemCenter,
			},
			want: image.Rect(25, 25, 75, 75),
		},
		{
			name: "Row - End, Start",
			flex: &View{
				Width:      w,
				Height:     h,
				Direction:  Row,
				Justify:    JustifyEnd,
				AlignItems: AlignItemStart,
			},
			want: image.Rect(50, 0, 100, 50),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flexItemBounds(tt.flex, child)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFlexWrap(t *testing.T) {
	flex := &View{
		Width:      200,
		Height:     200,
		Direction:  Row,
		Justify:    JustifyStart,
		AlignItems: AlignItemStart,
		Wrap:       Wrap,
	}

	mocks := [3]mockHandler{}
	flex.AddChild(&View{Width: 100, Height: 100, Handler: &mocks[0]})
	flex.AddChild(&View{Width: 100, Height: 100, Handler: &mocks[1]})
	flex.AddChild(&View{Width: 100, Height: 100, Handler: &mocks[2]})

	flex.Update()
	flex.Draw(nil)

	// (0,0)
	// ┌───────────────(100,0)───────────┐
	// │box1            │box2            │
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
	// └──────────────(100,200)──────────┘
	// 															 (200,200)

	assert.Equal(t, image.Rect(0, 100, 100, 200), mocks[2].Frame)
}

func TestAbsolutePos(t *testing.T) {
	left, top := 20, 30
	f1 := &View{
		Width:      100,
		Height:     200,
		Left:       left,
		Top:        top,
		Position:   PositionAbsolute,
		Direction:  Row,
		Justify:    JustifyCenter,
		AlignItems: AlignItemCenter,
		Wrap:       Wrap,
	}

	mock := mockHandler{}

	f1.AddChild(&View{Width: 30, Height: 40, Handler: &mock})
	f1.Update()
	f1.Draw(nil)

	//  (0,0)
	//  ┌───────────────────────────────────┐
	//  │                                   │
	//  │                                   │
	//  │         (100,50)                  │
	//  │           ┌───────────────────────┤
	//  │           │ flex                  │
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
	//  └───────────┴───────────────────────┘
	//                                  (150,200)

	w, h := 30, 40
	x, y := 100/2-w/2+left, 200/2-h/2+top
	require.Equal(t, image.Rect(x, y, x+w, y+h), mock.Frame)
}

func TestAbsolutePosNested(t *testing.T) {
	f1 := &View{
		Width:      150,
		Height:     200,
		Direction:  Row,
		Justify:    JustifyStart,
		AlignItems: AlignItemCenter,
		Wrap:       Wrap,
	}

	f2 := &View{
		Width:      50,
		Height:     150,
		Left:       100,
		Top:        50,
		Position:   PositionAbsolute,
		Direction:  Row,
		Justify:    JustifyCenter,
		AlignItems: AlignItemCenter,
		Wrap:       Wrap,
	}

	f1.AddChild(f2)

	mock := mockHandler{}

	f2.AddChild(&View{Width: 30, Height: 40, Handler: &mock})
	f1.Update()
	f1.Draw(nil)

	//  (0,0)
	//  ┌───────────────────────────────────┐
	//  │                                   │
	//  │                                   │
	//  │         (100,50)                  │
	//  │           ┌───────────────────────┤
	//  │           │ flex                  │
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
	//  └───────────┴───────────────────────┘
	//                                  (150,200)

	require.Equal(t, image.Rect(100, 50, 150, 200), f2.frame)

	w, h := 30, 40
	x, y := 100+50/2-w/2, 50+150/2-h/2
	require.Equal(t, image.Rect(x, y, x+w, y+h), mock.Frame)
}

func TestNesting(t *testing.T) {
	parent := &View{
		Width:      300,
		Height:     500,
		Direction:  Column,
		Justify:    JustifyCenter,
		AlignItems: AlignItemCenter,
		Left:       100,
		Top:        50,
		Position:   PositionAbsolute,
	}

	child := &View{
		Width:      100,
		Height:     200,
		Direction:  Column,
		Justify:    JustifyEnd,
		AlignItems: AlignItemEnd,
	}

	parent.AddChild(child)

	item := &mockHandler{}

	child.AddChild(&View{
		Width:   30,
		Height:  40,
		Handler: item,
	})

	parent.Update()
	parent.Draw(nil)

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
	// expected item Frame:
	// x = 300-30 = 270 to 300
	// y = 400-40 = 360 to 400

	want := image.Rect(270, 360, 300, 400)
	require.Equal(t, want, item.Frame)
}

func TestMargin(t *testing.T) {
	var tests = []struct {
		Flex *View
		View *View
		Want image.Rectangle
	}{
		{
			Flex: &View{
				Width:      100,
				Height:     100,
				Direction:  Row,
				Justify:    JustifyCenter,
				AlignItems: AlignItemCenter,
			},
			View: &View{
				Width:      50,
				Height:     50,
				MarginLeft: 20,
			},
			Want: image.Rect(25+10, 25, 75+10, 75),
		},
		{
			Flex: &View{
				Width:      100,
				Height:     100,
				Direction:  Column,
				Justify:    JustifyCenter,
				AlignItems: AlignItemCenter,
			},
			View: &View{
				Width:     50,
				Height:    50,
				MarginTop: 20,
			},
			Want: image.Rect(25, 25+10, 75, 75+10),
		},
		{
			Flex: &View{
				Width:      100,
				Height:     100,
				Direction:  Row,
				Justify:    JustifyEnd,
				AlignItems: AlignItemStart,
			},
			View: &View{
				Width:       50,
				Height:      50,
				MarginTop:   10,
				MarginRight: 10,
			},
			Want: image.Rect(40, 10, 90, 60),
		},
		{
			Flex: &View{
				Width:      100,
				Height:     100,
				Direction:  Column,
				Justify:    JustifyEnd,
				AlignItems: AlignItemEnd,
			},
			View: &View{
				Width:        50,
				Height:       50,
				MarginRight:  10,
				MarginBottom: 10,
			},
			Want: image.Rect(40, 40, 90, 90),
		},
	}

	for _, tt := range tests {
		mock := &mockHandler{}
		tt.View.Handler = mock
		tt.Flex.AddChild(tt.View)
		tt.Flex.Update()
		tt.Flex.Draw(nil)

		assert.Equal(t, tt.Want, mock.Frame)
	}
}

func TestMarginedItemPosition(t *testing.T) {
	flex := &View{
		Width:      200,
		Height:     200,
		Direction:  Column,
		Justify:    JustifyStart,
		AlignItems: AlignItemCenter,
	}

	mocks := [2]mockHandler{}

	view1 := &View{
		Width:      200,
		Height:     50,
		MarginTop:  10,
		Direction:  Column,
		Justify:    JustifyStart,
		AlignItems: AlignItemCenter,
		Handler:    &mocks[0],
	}
	flex.addChild(view1)

	view2 := &View{
		Width:   200,
		Height:  10,
		Handler: &mocks[1],
	}
	view1.addChild(view2)

	flex.Update()
	flex.Draw(nil)

	require.Equal(t, image.Rect(0, 10, 200, 60), mocks[0].Frame)
	require.Equal(t, image.Rect(0, 10, 200, 20), mocks[1].Frame)
}

func TestMultiMarginedWrapRowItems(t *testing.T) {
	flex := &View{
		Width:        200,
		Height:       200,
		Direction:    Row,
		Justify:      JustifyStart,
		AlignItems:   AlignItemCenter,
		AlignContent: AlignContentCenter,
		Wrap:         Wrap,
	}

	mocks := [4]mockHandler{}
	view := View{
		Width:      85,
		Height:     85,
		MarginTop:  10,
		MarginLeft: 10,
	}

	for i := 0; i < 4; i++ {
		v := view
		v.Handler = &mocks[i]
		flex.AddChild(&v)
	}

	flex.Update()
	flex.Draw(nil)

	assert.Equal(t, image.Rect(10, 15, 10+85, 15+85), mocks[0].Frame)
	assert.Equal(t, image.Rect(105, 15, 105+85, 15+85), mocks[1].Frame)
	assert.Equal(t, image.Rect(10, 110, 10+85, 110+85), mocks[2].Frame)
	assert.Equal(t, image.Rect(105, 110, 105+85, 110+85), mocks[3].Frame)
}

func TestRemoveChild(t *testing.T) {
	w, h := 100, 100

	flex := &View{
		Width:      w,
		Height:     h,
		Direction:  Row,
		Justify:    JustifyCenter,
		AlignItems: AlignItemCenter,
	}

	mocks := [2]mockHandler{}
	views := [2]*View{}

	for i := 0; i < 2; i++ {
		views[i] = &View{
			Width:   50,
			Height:  50,
			Handler: &mocks[i],
		}
		flex.AddChild(views[i])
	}

	flex.Update()
	flex.Draw(nil)

	require.Equal(t, mocks[0].Frame, image.Rect(0, 25, 50, 75))
	require.Equal(t, mocks[1].Frame, image.Rect(50, 25, 100, 75))

	flex.RemoveChild(views[0])
	flex.Update()
	flex.Draw(nil)

	require.Equal(t, mocks[1].Frame, image.Rect(25, 25, 75, 75))
}

func TestAutoExpanding(t *testing.T) {
	flex := &View{
		Width:      1000,
		Height:     1000,
		Direction:  Row,
		Justify:    JustifyCenter,
		AlignItems: AlignItemStretch,
	}

	mocks := [2]mockHandler{}
	for i := 0; i < 2; i++ {
		v := &View{
			Grow:    1,
			Handler: &mocks[i],
		}
		flex.AddChild(v)
	}

	flex.Update()
	flex.Draw(nil)

	assert.Equal(t, image.Rect(0, 0, 500, 1000), mocks[0].Frame)
	assert.Equal(t, image.Rect(500, 0, 1000, 1000), mocks[1].Frame)
}

func TestNestedChildrenGrow(t *testing.T) {
	flex := &View{
		Width:      1000,
		Height:     1000,
		Direction:  Row,
		Justify:    JustifyCenter,
		AlignItems: AlignItemStretch,
	}

	child := &View{
		Direction:  Column,
		Justify:    JustifyCenter,
		AlignItems: AlignItemStretch,
		Grow:       1,
	}

	flex.AddChild(child)

	child2 := &View{
		Direction:  Row,
		Justify:    JustifyCenter,
		AlignItems: AlignItemStretch,
		Grow:       1,
	}

	child.AddChild(child2)

	mocks := [2]mockHandler{}
	for i := 0; i < 2; i++ {
		v := &View{
			Grow:    1,
			Handler: &mocks[i],
		}
		child2.AddChild(v)
	}

	flex.Update()
	flex.Draw(nil)

	assert.Equal(t, image.Rect(0, 0, 500, 1000), mocks[0].Frame)
	assert.Equal(t, image.Rect(500, 0, 1000, 1000), mocks[1].Frame)
}

func TestNestedChildGrow(t *testing.T) {
	flex := &View{
		Width:      1000,
		Height:     1000,
		Direction:  Column,
		AlignItems: AlignItemStretch,
		Justify:    JustifyCenter,
	}

	mock := mockHandler{}

	flex.AddChild(
		(&View{
			Direction:  Column,
			AlignItems: AlignItemStretch,
			Justify:    JustifyCenter,
			Grow:       1,
		}).AddChild(
			&View{
				Direction:  Row,
				AlignItems: AlignItemStretch,
				Justify:    JustifyCenter,
				Grow:       1,
				Handler:    &mock,
			},
		),
	)

	flex.Update()
	flex.Draw(nil)

	assert.Equal(t, image.Rect(0, 0, 1000, 1000), mock.Frame)
}

func flexItemBounds(parent *View, child *View) image.Rectangle {
	mock := &mockHandler{}
	child.Handler = mock

	parent.AddChild(child)
	parent.Update()
	parent.Draw(nil)

	return mock.Frame
}
