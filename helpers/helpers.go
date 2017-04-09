package helpers

import (
    "strings"
    "path/filepath"
)

func RemoveFileExt(filename string) string {
    return strings.TrimSuffix(filename, filepath.Ext(filename))
}