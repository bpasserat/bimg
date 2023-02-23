package bimg

import (
	"io"
	"io/ioutil"
)

// Read reads all the content of the given file path
// and returns it as byte buffer.
func Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Write writes the given byte buffer into disk
// to the given file path.
func Write(path string, buf []byte) error {
	return ioutil.WriteFile(path, buf, 0644)
}

func ReadFromReader(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}

// Write writes the given byte buffer into disk
// to the given file path.
func WriteToWriter(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	return err
}
