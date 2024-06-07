package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rdawson46/got/ui"
)

func main() {
    var name string
    if len(os.Args) < 2 {
        name = "."
    } else {
        name = os.Args[1]
    }

    /*
    gitter, err := utils.NewGit(name)

    if err != nil {
        fmt.Println("gitter broke")
        return
    }

    fmt.Println(gitter.Entries)
    */

    p := tea.NewProgram(ui.InitializeModel(), tea.WithAltScreen())

    if _, err := p.Run(); err != nil {
        fmt.Println("Broke:", err)
        os.Exit(1)
    }
}
