// Package utfmgr
// Created by RTT.
// Author: teocci@yandex.com on 2021-Oct-22
package utfmgr

import (
	"io"
	"os"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// UTFReadCloser describes a ReadCloser structure for this package.
type UTFReadCloser interface {
	Read(p []byte) (n int, err error)
	Close() error
}

// ReadCloser is a read-closer for this package.
type readCloser struct {
	file   *os.File
	reader io.Reader
}

// Read implements the standard Reader interface.
func (u readCloser) Read(p []byte) (n int, err error) {
	return u.reader.Read(p)
}

// Close implements the standard Closer interface.
func (u readCloser) Close() error {
	if u.file != nil {
		return u.file.Close()
	}
	return nil
}

// NewReader wraps a Reader to decode Unicode to UTF-8 as it reads.
func NewReader(r io.Reader, d EncodingHint) UTFReadCloser {
	var decoder *encoding.Decoder
	switch d {
	case UTF8:
		// Make a transformer that assumes UTF-8 but abides by the BOM.
		decoder = unicode.UTF8.NewDecoder()
	case UTF16LE:
		// Make a transformer that decodes MS-Windows (16LE) UTF files:
		winUTF := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
		// Make a transformer that is like winUTF, but abides by BOM if found:
		decoder = winUTF.NewDecoder()
	case UTF16BE:
		// Make a transformer that decodes UTF-16BE files:
		utf16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
		// Make a transformer that is like utf16be, but abides by BOM if found:
		decoder = utf16be.NewDecoder()
	}

	// Make a Reader that uses utf16bom:
	if rc, ok := r.(readCloser); ok {
		rc.reader = transform.NewReader(rc.file, unicode.BOMOverride(decoder))
		return rc
	}

	return readCloser{
		reader: transform.NewReader(r, unicode.BOMOverride(decoder)),
	}
}
