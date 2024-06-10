package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rdawson46/got/ui/fileModel"
	"github.com/rdawson46/got/ui/gitModel"
	"github.com/rdawson46/got/utils"
)

type focus int

const (
    fileTee focus = iota
    gitMenu
)

type Model struct {
    width  int
    height int
    focus  focus
    dir    string
    gitM   gitModel.GitModel
    fileM  fileModel.FileModel
}

func (m Model) View() string {
    s := ""

    s = fmt.Sprintf("%s\nHeight: %d\n", s, m.height)
    s = fmt.Sprintf("%s\nWidth: %d\n", s, m.width)

    joiner := lipgloss.JoinHorizontal(
        lipgloss.Center,
        m.fileM.View(),
        m.gitM.View(),
    )

    s = fmt.Sprintf("%s\n%s", s, joiner)

    return joiner
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case "tab":
            // move focus
            m.focus = (m.focus + 1) % 2
        default:
            switch m.focus {
            case fileTee:
                // NOTE: ignoring cmd for now
                f, _ := m.fileM.Update(msg)
                m.fileM = f.(fileModel.FileModel)
            case gitMenu:
                g, _ := m.gitM.Update(msg)
                m.gitM = g.(gitModel.GitModel)
            }
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
        fmt.Println("Broke when converting to absolute path")
        os.Exit(1)
    }


    fileM, err := fileModel.InitialFileModel(name)

    if err != nil {
        fmt.Println("Error with initial file model")
        os.Exit(1)
    }

    return Model{
        dir: name,
        fileM: fileM,
    }
}

