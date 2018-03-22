package main

import (
	"fmt"
	"io"
	"os"
)

type alphaReader string

func (a alphaReader) Read(p []byte) (int, error) {
	count := 0
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		if (a[i] >= 'A' && a[i] <= 'Z') || (a[i] >= 'a' && a[i] <= 'z') {
			fmt.Println(a[i])
			p[i] = a[i]
		}
		count++
	}
	return count, io.EOF
}

func main() {
	str := alphaReader("Hello5")
	io.Copy(os.Stdout, &str)
	fmt.Println()

}
