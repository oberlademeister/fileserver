package indexmaker

import (
	"fmt"
	"io"
	"os"
)

// IndexRenderer makes making the HTML output pluggable
type IndexRenderer func(w io.Writer, dir string, fis []os.FileInfo) error

// GetDefaultRenderer renders an index
func GetDefaultRenderer(config *Config) IndexRenderer {
	return func(w io.Writer, dir string, fis []os.FileInfo) error {
		fmt.Fprintf(w, "Path %s\n", dir)
		return nil
	}
}
