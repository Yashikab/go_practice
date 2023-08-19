package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {
	switch {
	case ('A' <= c && c <= 'Z'):
		return (c-'A'+13)%26 + 'A'
	case ('a' <= c && c <= 'z'):
		return (c-'a'+13)%26 + 'a'
	default:
		return c
	}
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
