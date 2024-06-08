package utils

import (
    "os"
    "path/filepath"
)

func IsDir(name string) (bool, error) {
    info, err := os.Stat(name)

    if err != nil {
        return false, err
    }

    return info.IsDir(), nil
}

// NOTE: returns path to user home if error occurs when converting
func ToAbs(relative string) (string, error) {
    abs, err := filepath.Abs(relative)

    if err != nil {
        return os.UserHomeDir()
    }

    return abs, nil
}
