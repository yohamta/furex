# furex
A simple UI framework with a subset of flexbox layout for [Ebiten](https://ebiten.org/).

## Features
| Feature                 | Supported | Note                                                                                                                   |
|-------------------------|------------------|------------------------------------------------------------------------------------------------------------------------|
| Flexbox layout          | o                | Supports a subset of flexbox layout spec.                                                                              |
| Button event handling   | o                | Supports both of touch and mouse. See [example](https://github.com/yohamta/furex/blob/master/examples/button/main.go). |
| Touch handler interface | o                | Able to handle each touch ID on component.                                                                             |
| Mouse handler           | o                | Able to handle left click on component.                                                                                |

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
	// Create flex container
	rootFlex := furex.NewFlex(0, 0, screenWidth, screenHeight)

	// Set the options for flex layout
	rootFlex.Direction = furex.Row
	rootFlex.Justify = furex.JustifyCenter
	rootFlex.AlignItems = furex.AlignItemCenter
	rootFlex.AlignContent = furex.AlignContentCenter
	rootFlex.Wrap = furex.Wrap

	// Make flex children
	for i := 0; i < 20; i++ {
		// Each flex children must have fixed size (width and height) so far
		// In this example the width is 50 and the height is 50
		rootFlex.AddChild(furex.NewBox(50, 50, colors[i%3]))
	}

	// Layer: A layer can be stacked on other layers
	//        so you can make complex UI with multiple layers.
	layer := furex.NewLayerWithContainer(rootFlex)

	// View: A view handles multiple layers of the UI
	//                  and also the UI events such as touches or mouse click.
	g.view = furex.NewView()
	g.view.Layout(0, 0, screenWidth, screenHeight)

	// Add the layer to the view
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
This simple game is using furex for UI interaction.
https://github.com/yohamta/godanmaku
