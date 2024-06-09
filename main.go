package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rdawson46/got/ui"
    // "github.com/go-git/go-git"
)

func main() {
    p := tea.NewProgram(ui.InitializeModel(), tea.WithAltScreen())

    if _, err := p.Run(); err != nil {
        fmt.Println("Broke:", err)
        os.Exit(1)
    }
}
