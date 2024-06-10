package gitModel

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/go-git/go-git/v5"
    "github.com/charmbracelet/lipgloss"
)

type GitModel struct {
    repo *git.Repository
}

func InitialGitModel() GitModel {
    return GitModel{
        repo: nil,
    }
}

func (g GitModel) Init() tea.Cmd {
    return nil
}

func (g GitModel) View() string {
    wrap := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder(), true, true, true, true).
        BorderForeground(lipgloss.Color("#474747"))


    return wrap.Render("git model view")
}

func (g GitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return g, nil
}
