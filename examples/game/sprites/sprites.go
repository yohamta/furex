package sprites

import (
	"encoding/xml"
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

var (
	sprites = map[string]*ganim8.Sprite{}
)

type textureAtlas struct {
	SubTextures []subTexture `xml:"SubTexture"`
}

type subTexture struct {
	Name   string `xml:"name,attr"`
	X      int    `xml:"x,attr"`
	Y      int    `xml:"y,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

type LoadOpts struct {
	PanelOpts map[string]PanelOpts
}

func LoadSprites(xmlPath string, imgPath string, opts LoadOpts) {
	dat, err := os.ReadFile(xmlPath)
	if err != nil {
		panic(err)
	}
	var atlas textureAtlas
	if err := xml.Unmarshal(dat, &atlas); err != nil {
		panic(err)
	}

	img, _, err := ebitenutil.NewImageFromFile(imgPath)
	if err != nil {
		panic(err)
	}

	rects := map[string]image.Rectangle{}
	for _, s := range atlas.SubTextures {
		g := ganim8.NewGrid(s.Width, s.Height, img.Bounds().Dx(), img.Bounds().Dy(), s.X, s.Y)
		rect := image.Rect(s.X, s.Y, s.X+s.Width, s.Y+s.Height)
		sprites[s.Name] = ganim8.NewSprite(img, g.Frames())
		rects[s.Name] = rect
	}

	for k, o := range opts.PanelOpts {
		r, ok := rects[k]
		if !ok {
			panic("panel not found: " + k)
		}
		panels := createPanels(img, r, o)
		for kk, v := range panels {
			sprites[fmt.Sprintf("%s_%s", k, kk)] = v
		}
	}
}

func Get(name string) *ganim8.Sprite {
	if s, ok := sprites[name]; ok {
		return s
	}
	panic("sprite not found: " + name)
}
