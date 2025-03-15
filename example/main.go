package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
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

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list    list.Model
	grid    *grid.StackedGrid
	example int // -1 for menu, 0-3 for examples
	width   int
	height  int
}

func initialModel() model {
	items := []list.Item{
		item{title: "Basic Grid", desc: "Simple grid layout example"},
		item{title: "Framed Grid", desc: "Grid with framed components"},
		item{title: "Nested Grid", desc: "Grid with nested grid components"},
		item{title: "Complex Layout", desc: "Advanced grid layout example"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "BubbleGrid Examples"
	l.SetShowHelp(false)

	return model{
		list:    l,
		example: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "backspace":
			if m.example != -1 {
				m.example = -1
				m.list.SetWidth(m.width)
				m.list.SetHeight(m.height)
				return m, nil
			}
		case "enter":
			if m.example == -1 {
				m.example = m.list.Index()
				m.grid = getExampleGrid(m.example)
				m.grid.SetSize(m.width, m.height)
				return m, nil
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.example == -1 {
			m.list.SetWidth(msg.Width)
			m.list.SetHeight(msg.Height)
		} else {
			m.grid.SetSize(msg.Width, msg.Height)
		}
	}

	if m.example == -1 {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	if m.example == -1 {
		return m.list.View()
	}
	return m.grid.Render()
}

func getExampleGrid(example int) *grid.StackedGrid {
	g := grid.NewStackedGrid()

	switch example {
	case 0:
		return basicGridExample()
	case 1:
		return framedGridExample()
	case 2:
		return nestedGridExample()
	case 3:
		return complexLayoutExample()
	}

	return g
}

func basicGridExample() *grid.StackedGrid {
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
	return g
}

func framedGridExample() *grid.StackedGrid {
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
	return g
}

func nestedGridExample() *grid.StackedGrid {
	outer := grid.NewStackedGrid()
	inner := grid.NewStackedGrid()

	inner.AddItem(frame.NewFrame(SimpleItem{"Nested 1"}), grid.ItemOptions{Column: 0})
	inner.AddItem(frame.NewFrame(SimpleItem{"Nested 2"}), grid.ItemOptions{Column: 1})

	outer.AddItem(inner, grid.ItemOptions{Column: 0})
	outer.AddItem(frame.NewFrame(SimpleItem{"Main Grid"}), grid.ItemOptions{Column: 0})
	return outer
}

func complexLayoutExample() *grid.StackedGrid {
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

	// Add a nested grid
	nested := grid.NewStackedGrid()
	nested.AddItem(frame.NewFrame(SimpleItem{"Nested 1"}), grid.ItemOptions{Column: 0, ExpandVertical: true})
	nested.AddItem(frame.NewFrame(SimpleItem{"Nested 2"}), grid.ItemOptions{Column: 1, ExpandVertical: true})
	nested.AddItem(frame.NewFrame(SimpleItem{"Nested 3"}), grid.ItemOptions{Column: 1})
	g.AddItem(nested, grid.ItemOptions{
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

	return g
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
