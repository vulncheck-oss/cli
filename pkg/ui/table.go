package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/vulncheck-oss/sdk"
	"golang.org/x/term"
	"strings"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#6667ab"))

type tableModel struct {
	table  table.Model
	action func(index string) error
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if err := m.action(m.table.SelectedRow()[0]); err != nil {
				return m, tea.Quit
			}
			return m, tea.Quit
			/*
				return m, tea.Batch(
					m.action(m.table.SelectedRow()[0]),
				)
			*/
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func IndicesRows(indices []sdk.IndicesMeta, search string) []table.Row {
	var rows []table.Row
	for _, index := range indices {
		if search != "" && !strings.Contains(index.Name, search) && !strings.Contains(index.Description, search) {
			continue
		}
		rows = append(rows, table.Row{
			index.Name,
			index.Description,
			index.Href,
		})
	}
	return rows
}

func IndicesBrowse(indices []sdk.IndicesMeta, search string, action func(index string) error) error {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Description", Width: TermWidth() - 52},
		{Title: "URL", Width: 20},
	}

	rows := IndicesRows(indices, search)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(TermHeight()-10),
		table.WithWidth(TermWidth()-5),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6667ab")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#34d399")).
		Bold(false)
	t.SetStyles(s)

	m := tableModel{t, action}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running program: %v", err)
	}

	return nil
}

func IndicesList(indices []sdk.IndicesMeta, search string) error {

	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("Name", "Description", "Href").Width(TermWidth())

	for _, index := range indices {
		if search != "" && !strings.Contains(index.Name, search) && !strings.Contains(index.Description, search) {
			continue
		}
		t.Row(index.Name, index.Description, index.Href)
	}

	fmt.Println(t)
	return nil
}

func TermWidth() int {
	width, _, _ := term.GetSize(0)
	return width
}

func TermHeight() int {
	_, height, _ := term.GetSize(0)
	return height
}

func CpeMeta(cpe sdk.CpeMeta) error {
	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("Part", "Vendor", "Product", "Version", "Update", "Edition").
		Row(cpe.Part, cpe.Vendor, cpe.Product, cpe.Version, cpe.Update, cpe.Edition).Width(TermWidth())
	fmt.Println(t)
	return nil
}

func PurlMeta(purl sdk.PurlMeta) error {
	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("Type", "Namespace", "Nme", "Version", "Qualifiers", "Subpath").
		Row(purl.Type, purl.Namespace, purl.Name, purl.Version, strings.Join(purl.Qualifiers, ","), purl.Subpath).
		Width(TermWidth())
	fmt.Println(t)
	return nil
}

func SbomList(sboms []sdk.SbomMeta) error {
	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("ID", "Name", "Repository", "Created At", "Updated At").
		Width(TermWidth())

	for _, sbom := range sboms {
		t.Row(sbom.Id, sbom.Name, sbom.Repository, sbom.CreatedAt, sbom.UpdatedAt)
	}
	fmt.Println(t)
	return nil
}

func SbomScans(scans []sdk.SbomScanMeta) error {
	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("ID", "Status", "Created At", "Updated At").
		Width(TermWidth())

	for _, scan := range scans {
		t.Row(scan.Id, scan.Status, scan.CreatedAt, scan.UpdatedAt)
	}
	fmt.Println(t)
	return nil
}
