package main

import (
	"bufio"
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

	// data, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// // Read folders

	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// Create a file

	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Hi Go. ")
	f.WriteString("Nice language")

	bytes := []byte("Hello Golang")

	f.Write(bytes)

	// Read and write to another file (Streaming fashion)

	sourceFIle, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	defer sourceFIle.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFIle)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		e := writer.WriteByte(b)
		if err != nil {
			panic(e)
		}

	}
	writer.Flush()
	fmt.Println("Written to new file successfully")

	// Delete a file

	er := os.Remove("example2.txt")
	if er != nil {
		panic(er)
	}

	fmt.Println("File deleted successfully")

}
