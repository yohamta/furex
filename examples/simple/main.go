package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/components"
)

type Game struct {
	init   bool
	screen screen
	gameUI *furex.View
}

type screen struct {
	Width  int
	Height int
}

func (g *Game) Update() error {
	if !g.init {
		g.init = true
		g.setupUI()
	}
	g.gameUI.Update()
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
	g.gameUI = &furex.View{
		Width:      g.screen.Width,
		Height:     g.screen.Height,
		Direction:  furex.Column,
		Justify:    furex.JustifySpaceBetween,
		AlignItems: furex.AlignItemCenter,
	}
	g.gameUI.AddChildren(
		&furex.View{
			Width:  100,
			Height: 100,
			Handler: &components.Box{
				Color: color.RGBA{0xff, 0, 0, 0xff},
			},
		},
		&furex.View{
			Width:  100,
			Height: 100,
			Handler: &components.Box{
				Color: color.RGBA{0, 0xff, 0, 0xff},
			},
		})
}

func main() {
	ebiten.SetWindowSize(480, 640)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
