# Bubble Grid

A flexible stacked grid layout component for [Bubble Tea](https://github.com/charmbracelet/bubbletea) TUI applications.

## Features

- Create multi-column grid layouts
- Stack items vertically within columns
- Configure item placement and expansion
- Support for nested grids
- Scrollable content with viewport support

## Installation

```bash
go get github.com/shahar3/bubble-grid
```

## Usage

Here's a simple example of how to use the grid:

```go
package main

import (
    "github.com/shahar3/bubble-grid"
)

// Create a component that implements GridItem
type MyItem struct {
    content string
}

func (m MyItem) Render() string {
    return m.content
}

func main() {
    // Create a new grid with 3 columns
    g := grid.New(3)

    // Add items to the grid
    g.AddItem(MyItem{"Item 1"}, grid.GridOptions{
        Column: 0,
        ExpandVertical: false,
        MinHeight: 1,
    })

    // Add more items...
}
```

### Grid Options

When adding items to the grid, you can specify the following options:

- `Column`: Which column to place the item in (0-based index)
- `ExpandVertical`: Whether the item should expand to fill available vertical space
- `MinHeight`: Minimum height for the item
- `MaxHeight`: Maximum height for the item (0 means unlimited)

### Nested Grids

You can create nested grids by adding a grid as an item to another grid:

```go
// Create main grid
mainGrid := grid.New(2)

// Create nested grid
nestedGrid := grid.New(2)
nestedGrid.AddItem(MyItem{"Nested 1"}, grid.GridOptions{...})

// Add nested grid to main grid
mainGrid.AddItem(nestedGrid, grid.GridOptions{
    Column: 0,
    ExpandVertical: true,
})
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License
