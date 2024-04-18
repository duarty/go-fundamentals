package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//string
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	fileLen, err := f.WriteString("john doe")
	println(fileLen)
	f.Close()

	openedFile, err := os.ReadFile("file.txt")
	fmt.Println(string(openedFile))

	//bytes
	fb, err := os.Create("fileBytes.txt")
	if err != nil {
		panic(err)
	}
	b, err := fb.Write([]byte("asdasdasd"))

	if err != nil {
		panic(err)
	}
	println(b)
	fb.Close()

	//Buffer
	fopen, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fopen)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
