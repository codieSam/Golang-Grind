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

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()
	// buff := make([]byte, 12)
	// d, err := f.Read(buff)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buff); i++ {
	// 	fmt.Println("Data: ", d, string(buff[i]))
	// }

	// we can simply read the file using ReadFile method

	data, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// Read folders

	dir, err := os.Open("../")

	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)

	for _, fi := range fileInfo {
		fmt.Println(fi.Name(), fi.IsDir())
	}

}
