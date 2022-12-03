package main

import (
	"image"
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/common/graphic"
)

type Game struct {
	initOnce sync.Once
	screen   screen
	gameUI   *furex.View
}

type screen struct {
	Width  int
	Height int
}

func (g *Game) Update() error {
	g.initOnce.Do(func() {
		g.setupUI()
	})
	g.gameUI.UpdateWithSize(ebiten.WindowSize())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.screen.Width = outsideWidth
	g.screen.Height = outsideHeight
	return g.screen.Width, g.screen.Height
}

func NewGame() (*Game, error) {
	game := &Game{}
	return game, nil
}

func (g *Game) setupUI() {
	colors := []color.Color{
		color.RGBA{0xaa, 0, 0, 0xff},
		color.RGBA{0, 0xaa, 0, 0xff},
		color.RGBA{0, 0, 0xaa, 0xff},
	}

	g.gameUI = &furex.View{
		Width:        g.screen.Width,
		Height:       g.screen.Height,
		Direction:    furex.Row,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
		AlignContent: furex.AlignContentCenter,
		Wrap:         furex.Wrap,
	}

	for i := 0; i < 20; i++ {
		g.gameUI.AddChild(&furex.View{
			Width:  100,
			Height: 100,
			Handler: &Box{
				Color: colors[i%len(colors)],
			},
		})
	}
}

type Box struct {
	Color color.Color
}

var _ furex.DrawHandler = (*Box)(nil)

func (b *Box) HandleDraw(screen *ebiten.Image, frame image.Rectangle) {
	// TODO: replace with ebiten/vector utility functions
	graphic.FillRect(screen, &graphic.FillRectOpts{
		Rect: frame, Color: b.Color,
	})
}

func main() {
	ebiten.SetWindowSize(480, 640)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
