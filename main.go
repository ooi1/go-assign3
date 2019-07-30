// function: go run main.go myfile.txt

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	myArgs := os.Args

	file, err := os.Open(myArgs[1])

	if err != nil {
		fmt.Println("Its fatal!")
		os.Exit(1)
	}

	// data := make([]byte, 999)
	// count, err := file.Read(data)

	// if err != nil {
	// 	fmt.Println("Its fatal!")
	// }

	// fmt.Printf("count: %d, data: %q\n", count, data[:count])
	io.Copy(os.Stdout, file)
}
