package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/furex"
)

type Game struct {
	viewController *furex.ViewController
	rootView       *furex.View
}

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

var (
	screenWidth     = 240
	screenHeight    = 360
	isWindowSizeSet = false
	isInitialized   = false
)

func (g *Game) Update(screen *ebiten.Image) error {
	if isInitialized == false {
		if isWindowSizeSet {
			g.buildUI()
		}
		return nil
	}
	g.viewController.Update()
	return nil
}

func (g *Game) SetWindowSize(width, height int) {
	screenHeight = int(float64(screenWidth) / float64(width) * float64(height))
	isWindowSizeSet = true
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.viewController.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) buildUI() {
	g.viewController = furex.NewViewController()
	g.viewController.SetFrame(0, 0, screenWidth, screenHeight)

	// root flex container
	rootFlex := furex.NewFlex(screenWidth, screenHeight)
	rootFlex.Direction = furex.Column
	rootFlex.Justify = furex.JustifySpaceBetween
	flexContainerView := furex.NewView(rootFlex)
	g.rootView = flexContainerView
	g.viewController.SetRootView(flexContainerView)

	b0 := furex.NewView(furex.NewBox(100, 100, color.RGBA{0xff, 0, 0, 0xff}))
	g.rootView.AddSubView(b0)

	b1 := furex.NewView(furex.NewBox(100, 100, color.RGBA{0, 0xff, 0, 0xff}))
	g.rootView.AddSubView(b1)
}

func main() {
	windowSize := image.Point{480, 800}
	ebiten.SetWindowSize(windowSize.X, windowSize.Y)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	game.SetWindowSize(windowSize.X, windowSize.Y)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
