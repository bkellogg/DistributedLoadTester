package btp

import (
	"os"
	"path/filepath"
)

// currentDir returns the directory that the application
// is being exucuted in
func currentDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
