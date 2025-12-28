package examples

import (
	"fmt"
	"io/fs"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// PrintFiles recursively lists files from any filesystem using the fs.FS interface.
func PrintFiles(fsys fs.FS) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
}

// LoadConfig reads a file using fs.ReadFile with any filesystem implementation.
func LoadConfig(fsys fs.FS, name string) ([]byte, error) {
	return fs.ReadFile(fsys, name)
}

// GetFileInfo gets file info from any filesystem implementation.
func GetFileInfo(fsys fs.FS, name string) (fs.FileInfo, error) {
	f, err := fsys.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f.Stat()
}

// ListDetails lists directory entries with information from any filesystem.
func ListDetails(fsys fs.FS) error {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return err
	}

	for _, d := range entries {
		info, _ := d.Info()
		fmt.Printf("%s: %d bytes (is dir: %v)\n", d.Name(), info.Size(), d.IsDir())
	}
	return nil
}
