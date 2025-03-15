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
		// count the number of items that have ExpandVertical set to true
		expandingItems := 0
		for _, item := range items {
			if item.Options.ExpandVertical {
				expandingItems++
			}
		}

		var heightPerItem int
		var heightPerExpandedItem int
		var remainderHeight int
		if expandingItems > 0 {
			// If there are expanding items, we need to calculate the height per item
			heightPerExpandedItem, remainderHeight = g.calculateHeightPerExpandedItem(items, expandingItems)
			heightPerItem = 1
		} else {
			// Calculate the height per item
			heightPerItem = g.height / len(items)
		}

		// Get the available height for each item
		totalLines := 0
		for i, item := range items {
			if i == len(items)-1 {
				if item.Options.ExpandVertical {
					heightPerExpandedItem += remainderHeight
				} else {
					if heightPerItem > 1 {
						remainderHeight := g.height % len(items)
						heightPerItem += remainderHeight
					} else {
						heightPerItem = g.getItemNaturalHeight(item.Item) + remainderHeight
					}
				}
			}

			// Check if the item is a Frame using type assertion
			if frame, ok := item.Item.(Sizer); ok {
				if item.Options.ExpandVertical {
					renderedItems[i] = frame.SetSize(colWidth, heightPerExpandedItem).Render()
					totalLines += heightPerExpandedItem
				} else {
					renderedItems[i] = frame.SetSize(colWidth, heightPerItem).Render()
					totalLines += heightPerItem
				}
			} else {
				if item.Options.ExpandVertical {
					renderedItems[i] = lipgloss.NewStyle().Height(heightPerExpandedItem).Width(colWidth).Render(item.Render())
					totalLines += heightPerExpandedItem
				} else {
					renderedItems[i] = lipgloss.NewStyle().Height(heightPerItem).Width(colWidth).Render(item.Render())
					totalLines += heightPerItem
				}
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

// calculateHeightPerExpandedItem calculates the height per expanded item and the remainder height
func (g *StackedGrid) calculateHeightPerExpandedItem(items []ItemWithOptions, expandingItems int) (int, int) {
	// Get the natural height of all the items that are not expanding
	naturalHeight := 0
	for _, item := range items {
		if !item.Options.ExpandVertical {
			naturalHeight += g.getItemNaturalHeight(item)
		}
	}

	return (g.height - naturalHeight) / expandingItems, (g.height - naturalHeight) % expandingItems
}

func (g *StackedGrid) SetSize(width, height int) {
	g.width = width
	g.height = height
}

func (g *StackedGrid) getItemNaturalHeight(item Item) int {
	// Count the number of lines in the item's render
	lines := strings.Split(item.Render(), "\n")
	return len(lines)
}
