package gitModel

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var highlight = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))

func (g GitModel) status() (string, error) {
    tree, err := g.repo.Worktree()
    
    if err != nil {
        fmt.Println("No work tree")
        return "", err
    }

    status, err := tree.Status()

    if err != nil {
        fmt.Println("status broke")
        return "", err
    }

    s := ""

    if len(status) == 0 {
        s = "Working branch is up to date"
    } else{
        for k, v := range status {
            sign := highlight.Render(string(v.Worktree))
            s = fmt.Sprintf("%s%s %s\n", s, k, sign)
        }
    }

    return s, nil
}
