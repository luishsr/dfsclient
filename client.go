package main

import (
	"fmt"
	"os"

	"github.com/luishsr/filesys"
)

func main() {
	nodes := []filesys.Node{
		{"localhost", 8000},
	}

	fs := filesys.NewSimpleDistributedFileSystem(nodes)

	// Upload a file
	testFile, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer testFile.Close()

	err = fs.Put(testFile)
	if err != nil {
		fmt.Println("Failed to upload file:", err)
		return
	}
	fmt.Println("File uploaded successfully")

	// Download the file
	downloadedFile, err := fs.Get("test.txt")
	if err != nil {
		fmt.Println("Failed to download file:", err)
		return
	}
	defer downloadedFile.Close()
	defer os.Remove(downloadedFile.Name())

	fmt.Println("File downloaded successfully:", downloadedFile.Name())
}
