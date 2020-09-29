package furex

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Rect struct {
	X, Y, W, H int
}

var imgRect *ebiten.Image

func createRectImg() *ebiten.Image {
	if imgRect != nil {
		return imgRect
	}
	imgRect, err := ebiten.NewImage(1, 1, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	return imgRect
}

func FillRect(target *ebiten.Image, r Rect, clr color.Color) {
	img := createRectImg()
	img.Fill(clr)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X)*(1/float64(r.W)), float64(r.Y)*(1/float64(r.H)))
	op.GeoM.Scale(float64(r.W), float64(r.H))

	target.DrawImage(img, op)
}

func DrawRect(target *ebiten.Image, r Rect, clr color.Color, width int) {
	FillRect(target, Rect{r.X, r.Y, width, r.H}, clr)
	FillRect(target, Rect{r.X + r.W - width, r.Y, width, r.H}, clr)
	FillRect(target, Rect{r.X, r.Y, r.W, width}, clr)
	FillRect(target, Rect{r.X, r.Y + r.H - width, r.W, width}, clr)
}
