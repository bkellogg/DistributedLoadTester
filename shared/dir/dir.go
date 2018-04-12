package dir

import (
	"os"
	"path/filepath"
)

// CurrentDir returns the directory that the application
// is being exucuted in
func CurrentDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
