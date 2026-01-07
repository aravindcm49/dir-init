package models

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Item represents a selectable item
type Item struct {
	Code        string
	Description string
	IsCustom    bool
}

// Mode represents the current UI mode
type Mode int

const (
	ModeSearch Mode = iota
	ModeCustomConfirm
	ModeCustomInput
)

// SelectorModel is the Bubble Tea model for selection
type SelectorModel struct {
	title       string
	searchInput textinput.Model
	descInput   textinput.Model
	allItems    []Item
	filtered    []Item
	cursor      int
	selected    *Item
	mode        Mode

	// Custom input state
	customCode string
	customDesc string
	saveChoice int // 0=use once, 1=save

	// State
	quitting  bool
	confirmed bool
}

// NewSelector creates a new selector model
func NewSelector(title string, items []Item) SelectorModel {
	ti := textinput.New()
	ti.Placeholder = "Type to search..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 50

	di := textinput.New()
	di.Placeholder = "Enter description..."
	di.CharLimit = 100
	di.Width = 50

	return SelectorModel{
		title:       title,
		searchInput: ti,
		descInput:   di,
		allItems:    items,
		filtered:    items,
		mode:        ModeSearch,
		cursor:      0,
	}
}

// Init initializes the model
func (m SelectorModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages
func (m SelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.mode {
		case ModeSearch:
			return m.updateSearch(msg)
		case ModeCustomConfirm:
			return m.updateCustomConfirm(msg)
		case ModeCustomInput:
			return m.updateCustomInput(msg)
		}
	}
	return m, nil
}

func (m SelectorModel) updateSearch(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyEsc:
		m.quitting = true
		return m, tea.Quit

	case tea.KeyEnter:
		if len(m.filtered) > 0 {
			// Select current item
			m.selected = &m.filtered[m.cursor]
			m.confirmed = true
			return m, tea.Quit
		} else if m.searchInput.Value() != "" {
			// No matches - create custom (one-time use)
			m.customCode = m.searchInput.Value()
			m.saveChoice = 0
			m.selected = &Item{
				Code:        m.customCode,
				Description: m.customCode,
				IsCustom:    true,
			}
			m.confirmed = true
			return m, tea.Quit
		}

	case tea.KeyCtrlS:
		if len(m.filtered) == 0 && m.searchInput.Value() != "" {
			// No matches - create custom and save to config
			m.customCode = m.searchInput.Value()
			m.saveChoice = 1
			m.mode = ModeCustomInput
			m.descInput.Focus()
			return m, textinput.Blink
		}

	case tea.KeyUp, tea.KeyRight:
		// Increment cursor (move selection forward)
		if m.cursor < len(m.filtered)-1 {
			m.cursor++
		}

	case tea.KeyDown, tea.KeyLeft:
		// Decrement cursor (move selection backward)
		if m.cursor > 0 {
			m.cursor--
		}

	default:
		// Only treat as search input if it's a printable character
		if msg.Type == tea.KeyRunes {
			// Update search input
			var cmd tea.Cmd
			m.searchInput, cmd = m.searchInput.Update(msg)

			// Filter items
			m.filterItems()
			m.cursor = 0

			return m, cmd
		}
	}

	return m, nil
}

func (m SelectorModel) updateCustomConfirm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		// Enter = Use once (don't save) - skip description
		m.saveChoice = 0
		m.selected = &Item{
			Code:        m.customCode,
			Description: m.customCode, // Use code as description for one-time items
			IsCustom:    true,
		}
		m.confirmed = true
		return m, tea.Quit

	case tea.KeyCtrlS:
		// Ctrl+S = Save to config - ask for description
		m.saveChoice = 1
		m.mode = ModeCustomInput
		m.descInput.Focus()
		return m, textinput.Blink

	case tea.KeyEsc:
		m.mode = ModeSearch
		m.searchInput.Focus()
		return m, textinput.Blink
	}

	return m, nil
}

func (m SelectorModel) updateCustomInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEnter:
		// Confirm custom input
		m.customDesc = m.descInput.Value()
		m.selected = &Item{
			Code:        m.customCode,
			Description: m.customDesc,
			IsCustom:    true,
		}
		m.confirmed = true
		return m, tea.Quit

	case tea.KeyEsc:
		m.mode = ModeCustomConfirm
		return m, nil
	}

	// Update description input
	var cmd tea.Cmd
	m.descInput, cmd = m.descInput.Update(msg)
	return m, cmd
}

func (m *SelectorModel) filterItems() {
	query := strings.ToLower(m.searchInput.Value())
	if query == "" {
		m.filtered = m.allItems
		return
	}

	m.filtered = []Item{}
	for _, item := range m.allItems {
		text := strings.ToLower(item.Code + " " + item.Description)
		if strings.Contains(text, query) {
			m.filtered = append(m.filtered, item)
		}
	}
}

// View renders the UI
func (m SelectorModel) View() string {
	if m.quitting || m.confirmed {
		return ""
	}

	var s strings.Builder

	// Title
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("220"))
	s.WriteString(titleStyle.Render(m.title) + "\n\n")

	switch m.mode {
	case ModeSearch:
		s.WriteString(m.viewSearch())
	case ModeCustomConfirm:
		s.WriteString(m.viewCustomConfirm())
	case ModeCustomInput:
		s.WriteString(m.viewCustomInput())
	}

	return s.String()
}

func (m SelectorModel) viewSearch() string {
	var s strings.Builder

	// Search input
	s.WriteString(m.searchInput.View() + "\n\n")

	// Filtered items
	if len(m.filtered) == 0 {
		hintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		if m.searchInput.Value() != "" {
			s.WriteString(hintStyle.Render("No matches found. Press Enter = use once | Ctrl+S = save to config\n"))
		}
	} else {
		for i, item := range m.filtered {
			if i == m.cursor {
				selectedItemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Bold(true)
				s.WriteString(selectedItemStyle.Render("> " + item.Code + " - " + item.Description))
			} else {
				itemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
				s.WriteString(itemStyle.Render("  " + item.Code + " - " + item.Description))
			}
			s.WriteString("\n")
		}
	}

	return s.String()
}

func (m SelectorModel) viewCustomConfirm() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Create '%s'?\n", m.customCode))
	hintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	s.WriteString(hintStyle.Render("Press Enter = use once | Ctrl+S = save to config\n"))

	return s.String()
}

func (m SelectorModel) viewCustomInput() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Code: %s\n", m.customCode))
	s.WriteString("Description: " + m.descInput.View() + "\n")

	return s.String()
}

// GetSelected returns the selected item
func (m SelectorModel) GetSelected() *Item {
	return m.selected
}

// ShouldSave returns whether to save to config
func (m SelectorModel) ShouldSave() bool {
	return m.saveChoice == 1
}
