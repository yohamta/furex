package sprites

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

type PanelOpts struct {
	Border int
	Center int
}

func createPanels(img *ebiten.Image, r image.Rectangle, opts PanelOpts) map[string]*ganim8.Sprite {
	ret := map[string]*ganim8.Sprite{}
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	border := opts.Border
	center := opts.Center
	cx, cy := r.Min.X+r.Dx()/2, r.Min.Y+r.Dy()/2

	// top left
	g := ganim8.NewGrid(border, border, w, h, r.Min.X, r.Min.Y)
	ret["top_left"] = ganim8.NewSprite(img, g.Frames())
	// top
	g = ganim8.NewGrid(center, border, w, h, cx-center/2, r.Min.Y)
	ret["top"] = ganim8.NewSprite(img, g.Frames())
	// top right
	g = ganim8.NewGrid(border, border, w, h, r.Min.X+r.Dx()-border, r.Min.Y)
	ret["top_right"] = ganim8.NewSprite(img, g.Frames())
	// left
	g = ganim8.NewGrid(border, center, w, h, r.Min.X, cy-center/2)
	ret["left"] = ganim8.NewSprite(img, g.Frames())
	// center
	g = ganim8.NewGrid(center, center, w, h, cx-center/2, cy-center/2)
	ret["center"] = ganim8.NewSprite(img, g.Frames())
	// right
	g = ganim8.NewGrid(border, center, w, h, r.Min.X+r.Dx()-border, cy-center/2)
	ret["right"] = ganim8.NewSprite(img, g.Frames())
	// bottom left
	g = ganim8.NewGrid(border, border, w, h, r.Min.X, r.Min.Y+r.Dy()-border)
	ret["bottom_left"] = ganim8.NewSprite(img, g.Frames())
	// bottom
	g = ganim8.NewGrid(center, border, w, h, cx-center/2, r.Max.Y-border)
	ret["bottom"] = ganim8.NewSprite(img, g.Frames())
	// bottom right
	g = ganim8.NewGrid(border, border, w, h, r.Min.X+r.Dx()-border, r.Min.Y+r.Dy()-border)
	ret["bottom_right"] = ganim8.NewSprite(img, g.Frames())

	return ret
}
