# furex

A simple UI framework for [Ebiten](https://ebiten.org/) with a subset of flexbox layout specifications.
[GoDoc](https://pkg.go.dev/github.com/yohamta/furex)

## Motivation

When I was developing React Native apps, I thought the Flexbox layout was a very intuitive UI, so I thought it would be great if I could use the same concept for the UI of the games I make with Ebiten. I hope this library will help others with the same idea. Any types of Issue/PRs are welcomed :)

## Features

- Flexbox layout
  - Supports a subset of flexbox layout spec.
- Custom component
  - Supports any component that implements `DrawHandler` and `UpdateHandler`. See the [example](https://github.com/yohamta/furex/blob/master/components/box.go).
- Button comopnent
  - Able to handle touch ID on components that implements the `ButtonHandler` interface. See the [example](https://github.com/yohamta/furex/blob/master/components/button.go).
- Touch handling
  - Able to handle touch ID on components that implements the `TouchHandler` interface. 
- Mouse click / move handling
  - Able to handle left click on components that implements the `MouseHandler` interface. 
- Margin
- Absolute position
- Nesting view
- Wrap children
- Removing child view

## Layout Example
To check all examples, visit [this page](https://github.com/yohamta/furex/tree/main/examples).

[Full source code of the example](https://github.com/yohamta/furex/blob/master/examples/nesting/main.go)

<image src="https://user-images.githubusercontent.com/1475839/165524288-53827304-731e-4f33-81cd-26bb6a42e0d4.png" width="500px" />

## Simple Usage

```sh
go get github.com/yohamta/furex/v2
```

[Full source code of simple usage example](https://github.com/yohamta/furex/blob/master/examples/wrap/main.go)

```go
import "github.com/yohamta/furex/v2"

type Game struct {
  init   bool
  screen screen
  gameUI *furex.View
}

func (g *Game) Update() error {
  if !g.init {
    g.init = true
    g.setupUI()
  }
  g.gameUI.Update()
  return nil
}

func (g *Game) setupUI() {
  // create root view
  g.gameUI = &furex.View{
    Width:        g.screen.Width,
    Height:       g.screen.Height,
    Direction:    furex.Row,
    Justify:      furex.JustifyCenter,
    AlignItems:   furex.AlignItemCenter,
    AlignContent: furex.AlignContentCenter,
    Wrap:         furex.Wrap,
  }

  // create child view
  for i := 0; i < 20; i++ {
    g.gameUI.AddChild(&furex.View{
      Width:  100,
      Height: 100,
      Handler: &components.Box{
        Color: colors[i%len(colors)],
      },
    })
  }
}

func (g *Game) Draw(screen *ebiten.Image) {
  // Draw the UI tree
  g.gameUI.Draw(screen)
}

var colors = []color.Color{
  color.RGBA{0xaa, 0, 0, 0xff},
  color.RGBA{0, 0xaa, 0, 0xff},
  color.RGBA{0, 0, 0xaa, 0xff},
}
```

### Result
<image src="https://user-images.githubusercontent.com/1475839/133445715-b94b8c7f-bcd3-4aef-b7a4-b58bbb29d556.png" width="500px" />

