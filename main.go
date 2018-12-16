package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	// Set console/terminal to raw mode
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, oldState)

	// Get input and send to channel
	ch := make(chan [1]byte)
	go func() {
		var c [1]byte
		_, err = os.Stdin.Read(c[:])
		if err != nil {
			panic(err)
		}
		ch <- c
	}()
	out := <-ch
	if string(out[:]) == "a" {
		fmt.Print("test 1")
	} else {
		fmt.Print("test 2")
	}
}
