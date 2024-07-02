//https://go.dev/tour/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(byte []byte) (n int,err error) {
	n, err = reader.r.Read(byte)
	for i := 0; i < n ; i++ {
		byte[i] = rot13(byte[i])
	}
	return // no need to pass n and err variable in return statement. Compiler is smart 😉
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}


func rot13(b byte) (byte){
		char := b
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			b += 13
		} else if (char > 'N' && char <= 'Z') || (char > 'm' && char <= 'z') {
			b -= 13
		}
		return b
}
