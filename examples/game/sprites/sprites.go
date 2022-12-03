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

type PanelOpts struct {
	Name   string
	Border int
}

type LoadOpts struct {
	PanelOpts []PanelOpts
}

func LoadSprites(opts LoadOpts) {
	dat, err := os.ReadFile("assets/uipack_rpg_sheet.xml")
	if err != nil {
		panic(err)
	}
	var atlas textureAtlas
	if err := xml.Unmarshal(dat, &atlas); err != nil {
		panic(err)
	}

	img, _, err := ebitenutil.NewImageFromFile("assets/uipack_rpg_sheet.png")
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

	for _, o := range opts.PanelOpts {
		r, ok := rects[o.Name]
		if !ok {
			panic("panel not found: " + o.Name)
		}
		panels := createPanels(img, r, o.Border)
		for k, v := range panels {
			sprites[fmt.Sprintf("%s_%s", o.Name, k)] = v
		}
	}
}

func Get(name string) *ganim8.Sprite {
	if s, ok := sprites[name]; ok {
		return s
	}
	panic("sprite not found: " + name)
}
