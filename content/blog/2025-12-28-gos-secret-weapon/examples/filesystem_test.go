package examples

import (
	"fmt"
	"testing/fstest"
)

// ExampleLoadConfig demonstrates accepting fs.FS
func ExampleLoadConfig() {
	fsys := fstest.MapFS{
		"config.json": &fstest.MapFile{Data: []byte(`{"env": "prod"}`)},
	}
	data, _ := LoadConfig(fsys, "config.json")
	fmt.Println(string(data))
	// Output: {"env": "prod"}
}

// ExampleGetFileInfo demonstrates accepting fs.FS with file info
func ExampleGetFileInfo() {
	fsys := fstest.MapFS{
		"config.json": &fstest.MapFile{Data: []byte(`{"env": "prod"}`)},
	}
	info, _ := GetFileInfo(fsys, "config.json")
	fmt.Println("File size:", info.Size())
	// Output: File size: 15
}
