package ui

import (
    "os"
    "fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rdawson46/got/utils"
)

type Model struct {
    dir string
    gitter *utils.Gitter
}

func (m Model) View() string {
    return "running"
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, nil
}

func (m Model) Init() tea.Cmd {
    return tea.Batch()
}

func InitializeModel() Model {
    var name string
    if len(os.Args) < 2 {
        name = "."
    } else {
        name = os.Args[1]
    }

    gitter, err := utils.NewGit(name)

    if err != nil {
        fmt.Println("gitter broke")
        os.Exit(1)
    }

    return Model{
        gitter: gitter,
        dir: name,
    }
}

