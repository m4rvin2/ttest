package tabwriter

import (
	"io"
	"text/tabwriter"
)

// New constructs and returns a pointer to a Writer struct with the specified
// output, minimum width tab width, padding, padding character, and flags. For
// more information on these parameters, please refer to the official
// documentation for the standard tabwriter package: https://golang.org/pkg/text/tabwriter/
func New(output io.Writer, minimumWidth, tabWidth, padding int, paddingCharacter byte, flags uint) *Writer {
	return &Writer{
		tabwriter.NewWriter(output, minimumWidth, tabWidth, padding, paddingCharacter, flags),
	}
}

// Writer embeds a pointer to a tabwriter.Writer struct
type Writer struct {
	*tabwriter.Writer
}

// MustFlush ¯\_(ツ)_/¯
func (w *Writer) MustFlush() {
	if e := w.Writer.Flush(); e != nil {
		panic(e)
	}
}
