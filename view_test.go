package furex

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddChildUpdateRemove(t *testing.T) {
	view := &View{
		Width:      100,
		Height:     100,
		Direction:  Row,
		Justify:    JustifyStart,
		AlignItems: AlignItemStart,
	}

	mock := &mockHandler{}
	child := &View{
		Width:   10,
		Height:  10,
		Handler: mock,
	}
	view.AddChild(child)
	require.True(t, view.isDirty)

	view.Update()
	require.True(t, mock.IsUpdated)

	view.Draw(nil)
	require.Equal(t, image.Rect(0, 0, 10, 10), mock.Frame)

	require.True(t, view.RemoveChild(child))
	require.Equal(t, 0, len(view.children))
}

func TestAddChildren(t *testing.T) {
	view := &View{
		Width:      100,
		Height:     100,
		Direction:  Row,
		Justify:    JustifyStart,
		AlignItems: AlignItemStart,
	}

	mocks := [2]mockHandler{}
	view.AddChildren(
		&View{
			Width:   10,
			Height:  10,
			Handler: &mocks[0],
		},
		&View{
			Width:   10,
			Height:  10,
			Handler: &mocks[1],
		},
	)

	view.Update()
	require.True(t, mocks[0].IsUpdated)
	require.True(t, mocks[1].IsUpdated)

	view.Draw(nil)
	require.Equal(t, image.Rect(0, 0, 10, 10), mocks[0].Frame)
	require.Equal(t, image.Rect(10, 0, 20, 10), mocks[1].Frame)

	view.RemoveAll()
	require.Equal(t, 0, len(view.children))
}
