package fileModel

import "io/fs"

// model to represent file tree
type FileModel struct {
    path    string
    Entires []fs.DirEntry
    current int
}
