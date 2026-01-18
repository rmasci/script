package script

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"os"
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
