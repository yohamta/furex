<h1>Furex</h1>

Furex is a minimal GUI framework built on top of [Ebitengine](https://ebiten.org/), a 2D game engine for Go. It provides support for the [Flex Layout Algorithm](https://www.w3.org/TR/css-flexbox-1/#layout-algorithm), a powerful tool for laying out items of different sizes and dimensions. Furex is not a component library, but rather a framework for positioning and stacking virtual widgets, handling button and touch events for mobile, and more. How these widgets are rendered is up to the user.

<p align="center">
  <img width="480" src="./assets/example.gif">
</p>

Assets by [Kenney](https://kenney.nl). Fonts by [00ff](http://www17.plala.or.jp/xxxxxxx/00ff/).

Full source code of the example is [here](examples/game/main.go).

## Contents

- [Contents](#contents)
- [Motivation](#motivation)
- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Debug](#debug)
- [Contributions](#contributions)

## Motivation

[Flexbox](https://www.w3.org/TR/css-flexbox/) is a popular layout mechanism in web development, used for creating responsive and flexible user interfaces. With Furex, we can bring this same concept to game development using Go and Ebitengine. This library aims to make it easier for developers with experience in web or ReactNative projects to create dynamic, user-friendly game UI.

## Features

Here are some of the key features of Furex:

- Flexbox layout: The UI layout can be configured using the properties of [View](https://pkg.go.dev/github.com/yohamta/furex/v2#View) instances, which can be thought of as equivalent to `DIV` elements in HTML. These views can be stacked or nested to create complex layouts.

- Custom widgets: `View` instances can receive a `Handler` which is responsible for drawing and updating the view. This allows users to create any type of UI component by implementing the appropriate handler interfaces, such as [DrawHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#DrawHandler), [UpdateHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#UpdateHandler), and more.

- Button support: To create a button, users can implement the [ButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#ButtonHandler) interface. This supports both touch and mouse input for button actions. See the Example Button for more details.

- Touch and mouse events: Furex provides support for handling touch events and positions using the [TouchHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#TouchHandler) interface, and mouse click events using the [MouseLeftButtonHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseLeftButtonHandler) interface. It also offers support for detecting mouse position events using the [MouseHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseHandler) interface, and mouse enter/leave events using the [MouseEnterLeaveHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#MouseEnterLeaveHandler) interface.

- Swipe gestures: Users can detect swipe gestures by implementing the [SwipeHandler](https://pkg.go.dev/github.com/yohamta/furex/v2#SwipeHandler) interface.

These are just a few examples of the capabilities of Furex. For more information, be sure to check out the [GoDoc](https://pkg.go.dev/github.com/yohamta/furex/v2) documentation.

## Getting Started

To get started with Furex, install Furex using the following command:

```sh
go get github.com/yohamta/furex/v2
```

## Usage

Here's a simple example of how to use Furex to create an UI in your game:

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
  <img width="592" src="./assets/greens.png">
</p>

## Debug

You can turn on the Debug Mode by setting the variable below.
```go
furex.Debug = true
```

<p align="center">
  <img width="592" src="./assets/debug.png">
</p>

## Contributions

Contributions are welcome! If you find a bug or have an idea for a new feature, feel free to open an issue or submit a pull request.

