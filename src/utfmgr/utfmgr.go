// Package utfmgr
// Created by RTT.
// Author: teocci@yandex.com on 2021-Oct-22
package utfmgr

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// EncodingHint indicates the file's encoding if there is no BOM.
type EncodingHint int

const (
	// UTF8 indicates the specified encoding.
	UTF8 EncodingHint = iota
	// UTF16LE indicates the specified encoding.
	UTF16LE
	// UTF16BE indicates the specified encoding.
	UTF16BE
	// WINDOWS indicates that the file came from a MS-Windows system
	WINDOWS = UTF16LE
	// POSIX indicates that the file came from Unix or Unix-like systems
	POSIX = UTF8
	// HTML5 indicates that the file came from the web
	// This is recommended by the W3C for use in HTML 5:
	// "For compatibility with deployed content, the byte order
	// mark (also known as BOM) is considered more authoritative
	// than anything else." http://www.w3.org/TR/encoding/#specification-hooks
	HTML5 = UTF8
)

// OpenFile is the equivalent of os.Open().
func OpenFile(name string, d EncodingHint) (UTFReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	rc := readCloser{file: f}

	return NewReader(rc, d), nil
}

// ReadFile is the equivalent of ioutil.ReadFile()
func ReadFile(name string, d EncodingHint) ([]byte, error) {
	file, err := OpenFile(name, d)
	if err != nil {
		return nil, err
	}
	defer CloseFile()(file)

	return ioutil.ReadAll(file)
}

// BytesReader is a convenience function that takes a []byte and decodes them to UTF-8.
func BytesReader(b []byte, d EncodingHint) io.Reader {
	return NewReader(bytes.NewReader(b), d)
}

func CloseFile() func(f UTFReadCloser) {
	return func(f UTFReadCloser) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
