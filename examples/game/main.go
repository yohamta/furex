package main

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/examples/game/sprites"
	"github.com/yohamta/furex/v2/examples/game/text"
	"github.com/yohamta/furex/v2/examples/game/widgets"

	_ "embed"
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
	g.gameUI.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{63, 124, 182, 255})
	g.gameUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.screen.Width = outsideWidth
	g.screen.Height = outsideHeight
	return g.screen.Width, g.screen.Height
}

func NewGame() (*Game, error) {
	text.LoadFonts()
	sprites.LoadSprites(
		"assets/images/uipack_rpg_sheet.xml",
		"assets/images/uipack_rpg_sheet.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{
				"panelInset_beige.png": {
					Border: 32,
					Center: 36,
				},
				"panel_brown.png": {
					Border: 32,
					Center: 36,
				},
			},
		})
	sprites.LoadSprites(
		"assets/images/uipackSpace_sheet.xml",
		"assets/images/uipackSpace_sheet.png",
		sprites.LoadOpts{
			PanelOpts: map[string]sprites.PanelOpts{
				"glassPanel_corners.png": {
					Border: 40,
					Center: 20,
				},
				"glassPanel_projection.png": {
					Border: 20,
					Center: 10,
				},
			},
		})
	game := &Game{}
	return game, nil
}

//go:embed assets/html/main.html
var mainHTML string

func (g *Game) setupUI() {
	g.gameUI = furex.Parse(mainHTML, &furex.ParseOptions{
		Width:  g.screen.Width,
		Height: g.screen.Height,
		Components: map[string]furex.Component{
			"panel":        &widgets.Panel{},
			"gauge-text":   &widgets.Text{Color: color.RGBA{50, 48, 41, 255}},
			"health-gauge": &widgets.Bar{Color: "Green", Value: .8},
			"mana-gauge":   &widgets.Bar{Color: "Blue", Value: .8},
			"button-big": func() furex.Handler {
				return &widgets.Button{
					Sprite:        "buttonLong_blue.png",
					SpritePressed: "buttonLong_blue_pressed.png",
					OnClick:       func() { println("button clicked") },
				}
			},
			"button-small": func() furex.Handler {
				return &widgets.Button{
					Sprite:        "buttonSquare_blue.png",
					SpritePressed: "buttonSquare_blue_pressed.png",
					OnClick:       func() { println("button clicked") },
				}
			},
			"close-button": &widgets.Button{
				Sprite:  "buttonRound_blue.png",
				OnClick: func() { println("button clicked") },
			},
			"close-button-sprite": &widgets.Sprite{
				Sprite: "iconCross_beige.png",
			},
			"bottom-button": func() furex.Handler {
				return &widgets.Button{
					Color:         color.RGBA{210, 178, 144, 255},
					Sprite:        "buttonSquare_brown.png",
					SpritePressed: "buttonSquare_brown_pressed.png",
					OnClick:       func() { println("button clicked") },
				}
			},
			"glass-button": func() furex.Handler {
				return &widgets.Panel{OnClick: func() { println("button clicked") }}
			},
			"play-game-text": &widgets.Text{
				Color:     color.RGBA{45, 73, 94, 255},
				HorzAlign: etxt.XCenter,
				VertAlign: etxt.YCenter,
			},
		},
	})

	// panels that draws mouse cursor
	g.gameUI.AddChild(
		&furex.View{
			Width:    g.screen.Width,
			Height:   g.screen.Height,
			Position: furex.PositionAbsolute,
			Left:     0,
			Top:      0,
			Handler:  &widgets.Mouse{},
		},
	)
}

func main() {
	ebiten.SetWindowSize(480, 640)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
