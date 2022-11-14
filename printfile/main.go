package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
