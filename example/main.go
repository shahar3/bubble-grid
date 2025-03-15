package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shahar3/bubble-grid/frame"
	"github.com/shahar3/bubble-grid/grid"
)

// SimpleItem is a basic implementation of GridItem
type SimpleItem struct {
	content string
}

func (s SimpleItem) Render() string {
	color := lipgloss.Color("#874BFD")
	return lipgloss.NewStyle().Background(color).Render(s.content)
}

// ExampleModel demonstrates how to use the StackedGrid
type ExampleModel struct {
	grid *grid.StackedGrid
}

func (m ExampleModel) Init() tea.Cmd {
	return nil
}

func (m ExampleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.grid.SetSize(msg.Width, msg.Height)
	}
	return m, nil
}

func (m ExampleModel) View() string {
	return m.grid.Render()
}

func BasicGridExample() {
	// Create a new grid with 3 columns
	g := grid.NewStackedGrid()

	// Add a simple item
	g.AddItem(SimpleItem{"Item 1"}, grid.ItemOptions{
		Column: 0,
	})

	// Add another simple item
	g.AddItem(SimpleItem{"Item 2"}, grid.ItemOptions{
		Column: 1,
	})

	// Add another simple item
	g.AddItem(SimpleItem{"Item 3"}, grid.ItemOptions{
		Column: 2,
	})

	model := ExampleModel{grid: g}
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

func GridWithFrameExample() {
	// Create a new grid with 3 columns
	g := grid.NewStackedGrid()

	// Add a simple item with a frame
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 1"}), grid.ItemOptions{
		Column: 0,
	})

	// Add a simple item with a frame
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 2"}), grid.ItemOptions{
		Column: 1,
	})

	// Add another framed item
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 3"}), grid.ItemOptions{
		Column: 2,
	})

	// Add a framed item in the second row
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 4"}), grid.ItemOptions{
		Column: 1,
	})

	model := ExampleModel{grid: g}
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

func GridWithExpandedFrameExample() {
	// Create a new grid with 3 columns
	g := grid.NewStackedGrid()

	// Add a simple item with a frame
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 1"}), grid.ItemOptions{
		Column: 0,
	})

	// Add a simple item with a frame
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 2"}), grid.ItemOptions{
		Column:         1,
		ExpandVertical: true,
	})

	// Add a framed item in the second row
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 3"}), grid.ItemOptions{
		Column:         1,
		ExpandVertical: true,
	})

	// Add an expanded framed item in the second row
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 4"}), grid.ItemOptions{
		Column:         1,
		ExpandVertical: false,
	})

	// Add another framed item
	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 5"}), grid.ItemOptions{
		Column: 2,
	})

	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 6"}), grid.ItemOptions{
		Column: 2,
	})

	g.AddItem(frame.NewFrame(SimpleItem{"Framed Item 7"}), grid.ItemOptions{
		Column: 2,
	})

	model := ExampleModel{grid: g}
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

func main() {
	GridWithExpandedFrameExample()
}
