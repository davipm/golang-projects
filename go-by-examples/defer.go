package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	_, _ = fmt.Fprintln(f, "creating")
}

func closeFile(f *os.File) {
	fmt.Println("writing")
	err := f.Close()

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.text")
	defer closeFile(f)
	writeFile(f)
}
