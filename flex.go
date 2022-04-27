// Referenced code: https://github.com/golang/exp/blob/master/shiny/widget/flex/flex.go
package furex

import (
	"fmt"
	"image"
	"math"
)

// Direction is the direction in which flex items are laid out
type Direction uint8

const (
	Row Direction = iota
	Column
)

// Justify aligns items along the main axis.
type Justify uint8

const (
	JustifyStart        Justify = iota // pack to start of line
	JustifyEnd                         // pack to end of line
	JustifyCenter                      // pack to center of line
	JustifySpaceBetween                // even spacing
	JustifySpaceAround                 // even spacing, half-size on each end
)

// FlexAlign represents align of flex children
type FlexAlign int

const (
	FlexCenter FlexAlign = iota
	FlexStart
	FlexEnd
	FlexSpaceBetween
)

// AlignItem aligns items along the cross axis.
type AlignItem uint8

const (
	AlignItemStart AlignItem = iota
	AlignItemEnd
	AlignItemCenter
)

// FlexWrap controls whether the container is single- or multi-line,
// and the direction in which the lines are laid out.
type FlexWrap uint8

const (
	NoWrap FlexWrap = iota
	Wrap
	WrapReverse
)

// AlignContent is the 'align-content' property.
// It aligns container lines when there is extra space on the cross-axis.
type AlignContent uint8

const (
	AlignContentStart AlignContent = iota
	AlignContentEnd
	AlignContentCenter
	AlignContentSpaceBetween
	AlignContentSpaceAround
	AlignContentStretch
)

type Position uint8

const (
	PositionStatic Position = iota
	PositionAbsolute
)

type flexEmbed struct {
	*View
}

// layout is the main routine that implements a subset of flexbox layout
// https://www.w3.org/TR/css-flexbox-1/#layout-algorithm
func (f *flexEmbed) layout(width, height int, container *containerEmbed) {
	// 9.2. Line Length Determination
	// Determine the available main and cross space for the flex items.
	containerMainSize := float64(f.mainSize(width, height))
	containerCrossSize := float64(f.crossSize(width, height))

	// Determine the flex base size and hypothetical main size of each item:
	var children []element
	for _, c := range container.children {
		if c.item.Position == PositionAbsolute {
			c.bounds = image.Rect(
				c.item.Left, c.item.Top, c.item.Left+c.item.Width, c.item.Top+c.item.Height,
			)
			c.item.frame = c.bounds
			continue
		}
		children = append(children, element{
			flexBaseSize: float64(f.flexBaseSize(c)),
			node:         c,
		})
	}

	// §9.3. Main Size Determination
	// Collect flex items into flex lines
	var lines []flexLine
	if f.Wrap == NoWrap {
		// Single line
		line := flexLine{child: make([]*element, len(children))}
		for i := range children {
			child := &children[i]
			child.mainMargin = f.mainMargin(child.node)
			line.child[i] = child
			line.mainSize += child.flexBaseSize +
				(child.mainMargin[0] + child.mainMargin[1])
		}
		lines = []flexLine{line}
	} else {
		// Multi line
		var line flexLine
		for i := range children {
			child := &children[i]
			child.mainMargin = f.mainMargin(child.node)

			// hypotheticalMainSize = flexBaseSize + main margin
			hypotheticalMainSize := child.flexBaseSize +
				(child.mainMargin[0] + child.mainMargin[1])

			if line.mainSize > 0 && line.mainSize+hypotheticalMainSize > containerMainSize {
				lines = append(lines, line)
				line = flexLine{}
			}
			line.child = append(line.child, child)
			line.mainSize += hypotheticalMainSize
		}

		if len(line.child) > 0 || len(children) == 0 {
			lines = append(lines, line)
		}
	}

	for l := range lines {
		line := &lines[l]

		// Calculate free space
		freeSpace := float64(f.mainSize(width, height))
		for _, child := range line.child {
			freeSpace -= (float64(f.flexBaseSize(child.node)) +
				(child.mainMargin[0] + child.mainMargin[1]))
		}

		// Distribute free space
		for _, child := range line.child {
			child.mainSize = float64(f.flexBaseSize(child.node))
		}
	}

	// §9.4. Cross Size Determination
	// Determine the hypothetical cross size of each item
	for l := range lines {
		for _, c := range lines[l].child {
			if c.node.item.Width != 0 || c.node.item.Height != 0 {
				c.crossMargin = f.crossMargin(c.node)
				c.crossSize = float64(
					f.crossSize(c.node.item.Width, c.node.item.Height),
				)
			} else {
				panic("flex: size must be > 0")
			}
		}
	}

	// §9.4.8 Calculate the cross size of each flex line.
	if len(lines) == 1 {
		// Single line
		lines[0].crossSize = containerCrossSize
	} else {
		// Multi line
		for l := range lines {
			line := &lines[l]
			max := 0.0
			for _, child := range line.child {
				if child.crossSize > max {
					max = child.crossSize +
						(child.crossMargin[0] + child.crossMargin[1])
				}
			}
			line.crossSize = max
		}
	}

	off := 0.0
	for l := range lines {
		line := &lines[l]
		line.crossOffset = off
		off += line.crossSize
	}

	// §9.4.9 align-content: stretch
	remCrossSize := containerCrossSize - off
	if f.AlignContent == AlignContentStretch && remCrossSize > 0 {
		add := remCrossSize / float64(len(lines))
		for l := range lines {
			line := &lines[l]
			line.crossOffset += float64(l) * add
			line.crossSize += add
		}
	}

	// §9.5. Main-Axis Alignment
	for l := range lines {
		line := &lines[l]
		total := 0.0
		for _, child := range line.child {
			total += child.mainSize +
				(child.mainMargin[0] + child.mainMargin[1])
		}
		remFree := containerMainSize - total
		off, spacing := 0.0, 0.0
		switch f.Justify {
		case JustifyStart:
		case JustifyEnd:
			off = remFree
		case JustifyCenter:
			off = remFree / 2
		case JustifySpaceBetween:
			spacing = remFree / float64(len(line.child)-1)
		case JustifySpaceAround:
			spacing = remFree / float64(len(line.child))
			off = spacing / 2
		}
		for _, child := range line.child {
			child.mainOffset = off + (child.mainMargin[0])
			off += spacing + child.mainSize +
				(child.mainMargin[0] + child.mainMargin[1])
		}
	}

	// §9.6. Cross axis alignment
	for l := range lines {
		line := &lines[l]
		for _, child := range line.child {
			child.crossOffset = line.crossOffset + (child.crossMargin[0])
			if child.crossSize == line.crossSize {
				continue
			}
			diff := line.crossSize - child.crossSize -
				(child.crossMargin[0] + child.crossMargin[1])
			switch f.AlignItems {
			case AlignItemStart:
				// already laid out correctly
			case AlignItemEnd:
				child.crossOffset = line.crossOffset + diff +
					(child.crossMargin[0])
			case AlignItemCenter:
				child.crossOffset = line.crossOffset + diff/2 +
					(child.crossMargin[0])
			}
		}
	}

	// §9.6.15 determine container cross size used
	crossSize := lines[len(lines)-1].crossOffset + lines[len(lines)-1].crossSize
	remFree := containerCrossSize - crossSize

	// §9.6.16 align flex lines, 'align-content'.
	if remFree > 0 {
		spacing, off := 0.0, 0.0
		switch f.AlignContent {
		case AlignContentStart:
			// already laid out correctly
		case AlignContentEnd:
			off = remFree
		case AlignContentCenter:
			off = remFree / 2
		case AlignContentSpaceBetween:
			spacing = remFree / float64(len(lines)-1)
		case AlignContentSpaceAround:
			spacing = remFree / float64(len(lines))
			off = spacing / 2
		}
		if f.AlignContent != AlignContentStart {
			for l := range lines {
				line := &lines[l]
				line.crossOffset += off
				for _, child := range line.child {
					child.crossOffset += off
				}
				off += spacing
			}
		}
	}

	// Layout complete. Update children position
	for l := range lines {
		line := &lines[l]
		for _, child := range line.child {
			switch f.Direction {
			case Row:
				child.node.bounds = image.Rect(
					round(child.mainOffset),
					round(child.crossOffset),
					round(child.mainOffset+child.mainSize),
					round(child.crossOffset+child.crossSize))
				child.node.item.setFrame(child.node.bounds.Add(f.frame.Min))
			case Column:
				child.node.bounds = image.Rect(
					round(child.crossOffset),
					round(child.mainOffset),
					round(child.crossOffset+child.crossSize),
					round(child.mainOffset+child.mainSize))
				child.node.item.setFrame(child.node.bounds.Add(f.frame.Min))
			default:
				panic(fmt.Sprint("flex: bad direction ", f.Direction))
			}
		}
	}
}

type element struct {
	node         *child
	flexBaseSize float64
	mainSize     float64
	mainOffset   float64
	mainMargin   []float64
	crossSize    float64
	crossOffset  float64
	crossMargin  []float64
}

type flexLine struct {
	mainSize    float64
	crossSize   float64
	crossOffset float64
	child       []*element
}

func (f *flexEmbed) mainSize(x, y int) int {
	switch f.Direction {
	case Row:
		return x
	case Column:
		return y
	default:
		panic(fmt.Sprint("flex: bad direction ", f.Direction))
	}
}

func (f *flexEmbed) crossSize(x, y int) int {
	switch f.Direction {
	case Row:
		return y
	case Column:
		return x
	default:
		panic(fmt.Sprint("flex: bad direction ", f.Direction))
	}
}

func (f *flexEmbed) mainMargin(c *child) []float64 {
	switch f.Direction {
	case Row:
		return []float64{
			float64(c.item.MarginLeft),
			float64(c.item.MarginRight)}
	case Column:
		return []float64{
			float64(c.item.MarginTop),
			float64(c.item.MarginBottom)}
	default:
		panic("unreachable")
	}
}

func (f *flexEmbed) crossMargin(c *child) []float64 {
	switch f.Direction {
	case Row:
		return []float64{
			float64(c.item.MarginTop),
			float64(c.item.MarginBottom)}
	case Column:
		return []float64{
			float64(c.item.MarginLeft),
			float64(c.item.MarginRight)}
	default:
		panic("unreachable")
	}
}

func (f *flexEmbed) flexBaseSize(c *child) int {
	if c.item.Width != 0 || c.item.Height != 0 {
		return f.mainSize(c.item.Width, c.item.Height)
	}
	panic("flex: size must be > 0")
}

func round(f float64) int {
	return int(math.Floor(f + .5))
}
