package ui

import (
	"fmt"
	"github.com/package-url/packageurl-go"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/sdk-go"
	"golang.org/x/term"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#6667ab"))

type tableModel struct {
	table       table.Model
	selectedID  string
	quitting    bool
	createEntry bool
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			os.Exit(0)
			return m, tea.Quit
		case "enter":
			m.selectedID = m.table.SelectedRow()[0]
			return m, tea.Quit
		case "c":
			if m.createEntry {
				m.selectedID = "createEntry"
				return m, tea.Quit
			}
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	return baseStyle.Render(m.table.View())
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

func IndicesBrowse(indices []sdk.IndicesMeta, search string) (string, error) {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Description", Width: TermWidth() - 52},
		{Title: "URL", Width: 20},
	}

	rows := IndicesRows(indices, search)

	m := newTableModel(columns, rows, false)

	p := tea.NewProgram(m)
	finalModel, err := p.Run()

	if err != nil {
		return "", fmt.Errorf("error running program: %v", err)
	}

	if finalModel, ok := finalModel.(tableModel); ok {
		if finalModel.quitting {
			return "", nil
		}
		return finalModel.selectedID, nil
	}

	return "", fmt.Errorf("unexpected model type")
}

func TokensRows(tokens []sdk.TokenData) []table.Row {
	var rows []table.Row
	for _, token := range tokens {
		rows = append(rows, table.Row{
			token.ID,
			token.GetSourceLabel(),
			token.GetLocationString(),
			token.GetHumanUpdatedAt(),
		})
	}
	return rows
}

func newTableModel(columns []table.Column, rows []table.Row, createEntry bool) tableModel {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(TermHeight()-11),
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

	return tableModel{table: t, createEntry: createEntry}
}

func TokensBrowse(tokens []sdk.TokenData) (string, error) {
	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Source", Width: 30},
		{Title: "Location", Width: TermWidth() - 67},
		{Title: "Last Activity", Width: 15},
	}

	rows := TokensRows(tokens)

	m := newTableModel(columns, rows, true)
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("error running program: %v", err)
	}

	if finalModel, ok := finalModel.(tableModel); ok {
		if finalModel.quitting {
			return "", nil
		}
		return finalModel.selectedID, nil
	}

	return "", fmt.Errorf("unexpected model type")
}

func TokensList(tokens []sdk.TokenData) error {

	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("ID", "Source", "Location", "Last Activity").Width(TermWidth())

	for _, token := range tokens {
		t.Row(token.ID, token.GetSourceLabel(), token.GetLocationString(), token.GetHumanUpdatedAt())
	}

	fmt.Println(t)
	return nil
}

func IndicesList(indices []sdk.IndicesMeta, search string) error {

	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
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
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("Part", "Vendor", "Product", "Version", "Update", "Edition").
		Row(cpe.Part, cpe.Vendor, cpe.Product, cpe.Version, cpe.Update, cpe.Edition).Width(TermWidth())
	fmt.Println(t)
	return nil
}

func PurlMeta(purl sdk.PurlMeta) error {
	t := ltable.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("Type", "Namespace", "Nme", "Version", "Qualifiers", "Subpath").
		Row(purl.Type, purl.Namespace, purl.Name, purl.Version, strings.Join(purl.Qualifiers, ","), purl.Subpath).
		Width(TermWidth())
	fmt.Println(t)
	return nil
}

func PurlInstance(purl packageurl.PackageURL) error {
	qualifiers := make([]string, 0, len(purl.Qualifiers))
	t := ltable.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("Type", "Namespace", "Name", "Version", "Qualifiers", "Subpath").
		Row(purl.Type, purl.Namespace, purl.Name, purl.Version, strings.Join(qualifiers, ","), purl.Subpath).
		Width(TermWidth())
	fmt.Println(t)
	return nil
}

func PurlVulns(vulns []sdk.PurlVulnerability) error {
	t := ltable.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("Detection", "Fixed Version").
		Width(TermWidth())

	for _, vuln := range vulns {
		t.Row(vuln.Detection, vuln.FixedVersion)
	}

	fmt.Println(t)
	return nil
}

func ScanResults(results []models.ScanResultVulnerabilities) error {
	t := ltable.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("CVE", "Name", "Version", "VulnCheck KEV", "CVSS Base", "CVSS Temporal", "Fixed").
		Width(TermWidth())

	for _, result := range results {
		inKev := lipgloss.NewStyle().Foreground(lipgloss.Color("#34d399")).Render("✔")
		if !result.InKEV {
			inKev = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("✘")
		}

		t.Row(result.CVE, result.Name, result.Version, inKev, result.CVSSBaseScore, result.CVSSTemporalScore, result.FixedVersions)
	}

	fmt.Println(t)
	return nil
}

func SingleColumnResults(results []string, title string) error {
	t := ltable.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers(title).
		Width(TermWidth())

	for _, result := range results {
		t.Row(result)
	}

	fmt.Println(t)
	return nil
}
