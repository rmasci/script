package script

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rmasci/csvtable"
)

func ZFile(path string) *Pipe {
	p := NewPipe()
	f, err := os.Open(path)
	if err != nil {
		return p.WithError(err)
	}

	// Read a small chunk to determine the file type
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return p.WithError(err)
	}

	// Reset the read pointer
	_, err = f.Seek(0, 0)
	if err != nil {
		return p.WithError(err)
	}

	var reader io.Reader
	if http.DetectContentType(buffer) == "application/x-gzip" {
		reader, err = gzip.NewReader(f)
		if err != nil {
			return p.WithError(err)
		}
	} else {
		reader = f
	}

	return p.WithReader(reader)
}

// Fields splits the input string into fields based on the specified input delimiter,
// selects the fields at the specified indices, and then joins these fields with the
// specified output delimiter. The indices are 1-based and negative indices count from
// the end of the fields. If an index is out of range, it is ignored.
//
// For example, Fields("|", ",", 1, 3) applied to the string "field1|field2|field3|field4"
// would return the string "field1,field3".
//
// Fields returns a new Pipe that will write the transformed string when its output is read.
func (p *Pipe) Fields(inDelim, outDelim string, a ...int) *Pipe {
	// if len(a) <= 0 {
	// 	return p
	// }
	return p.FilterScan(func(line string, w io.Writer) {
		var columns []string
		if inDelim == " " {
			columns = strings.Fields(line)
		} else {
			columns = strings.Split(line, inDelim)
		}
		// if user passes nothing to a, use all columns
		if len(a) <= 0 {
			for i := 1; i <= len(columns)+1; i++ {
				a = append(a, i)
			}
		}
		var out strings.Builder
		for _, c := range a {
			for i, col := range columns {
				col = strings.TrimSpace(col)
				if i == c-1 {
					fmt.Fprint(&out, col+outDelim)
				}
			}
		}
		fmt.Fprintf(w, "%s\n", strings.TrimSuffix(out.String(), outDelim))
	})
}

// Table takes CSV from the input pipeline, and creates a formatted table from the provided CSV string.
// The input from the PIPE is expected to be a CSV string, or it may not output at all. You can convert fields to CSV by using the p.Fields function.
// The function accepts an optional list of flags that can be used to customize the table's formatting.
// The flags can be used to specify the table's render style, alignment, header presence, wrapping, indentation, line separation between rows, spacing, columns, and delimiter.
// If the flags are not provided, the table is created with default settings.
// The function returns the formatted table as a string and an error if any occurred during the table creation.
// Try these with <render>-nohead when there is no header. Options are comma separated, "Render=grid","Header=one,two,three"
//
//		+-------------+-------------------------------------------+---------------------------+
//		| Option      | Description                               | Usage                     |
//		+=============+===========================================+===========================+
//		| Render      | How you want the table to be formatted    | Render=See Below          |
//		| Align       | Left Center Right                         | Align=Left                |
//		| NoHeader    | True/False                                | NoHeader=False            |
//		| Wrap        | Wrap Text in cells                        | Wrap=true                 |
//		| Header      | Specify Header                            | Header=One,Two,Three    |
//		+-------------+-------------------------------------------+---------------------------+
//	 Ex: script.Echo("one,two,three").Table("Render=grid","Header=one,two,three", "Align=right", "NoHeader=false", "Wrap=true")
//
// Render Formats:
//
//	+-----------+-------------------------------------+
//	| Render    | Output Format                       |
//	+===========+=====================================+
//	| mysql     | Looks like a MySQL Client Query     |
//	| grid      | Spreadsheet using Graphical Grid    |
//	| gridt     | Spreadsheet using text grid         |
//	| simple    | Simple Table                        |
//	| html      | Output in HTML Table                |
//	| tab       | Just text tab separated             |
//	| csv       | Output in CSV format                |
//	| plain     | Plain Table output                  |
//	+-----------+-------------------------------------+
func (p *Pipe) Table(options ...string) error {
	data, err := p.Bytes()
	if err != nil {
		p.SetError(err)
		return p.Error()
	}
	strData := string(data)
	if !strings.Contains(strData, ",") {
		p.SetError(fmt.Errorf("No Commas in string"))
		return p.Error()
	}
	strData, err = csvtable.Table(strData, options...)
	fmt.Println(strData)
	return nil
}

func Cat(path string) *Pipe {
	var reader io.Reader
	var err error

	f, err := os.Open(path)
	if err != nil {
		return NewPipe().WithError(err)
	}

	// Read the first 512 bytes to determine the MIME type
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return NewPipe().WithError(err)
	}

	// Detect the MIME type
	mimeType := http.DetectContentType(buf[:n])

	// If the MIME type is gzip, create a gzip reader
	if mimeType == "application/x-gzip" {
		reader, err = gzip.NewReader(bytes.NewReader(buf[:n]))
		if err != nil {
			return NewPipe().WithError(err)
		}
	} else {
		// If the MIME type is not gzip, continue to use the file as the reader
		// But first, we need to seek back to the beginning of the file
		_, err = f.Seek(0, 0)
		if err != nil {
			return NewPipe().WithError(err)
		}
		reader = f
	}

	return NewPipe().WithReader(reader)
}
