package gitModel

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-git/go-git/v5"
)

type GitModel struct {
    current       string
    repo          *git.Repository
    Active        bool
    Height, Width int
}

func InitialGitModel(name string) (GitModel, error) {
    repo, err := git.PlainOpen(name)

    if err != nil {
        fmt.Println("Error when making git repo")
        return GitModel{}, err
    }

    return GitModel{
        repo: repo,
        Active: false,
    }, nil
}

func (g GitModel) Init() tea.Cmd {
    return nil
}

func (g GitModel) View() string {
    wrap := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder(), true, true, true, true).
        BorderForeground(lipgloss.Color("#474747")).
        Width(g.Width).
        Height(g.Height)

    if g.Active {
        wrap = wrap.BorderForeground(lipgloss.Color("#00dd00"))
    }

    s := fmt.Sprintf("%+v\n\n%+v", *g.repo, g.repo.Storer)

    return wrap.Render(s)
}

func (g GitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return g, nil
}
