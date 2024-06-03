package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handing in golang")

	content := "This is me, pranit patil"

	file, err := os.Create("./newfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)
	file.Close()

}
