package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := OpenCVS("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	defer file.Close()
}

func OpenCVS(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}
	RemoveEmptyLines(file)
	return file, err

}

func RemoveEmptyLines(f *os.File) {
	panic(errors.New("Failed parse"))
}
