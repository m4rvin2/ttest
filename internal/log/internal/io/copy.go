package io

import "io"

// MustCopy ¯\_(ツ)_/¯
func MustCopy(destination io.Writer, source io.Reader) (written int64) {
	if w, e := io.Copy(destination, source); e != nil {
		panic(e)
	} else {
		return w
	}
}
