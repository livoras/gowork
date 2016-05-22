package main

import (
	"io"
	"os"
	"strings"
	// "fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i := range b[:n] {
		pos := b[i]
		if pos < 65 || pos > 122 || (pos > 90 && pos < 97) {

		} else {
			if pos > 97 {
				pos = pos + 13
				if (pos > 122) {
					pos = pos - 26
				}
			} else {
				pos = pos + 13
				if (pos > 90) {
					pos = pos - 26
				}
			}
		}
		b[i] = pos
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
