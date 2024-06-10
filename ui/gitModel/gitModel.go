package gitModel

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/go-git/go-git/v5"
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
    
    return "git model view"
}

func (g GitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return g, nil
}
