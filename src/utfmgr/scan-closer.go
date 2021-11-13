// Package utfmgr
// Created by RTT.
// Author: teocci@yandex.com on 2021-Oct-22
package utfmgr

import "bufio"

// UTFScanCloser describes a new utfmgr ScanCloser structure.
// It's similar to ReadCloser, but with a scanner instead of a reader.
type UTFScanCloser interface {
	Buffer(buf []byte, max int)
	Bytes() []byte
	Err() error
	Scan() bool
	Split(split bufio.SplitFunc)
	Text() string
	Close() error
}

type scanCloser struct {
	file    UTFReadCloser
	scanner *bufio.Scanner
}

// Buffer will run the Buffer function on the underlying bufio.Scanner.
func (sc scanCloser) Buffer(buf []byte, max int) {
	sc.scanner.Buffer(buf, max)
}

// Bytes will run the Bytes function on the underlying bufio.Scanner.
func (sc scanCloser) Bytes() []byte {
	return sc.scanner.Bytes()
}

// Err will run the Err function on the underlying bufio.Scanner.
func (sc scanCloser) Err() error {
	return sc.scanner.Err()
}

// Scan will run the Scan function on the underlying bufio.Scanner.
func (sc scanCloser) Scan() bool {
	return sc.scanner.Scan()
}

// Split will run the Split function on the underlying bufio.Scanner.
func (sc scanCloser) Split(split bufio.SplitFunc) {
	sc.scanner.Split(split)
}

// Text will return the text from the underlying bufio.Scanner.
func (sc scanCloser) Text() string {
	return sc.scanner.Text()
}

// Close will close the underlying file handle.
func (sc scanCloser) Close() error {
	return sc.file.Close()
}

// NewScanner is a convenience function that takes a filename and returns a scanner.
func NewScanner(name string, d EncodingHint) (UTFScanCloser, error) {
	f, err := OpenFile(name, d)
	if err != nil {
		return nil, err
	}

	return scanCloser{
		scanner: bufio.NewScanner(f),
		file:    f,
	}, nil
}
