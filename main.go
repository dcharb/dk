package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type mode int

const (
	NORMAL mode = iota
	INSERT
)

var edMode mode = NORMAL

func parse(c byte) bool {
	if c == 'c' {
		return true
	}
	return false
}

func main() {
	// Set console/terminal to raw mode
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	// Get input and send to parser
	for {
		c := make([]byte, 1)
		_, err = os.Stdin.Read(c)
		if err != nil {
			panic(err)
		}

		if !parse(c[0]) {
			break
		}
	}
}
