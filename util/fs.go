package util

import "path/filepath"

func FileNameWithExt(filename string) bool {
    fileExt := filepath.Ext(filename)
    return fileExt != ""
}
