package furex

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2/internal/graphic"
)

var (
	Debug           = false
	debugColor      = color.RGBA{0xff, 0, 0, 0xff}
	debugColorShift = ebiten.ColorM{}
)

func debugBorders(screen *ebiten.Image, root containerEmbed) {
	queue := []containerEmbed{}
	queue = append(queue, root)
	renderColor := resetDebugColor()

	for len(queue) > 0 {
		levelSize := len(queue)
		for levelSize != 0 {
			curr := queue[0]
			queue = queue[1:]

			graphic.DrawRect(screen, &graphic.DrawRectOpts{
				Rect:        curr.frame,
				Color:       renderColor,
				StrokeWidth: 2,
			})

			for _, c := range curr.children {
				queue = append(queue, c.item.containerEmbed)
			}
			levelSize--
		}

		renderColor = rotateDebugColor()
	}
}

func rotateDebugColor() color.Color {
	debugColorShift.RotateHue(1.66)
	return debugColorShift.Apply(debugColor)
}

func resetDebugColor() color.Color {
	debugColorShift = ebiten.ColorM{}
	return debugColor
}
