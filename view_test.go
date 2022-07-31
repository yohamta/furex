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
	require.Equal(t, view, view.AddChild(child))
	require.True(t, view.isDirty)

	view.Update()
	require.True(t, mock.IsUpdated)

	view.Draw(nil)
	require.Equal(t, image.Rect(0, 0, 10, 10), mock.Frame)

	require.True(t, view.RemoveChild(child))
	require.Equal(t, 0, len(view.children))
}

func TestUpdateWithSize(t *testing.T) {
	view := &View{
		Width:      100,
		Height:     100,
		Direction:  Row,
		Justify:    JustifyCenter,
		AlignItems: AlignItemCenter,
	}

	mock := &mockHandler{}
	child := &View{
		Width:   10,
		Height:  10,
		Handler: mock,
	}
	require.Equal(t, view, view.AddChild(child))

	view.UpdateWithSize(200, 200)
	require.True(t, mock.IsUpdated)

	view.Draw(nil)
	require.Equal(t, image.Rect(95, 95, 105, 105), mock.Frame)

}

func TestAddToParent(t *testing.T) {
	root := &View{
		Width:      100,
		Height:     100,
		Direction:  Row,
		Justify:    JustifyStart,
		AlignItems: AlignItemStart,
	}

	mock := &mockHandler{}

	child := (&View{
		Width:   10,
		Height:  10,
		Handler: mock,
	})

	require.Equal(t, child, child.AddTo(root))

	root.Update()
	require.True(t, mock.IsUpdated)

}

func TestAddChild(t *testing.T) {
	view := &View{
		Width:      100,
		Height:     100,
		Direction:  Row,
		Justify:    JustifyStart,
		AlignItems: AlignItemStart,
	}

	mocks := [2]mockHandler{}
	require.Equal(t, view, view.AddChild(
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
	))

	view.Update()
	require.True(t, mocks[0].IsUpdated)
	require.True(t, mocks[1].IsUpdated)

	view.Draw(nil)
	require.Equal(t, image.Rect(0, 0, 10, 10), mocks[0].Frame)
	require.Equal(t, image.Rect(10, 0, 20, 10), mocks[1].Frame)

	view.RemoveAll()
	require.Equal(t, 0, len(view.children))
}
