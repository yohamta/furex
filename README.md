# Furex

Furex is a minimal GUI framework for [Ebitengine](https://ebiten.org/) that supports a subset of [Flex Layout Algorithm](https://www.w3.org/TR/css-flexbox-1/#layout-algorithm).

For now, Furex is not a component library but a framework for positioning and stacking virtual widgets, handling button and touch events. How they are rendered is completely up to the user.

[GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2)

## Motivation

[Flexbox](https://www.w3.org/TR/css-flexbox/) is a good mechanism for laying out items of different sizes. I wanted to use the same concept for game UI because I have experience in Web and ReactNative projects. I hope the library helps other people with the same thoughts.

## Features

| Feature                  | How to use                                                                                                                                                                                                                                                                                                                        |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Flexbox layout           | The UI layout can be configured via the properties of [View](https://pkg.go.dev/github.com/yohamta/furex/v2#View) instances. We can think of a `View` as a `DIV` element in HTML, which can be stacked or nested.                                                                                                                 |
| Custom widgets           | A `View` can receive a `Handler` which draws and updates the `View`. We can create any type of UI components by implementing the handler interfaces (e.g., [DrawHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#DrawHandler), [UpdateHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#UpdateHandler), etc) |
| Buttons                  | To create a `Button`, you can implement the [ButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#ButtonHandler). It supports both touch and mouse. See the [Example Button](./examples/game/widgets/button.go) for more details.          |
| Touch events             | To handle `Touch` events and touch positions, you can implement [TouchHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#TouchHandler).                                                                                                                                                                                      |
| Mouse click events       | To handle `MouseClick` events, you can implement [MouseLeftButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseLeftButtonHandler).                                                                                                                                                                                 |
| Mouse move events        | To detect `Mouse` position events, you can implement [MouseHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseHandler).                                                                                                                                                                                                 |
| Mouse enter/leave events | To detect `MouseEnter`/`MouseLeave` events, you can implement [MouseEnterLeaveHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseEnterLeaveHandler).                                                                                                                                                                    |
| Swipe gestures           | To detect `Swipe` gestures you can implement [SwipeHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#SwipeHandler).                                                                                                                                                                                                         |
| Margins                  | A `View`'s margin can be configured by setting the field values, `MarginLeft`, `MarginTop`, `MarginRight`, and `MarginBottom`                                                                                                                                                                                                     |
| Absolute positions       | It is possible to place a `View` to an absolute position by setting `PositionAbsolute` to the `Position` property. Also you can set `Left`, and `Top` property values to set the `View` to specify a certain position.                                                                                                            |

## Install

```sh
go get github.com/yohamta/furex/v2
```

## Example
[Full source code of the example](examples/game/main.go)

Assets by [Kenney](https://kenney.nl).

<p align="center">
  <img width="480" height="640" src="./assets/example.gif">
</p>

## Usage

[Full source code of the example](examples/simple/main.go)

```go
type Game struct {
	initOnce sync.Once
	screen   screen
	gameUI   *furex.View
}

func (g *Game) Update() error {
	g.initOnce.Do(func() {
		g.setupUI()
	})
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.gameUI.Draw(screen)
}

func (g *Game) setupUI() {
	screen.Fill(color.RGBA{0x3d, 0x55, 0x0c, 0xff})
	colors := []color.Color{
		color.RGBA{0x3d, 0x55, 0x0c, 0xff},
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
	graphic.FillRect(screen, &graphic.FillRectOpts{
		Rect: frame, Color: b.Color,
	})
}
```

<p align="center">
  <img width="592" height="780" src="./assets/greens.png">
</p>

## Debug

You can turn on the Debug Mode by setting the variable below.
```go
furex.Debug = true
```

<p align="center">
  <img width="592" height="780" src="./assets/debug.png">
</p>

## How to contribute?

Feel free to contribute in any way you want. Share ideas, submit issues, create pull requests, adding examples, or adding widgets. Thank you!
