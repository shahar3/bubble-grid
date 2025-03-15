package grid

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Item represents an item that can be rendered in the grid
type Item interface {
	// Render returns the string representation of the item
	Render() string
}

// Sizer represents an item that can be sized
type Sizer interface {
	Item
	SetSize(width, height int) Sizer
}

type Options struct {
	FitScreen bool
}

// ItemOptions defines the configuration for a grid item
type ItemOptions struct {
	// Column specifies which column the item should be placed in (0-based)
	Column int
	// ExpandVertical determines if the item should expand to fill available vertical space
	ExpandVertical bool
}

type ItemWithOptions struct {
	Item
	Options ItemOptions
}

type StackedGrid struct {
	items   []ItemWithOptions
	width   int
	height  int
	options Options
}

func NewStackedGrid() *StackedGrid {
	return &StackedGrid{
		options: Options{
			FitScreen: true,
		},
	}
}

func NewStackedGridWithOptions(options Options) *StackedGrid {
	return &StackedGrid{
		options: options,
	}
}

func (g *StackedGrid) AddItem(item Item, options ItemOptions) {
	g.items = append(g.items, ItemWithOptions{
		Item:    item,
		Options: options,
	})
}

func (g *StackedGrid) Render() string {
	if g.width == 0 || g.height == 0 {
		return "Loading..."
	}

	// Group items by column
	columns := make(map[int][]ItemWithOptions)
	for _, item := range g.items {
		columns[item.Options.Column] = append(columns[item.Options.Column], item)
	}

	columnsOutput := g.RenderColumns(columns)

	// Join the columns horizontally with Lipgloss
	return lipgloss.JoinHorizontal(lipgloss.Top, columnsOutput...)
}

func (g *StackedGrid) RenderColumns(columns map[int][]ItemWithOptions) []string {
	renderedColumns := make([]string, len(columns))
	for column, items := range columns {
		// Get the width of the column
		colWidth := g.width / len(columns)
		// Generate the column with the width
		renderedColumns[column] = g.RenderColumn(items, colWidth)
	}

	return renderedColumns
}

func (g *StackedGrid) RenderColumn(items []ItemWithOptions, colWidth int) string {
	renderedItems := make([]string, len(items))
	// Check the grid options to see if we need to fit the column to the screen
	if g.options.FitScreen {
		// Calculate the height per item
		heightPerItem := g.height / len(items)
		// In case there is a remainder, add it to the last item
		remainderHeight := g.height % len(items)
		// Get the available height for each item
		for i, item := range items {
			if i == len(items)-1 {
				heightPerItem += remainderHeight
			}

			// Check if the item is a Frame using type assertion
			if frame, ok := item.Item.(Sizer); ok {
				renderedItems[i] = frame.SetSize(colWidth, heightPerItem).Render()
			} else {
				renderedItems[i] = lipgloss.NewStyle().Height(heightPerItem).Width(colWidth).Render(item.Render())
			}
		}

		return lipgloss.NewStyle().Width(colWidth).Render(strings.Join(renderedItems, "\n"))
	} else {
		for i, item := range items {
			renderedItems[i] = item.Render()
		}
	}

	return strings.Join(renderedItems, "\n")
}

func (g *StackedGrid) SetSize(width, height int) {
	g.width = width
	g.height = height
}
