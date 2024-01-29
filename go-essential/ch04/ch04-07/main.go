package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	file *os.File
}

func (c *Capper) Write(p []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("empty write")
	}
	str := string(p)
	str = strings.ToUpper(str)
	return c.file.Write([]byte(str))
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
