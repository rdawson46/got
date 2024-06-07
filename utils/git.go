package utils

import (
	"errors"
	"io/fs"
	"os"
)

type Gitter struct {
    Entries   []fs.DirEntry
    ParentDir string
}

func (l *Gitter) IsGit() bool {
    return false
}

func NewGit(dir string) (*Gitter, error) {
    // open parent dir
    b, err := IsDir(dir)

    if err != nil {
        return nil, err
    }

    if !b {
        return nil, errors.New("Not a directory")
    }

    entries, err := os.ReadDir(dir)

    if err != nil {
        return nil, err
    }

    return &Gitter{
        Entries: entries,
        ParentDir: dir,
    }, nil
}
