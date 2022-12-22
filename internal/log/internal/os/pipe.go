package os

import "os"

// MustPipe ¯\_(ツ)_/¯
func MustPipe() (reader File, writer File) {
	if r, w, e := os.Pipe(); e != nil {
		panic(e)
	} else {
		return File{r}, File{w}
	}
}
