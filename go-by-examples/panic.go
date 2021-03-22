package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/tpm/file")

	if err != nil {
		panic(err)
	}
}
