package gitModel

import (
	"os"
    "fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
)


func (g GitModel) status() string {
    ref, err := g.repo.Head()

    if err != nil {
        os.Exit(1)
    }

    commit, err := g.repo.CommitObject(ref.Hash())

    if err != nil {
        os.Exit(1)
    }

    tree, err := commit.Tree()

    if err != nil {
        os.Exit(1)
    }

    s := ""

    tree.Files().ForEach(func(f *object.File) error {
        s = fmt.Sprintf("%sblob %s    %s\n", s, f.Hash, f.Name)
        return nil
    })

    return s
}
