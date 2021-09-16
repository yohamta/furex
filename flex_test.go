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
		name      string
		flexSize  image.Point
		itemSize  image.Point
		direction furex.Direction
		justify   furex.Justify
		alignItem furex.AlignItem
		want      image.Rectangle
	}{
		{
			name:      "Column - Center, Center",
			flexSize:  flexSize,
			itemSize:  itemSize,
			direction: furex.Column,
			justify:   furex.JustifyCenter,
			alignItem: furex.AlignItemCenter,
			want:      image.Rect(25, 25, 75, 75),
		},
		{
			name:      "Column - Start, End",
			flexSize:  flexSize,
			itemSize:  itemSize,
			direction: furex.Column,
			justify:   furex.JustifyStart,
			alignItem: furex.AlignItemEnd,
			want:      image.Rect(50, 0, 100, 50),
		},
		{
			name:      "Row - Center, Center",
			flexSize:  flexSize,
			itemSize:  itemSize,
			direction: furex.Row,
			justify:   furex.JustifyCenter,
			alignItem: furex.AlignItemCenter,
			want:      image.Rect(25, 25, 75, 75),
		},
		{
			name:      "Row - End, Start",
			flexSize:  flexSize,
			itemSize:  itemSize,
			direction: furex.Row,
			justify:   furex.JustifyEnd,
			alignItem: furex.AlignItemStart,
			want:      image.Rect(50, 0, 100, 50),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flexItemBounds(
				tt.flexSize, tt.itemSize, flex.Direction(tt.direction), flex.Justify(tt.justify), flex.AlignItem(tt.alignItem),
			)
			if got != tt.want {
				t.Errorf("flexItemBounds(%s): got %v; want %v", tt.name, got, tt.want)
			}

		})
	}
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
