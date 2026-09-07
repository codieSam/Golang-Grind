package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Println("File Name: ", fileInfo.Name())
	// fmt.Println("File Size: ", fileInfo.Size())
	// fmt.Println("File Last modifies on: ", fileInfo.ModTime())

	// read file

	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()
	buff := make([]byte, 10)
	d, err := f.Read(buff)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(buff); i++ {
		fmt.Println("Data: ", d, string(buff[i]))
	}

}
