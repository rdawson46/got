package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rdawson46/got/ui/fileModel"
	"github.com/rdawson46/got/ui/gitModel"
	"github.com/rdawson46/got/utils"
)

type Model struct {
    width  int
    height int
    dir    string
    gitter *utils.Gitter
    gitM   gitModel.GitModel
    fileM  fileModel.FileModel
}

func (m Model) View() string {
    s := fmt.Sprintf("Parent: %s\n", m.gitter.ParentDir)

    for _, entry := range m.gitter.Entries {
        s = fmt.Sprintf("%s%s\n", s, entry.Name())
    }

    s = fmt.Sprintf("%s\nIs Git Dir: %t\n", s, m.gitter.IsGit())
    s = fmt.Sprintf("%s\nHeight: %d\n", s, m.height)
    s = fmt.Sprintf("%s\nWidth: %d\n", s, m.width)

    return s
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        m.height = msg.Height
        m.width = msg.Width
        return m, nil
    }
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

    name, err := utils.ToAbs(name)

    if err != nil {
        fmt.Println("Broke when converting")
        os.Exit(1)
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

