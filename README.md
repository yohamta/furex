# furex
A simple UI framework for [Ebiten](https://ebiten.org/) with a subset of flexbox layout specification.
[GoDoc](https://pkg.go.dev/github.com/yohamta/furex)

## Motivation
When I was making React Native apps, Flexbox layout felt useful and intuitive to me, so I thought it would be great if I can use the same concept to build user interface in my game using Ebiten framework. I would be so happy if it helps someone who has the same thoughts. I really appreciate if anyone can make any contribution to grow the library.

## Features
| Feature                 | Supported | Note                                                                                                                   |
|-------------------------|------------------|------------------------------------------------------------------------------------------------------------------------|
| Flexbox layout          | o                | Supports a subset of flexbox layout spec.                                                                              |
| Custom component   | o                | Supports any component that implements `Component` interface. See the [example](https://github.com/yohamta/furex/blob/master/examples/shared/box.go). |
| Button event handling   | o                | Supports both of touch and mouse click on components that implements `ButtonComponent` interface. See the [example](https://github.com/yohamta/furex/blob/master/examples/button/main.go). |
| Touch handler interface | o                | Able to handle touch ID on components that implements the `TouchHandler` interface.                                                                             |
| Mouse handler           | o                | Able to handle left click on components that implements the `MouseHandler` interface.                                                                                |
| Padding           | x                | To be implemented when needed.                                                     |
| Margin           | x                | To be implemented when needed.                                                      |

## Example
[Source Code](https://github.com/yohamta/furex/blob/master/examples/nesting/main.go)

<image src="https://user-images.githubusercontent.com/1475839/95682206-0279fa80-0c1f-11eb-8dd5-03bec58325e8.png" width="300px" />

## Simple Usage

```go
import "github.com/yohamta/furex"

var (
	colors = []color.Color{
		color.RGBA{0xff, 0, 0, 0xff},
		color.RGBA{0, 0xff, 0, 0xff},
		color.RGBA{0, 0, 0xff, 0xff},
	}
)

type Game struct {
	view *furex.View
}

// Initialize the UI
func (g *Game) initUI() {
	// Create flexbox container
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)

	// Set the options for flexbox
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	// Make flexbox items on flexbox container
	for i := 0; i < 20; i++ {
		// Each flexbox item must have fixed size so far.
		// In this case, the width is 50, height is 50.
		// Box component is just an example of a component.
		rootFlex.AddChild(NewBox(50, 50, colors[i%3]))
	}

	// Layer: A Layer can be stacked inside a View
	layer := furex.NewLayerWithContainer(rootFlex)

	// View: A View handles multiple layers of the UI and propagate UI events 
	//       such as touches or mouse click.
	g.view = furex.NewView()
	g.view.Layout(0, 0, screenWidth, screenHeight)

	// Add the layer to the View
	g.view.AddLayer(layer)
}

func (g *Game) Update() {
	// Update the UI 
	g.view.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the UI 
	g.view.Draw(screen)
}
```

### Result
<image src="https://user-images.githubusercontent.com/1475839/95682206-0279fa80-0c1f-11eb-8dd5-03bec58325e8.png" width="300px" />

## Use case
This [simple game](https://github.com/yohamta/godanmaku) is using furex for UI layout and interaction.

