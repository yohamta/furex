package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/furex"
)

type Game struct {
	vc *furex.ViewController
}

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

const desktopScreenScale = 2

var (
	screenWidth   int
	screenHeight  int
	isInitialized = false
)

func (g *Game) Update(screen *ebiten.Image) error {
	if isInitialized == false {
		g.buildUI()
		isInitialized = true
	}
	g.vc.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.vc.Draw(screen, image.Rect(0, 0, screenWidth, screenHeight))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) buildUI() {
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	for i := 0; i < 20; i++ {
		var c color.RGBA
		switch i % 3 {
		case 0:
			c = color.RGBA{0xff, 0, 0, 0xff}
		case 1:
			c = color.RGBA{0, 0xff, 0, 0xff}
		default:
			c = color.RGBA{0, 0, 0xff, 0xff}
		}
		rootFlex.AddChild(furex.NewBox(50, 50, c))
	}

	// view controller
	g.vc = furex.NewViewController()
	g.vc.SetRootView(rootFlex)
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(480, 640)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
