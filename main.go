package main

import (
	"fmt"
	"os"

	"github.com/rdawson46/got/utils"
)

func main() {
    var name string
    if len(os.Args) < 2 {
        name = "."
    } else {
        name = os.Args[1]
    }

    fmt.Println("Name: %s", name)

    gitter, err := utils.NewGit(name)

    if err != nil {
        fmt.Println("gitter broke")
        return
    }

    fmt.Println(gitter.Entries)
}
