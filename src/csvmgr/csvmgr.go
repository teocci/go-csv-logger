// Package csvmgr
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-30
package csvmgr

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/teocci/go-csv-logger/src/utfmgr"
)

const lineBreak = '\n'

func LineCounter(fn string) (count int) {
	f := OpenFile(fn)
	defer CloseFile()(f)

	buf := make([]byte, bufio.MaxScanTokenSize)
	lineSep := []byte{lineBreak}

	for {
		c, err := f.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			log.Fatal(err)
		}
	}
}

func OpenUTFFile(fn string) utfmgr.UTFReadCloser {
	f, err := utfmgr.OpenFile(fn, utfmgr.UTF8)
	HasError(err, true)

	return f
}

func UTFBufferFile(fn string) []byte {
	f, err := utfmgr.OpenFile(fn, utfmgr.UTF8)
	HasError(err, true)
	defer utfmgr.CloseFile()(f)

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(f); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func BufferFile(fn string) []byte {
	f, err := os.Open(fn)
	HasError(err, true)
	defer CloseFile()(f)

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(f); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func OpenReader(fn string) (*csv.Reader, *os.File) {
	f := OpenFile(fn)

	// Parse the file
	return csv.NewReader(f), f
}

func OpenFile(fn string) *os.File {
	f, err := os.Open(fn)
	HasError(err, true)

	return f
}

func CreateFile(fn string) *os.File {
	w, err := os.Create(fn)
	HasError(err, true)

	return w
}

func CloseFile() func(f *os.File) {
	return func(f *os.File) {
		err := f.Close()
		HasError(err, false)
	}
}

func FlushWriter() func(w *bufio.Writer) {
	return func(w *bufio.Writer) {
		err := w.Flush()
		HasError(err, false)
	}
}