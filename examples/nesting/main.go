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
	newButton := func() *furex.View {
		return (&furex.View{
			Width:        100,
			Height:       100,
			MarginTop:    5,
			MarginBottom: 10,
			MarginLeft:   5,
			MarginRight:  5,
			Handler: &components.Button{
				Text:    "Button",
				OnClick: func() { println("button clicked") },
			},
		})
	}

	g.gameUI = (&furex.View{
		Width:      g.screen.Width,
		Height:     g.screen.Height,
		Direction:  furex.Column,
		Justify:    furex.JustifySpaceBetween,
		AlignItems: furex.AlignItemCenter,
	}).AddChild(
		(&furex.View{
			Width:      g.screen.Width - 20,
			Height:     70,
			Justify:    furex.JustifySpaceBetween,
			AlignItems: furex.AlignItemCenter,
		}).AddChild(
			&furex.View{
				Width:   100,
				Height:  100,
				Handler: &components.Box{Color: color.RGBA{0xff, 0, 0, 0xff}},
			},
			&furex.View{
				Width:   200,
				Height:  60,
				Handler: &components.Box{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}},
			},
			&furex.View{
				Width:   100,
				Height:  100,
				Handler: &components.Box{Color: color.RGBA{0, 0xff, 0, 0xff}},
			},
		),
	).AddChild(&furex.View{
		Width:  200,
		Height: 50,
		Handler: &components.Button{
			Text:    "Button",
			OnClick: func() { println("button clicked") },
		},
	}).AddChild((&furex.View{
		Width:      g.screen.Width,
		Height:     140,
		Justify:    furex.JustifyCenter,
		AlignItems: furex.AlignItemEnd,
	}).AddChild(
		newButton(),
		newButton(),
		newButton(),
		newButton(),
	))
}

func main() {
	ebiten.SetWindowSize(480, 640)

	furex.Debug = true

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
