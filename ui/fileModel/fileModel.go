package fileModel

import (
	"fmt"
	"io/fs"
    "os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
    "github.com/rdawson46/got/utils"
)

// model to represent file tree
type FileModel struct {
    path          string
    Entires       []fs.DirEntry
    current       int
    Active        bool
    Height, Width int
}

// TODO: better error handling
func InitialFileModel(name string) (FileModel, error) {
    path, err := utils.ToAbs(name)
    dir, err := utils.IsDir(path)


    if err != nil {
        fmt.Println("Error with IsDir")
        os.Exit(1)
    }

    if !dir {
        path, err = os.UserHomeDir()

        if err != nil {
            fmt.Println("Could not find home di")
            os.Exit(1)
        }
    }

    entries, err := os.ReadDir(path)

    if err != nil {
        fmt.Println("couldn't read dir")
        os.Exit(1)
    }

    return FileModel{
        path: path,
        Entires: entries,
        current: 0,
        Active: true,
    }, nil
}

func (f FileModel) Init() tea.Cmd {
    return nil
}

func (f FileModel) View() string {
    title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.ANSIColor(44))

    s := title.Render(f.path)

    for i, entry := range f.Entires {
        name := entry.Name()

        if i == f.current {
            highlight := lipgloss.NewStyle().Foreground(lipgloss.Color("#cc0088"))

            name = highlight.Render(name)
        }

        s = fmt.Sprintf("%s\n%s", s, name)
    }

    wrap := lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder(), true, true, true, true).
        BorderForeground(lipgloss.Color("#474747")).
        Width(f.Width).
        Height(f.Height)

    if f.Active {
        wrap = wrap.BorderForeground(lipgloss.Color("#00dd00"))
    }

    return wrap.Render(s)
}

func (f FileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "j":
            f.current = min(f.current + 1, len(f.Entires) - 1)
        case "k":
            f.current = max(f.current - 1, 0)
        }
    }

    return f, nil
}
