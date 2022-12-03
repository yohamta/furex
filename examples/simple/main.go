package main

import (
	"image"
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
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
	screen.Fill(color.RGBA{0x3d, 0x55, 0x0c, 0xff})
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
		color.RGBA{0x59, 0x98, 0x1a, 0xff},
		color.RGBA{0x81, 0xb6, 0x22, 0xff},
		color.RGBA{0xec, 0xf8, 0x7f, 0xff},
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
	FillRect(screen, &FillRectOpts{
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
