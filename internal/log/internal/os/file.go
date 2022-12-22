package os

import "os"

// File embeds a pointer to an os.File struct
type File struct {
	*os.File
}

// MustClose ¯\_(ツ)_/¯
func (f File) MustClose() {
	if e := f.File.Close(); e != nil {
		panic(e)
	}
}
