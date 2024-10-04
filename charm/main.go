package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	config       *Configuration
	filteredSubs []Subscription
	selected     int
	filter       textinput.Model
	err          error
}

var (
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	headerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("203"))
	defaultStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("35"))
	matchStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
)

const (
	cursorWidth    = 2
	checkmarkWidth = 1
	gapWidth       = 1
	nameWidth      = 50
	idWidth        = 36
)

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Filter subscriptions..."
	ti.Focus()

	return model{
		filter:   ti,
		selected: 0,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if len(m.filteredSubs) > 0 {
				m.config.Subscriptions[m.selected].IsDefault = true
				return m, tea.Quit
			}
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if m.selected < len(m.filteredSubs)-1 {
				m.selected++
			}
		}

	case error:
		m.err = msg
		return m, nil
	}

	m.filter, cmd = m.filter.Update(msg)
	m.updateFilteredSubs()
	return m, cmd
}

func (m *model) updateFilteredSubs() {
	filter := strings.ToLower(m.filter.Value())
	m.filteredSubs = []Subscription{}
	for _, sub := range m.config.Subscriptions {
		if strings.Contains(strings.ToLower(sub.Name), filter) || strings.Contains(strings.ToLower(sub.ID.String()), filter) {
			m.filteredSubs = append(m.filteredSubs, sub)
		}
	}
	if m.selected >= len(m.filteredSubs) {
		m.selected = len(m.filteredSubs) - 1
	}
	if m.selected < 0 {
		m.selected = 0
	}
}

func highlightMatches(text, filter string) string {
	if filter == "" {
		return normalStyle.Render(text)
	}

	lowerText := strings.ToLower(text)
	lowerFilter := strings.ToLower(filter)
	var result strings.Builder

	for i := 0; i < len(text); {
		if i+len(lowerFilter) <= len(lowerText) && lowerText[i:i+len(lowerFilter)] == lowerFilter {
			result.WriteString(matchStyle.Render(text[i : i+len(filter)]))
			i += len(filter)
		} else {
			result.WriteString(normalStyle.Render(string(text[i])))
			i++
		}
	}

	return result.String()
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	s := titleStyle.Render("aztx") + "\n\n"
	s += m.filter.View() + "\n\n"

	// Table header
	nameHeader := headerStyle.Render(fmt.Sprintf("%-*s", nameWidth, "NAME"))
	idHeader := headerStyle.Render(fmt.Sprintf("%-*s", idWidth, "ID"))
	s += fmt.Sprintf("%*s%*s%s\n",
		cursorWidth+checkmarkWidth+gapWidth, "", // Space for cursor, checkmark, and gap
		nameWidth, nameHeader,
		idHeader)

	// Table rows
	for i, sub := range m.filteredSubs {
		cursor := "  "
		if i == m.selected {
			cursor = selectedStyle.Render("> ")
		}
		checkmark := " "
		if sub.IsDefault {
			checkmark = defaultStyle.Render("*")
		}
		name := lipgloss.NewStyle().Width(nameWidth).Render(highlightMatches(sub.Name, m.filter.Value()))
		id := lipgloss.NewStyle().Width(idWidth).Render(highlightMatches(sub.ID.String(), m.filter.Value()))
		s += fmt.Sprintf("%s%s%*s%s%s\n",
			cursor, checkmark,
			gapWidth, "",
			name,
			id)
	}

	s += "\n" + normalStyle.Render("Type to filter. Use arrow keys to navigate. Press Enter to select.") + "\n"

	return s
}

func loadConfig() (*Configuration, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %w", err)
	}

	configPath := filepath.Join(home, ".azure", "azureProfile.json")
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Remove BOM if present
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	var config Configuration
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}

func saveConfig(config *Configuration) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	configPath := filepath.Join(home, ".azure", "azureProfile.json")
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		fmt.Println("Please ensure your azureProfile.json file is properly formatted and encoded as UTF-8.")
		os.Exit(1)
	}

	m := initialModel()
	m.config = config
	m.filteredSubs = config.Subscriptions

	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}

	if m, ok := finalModel.(model); ok && len(m.filteredSubs) > 0 {
		selectedSub := m.filteredSubs[m.selected]
		for i := range m.config.Subscriptions {
			m.config.Subscriptions[i].IsDefault = (m.config.Subscriptions[i].ID == selectedSub.ID)
		}

		if err := saveConfig(m.config); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Switched to \"%s\" (%s)\n", selectedSub.Name, selectedSub.ID)
	}
}
