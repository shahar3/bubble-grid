package frame

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/shahar3/bubble-grid/grid"
)

// Frame is a component that wraps content with a border and padding
type Frame struct {
	content grid.Item
	style   lipgloss.Style
	width   int
	height  int
}

func NewFrame(content grid.Item) Frame {
	return Frame{
		content: content,
		style: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1),
		width:  0, // 0 means auto-size
		height: 0, // 0 means auto-size
	}
}

func (f Frame) Render() string {
	if f.width > 0 && f.height > 0 {
		innerWidth, innerHeight := f.GetInnerSize()

		// If the content is also a Sizer, propagate the inner dimensions
		if sizer, ok := f.content.(grid.Sizer); ok {
			innerContent := sizer.SetSize(innerWidth, innerHeight).Render()
			return f.style.Width(f.width).Height(f.height).Render(innerContent)
		}

		// Otherwise, just render the content with the available space
		return f.style.Width(innerWidth).Height(innerHeight).Render(f.content.Render())
	}
	return f.style.Render(f.content.Render())
}

func (f Frame) ChangeBorderColor(color lipgloss.Color) Frame {
	f.style = f.style.BorderForeground(color)
	return f
}

// SetSize sets the frame's width and height
func (f Frame) SetSize(width, height int) grid.Sizer {
	f.width = width
	f.height = height
	return f
}

// GetInnerSize returns the available size for content after accounting for borders and padding
func (f Frame) GetInnerSize() (width, height int) {
	if f.width == 0 || f.height == 0 {
		return 0, 0
	}

	horizontalSpace := f.style.GetHorizontalFrameSize() - 2
	verticalSpace := f.style.GetVerticalFrameSize() - 2

	return f.width - horizontalSpace, f.height - verticalSpace
}
