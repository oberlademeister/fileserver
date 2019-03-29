package indexmaker

import (
	"io"
	"os"
)

// IndexRenderer makes making the HTML output pluggable
type IndexRenderer func(io.Writer, []os.FileInfo) error

// DefaultRenderer renders an index
func DefaultRenderer(w io.Writer, fis []os.FileInfo) error {

	return nil
}
