package furex

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsInside(t *testing.T) {
	for _, tt := range []struct {
		r    image.Rectangle
		x, y int
		want bool
	}{
		{
			r: image.Rect(0, 0, 10, 10),
			x: 0, y: 0,
			want: true,
		},
		{
			r: image.Rect(10, 10, 10, 10),
			x: 10, y: 10,
			want: true,
		},
		{
			r: image.Rect(10, 10, 20, 20),
			x: 20, y: 20,
			want: true,
		},
	} {
		require.Equal(t, tt.want, isInside(&tt.r, tt.x, tt.y))
	}
}
