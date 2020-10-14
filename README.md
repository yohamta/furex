# furex
A simple UI framework with a subset of flexbox layout for [Ebiten](https://ebiten.org/).

## Example Usage

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
	view *furex.Controller
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

	// view controller: A controller handles multiple layers of the UI
	//                  and also the UI events such as touches or mouse click.
	g.view = furex.NewController()
	g.view.Layout(0, 0, screenWidth, screenHeight)

	// Add the layer to the view controller
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
![image](https://user-images.githubusercontent.com/1475839/95682206-0279fa80-0c1f-11eb-8dd5-03bec58325e8.png)