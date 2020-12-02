package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//取得路徑
	path, _ := os.Getwd()

	//取得bin
	selfbytes, _ := ReadFileBin(path + (os.Args[0])[1:])

	//計算md5sum並輸出
	fmt.Printf("%x\n", md5.Sum(selfbytes))
}

func ReadFileBin(filename string) ([]byte, error) {
	var content []byte

	file, _ := os.Open(filename) // For read access.

	for {
		b := make([]byte, 100)
		count, err := file.Read(b)
		if err != nil {
			if io.EOF == err {
				break
			} else {
				log.Fatal(err)
				return nil, err
			}
		}
		content = append(content, b[:count]...)
	}

	return content, nil
}
