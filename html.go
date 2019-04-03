package indexmaker

import (
	"fmt"
	"io"
	"os"
)

// IndexRenderer makes making the HTML output pluggable
type IndexRenderer func(w io.Writer, dir string, fis []os.FileInfo) error

const _hdr = `<html lang="en"><head><meta charset="utf-8"><link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous"></head>
`
const _ftr = `
</body></html>
`

// GetDefaultRenderer renders an index
func GetDefaultRenderer(config *Config) IndexRenderer {
	return func(w io.Writer, dir string, fis []os.FileInfo) error {
		fmt.Fprint(w, _hdr)
		fmt.Fprintf(w, "<table class=\"table\">")
		fmt.Fprint(w, "<head><th>name</th><th>size</th></head>")
		for _, fi := range fis {
			fmt.Fprint(w, "<tr>\n")
			if fi.IsDir() {
				fmt.Fprintf(w, "<td><a href=\"./%s\">%s</h></td><td>%s</td>\n", fi.Name(), fi.Name(), "(dir)")
			} else {
				fmt.Fprintf(w, "<td><a href=\"./%s\">%s</h></td><td>%d</td>\n", fi.Name(), fi.Name(), fi.Size())
			}
			fmt.Fprint(w, "</tr>\n")
		}
		fmt.Fprintf(w, "</table>")
		fmt.Fprint(w, _ftr)
		return nil
	}
}
