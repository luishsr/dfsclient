// Client is a sample implementation to test the
// distributed file system
package main

import (
	"fmt"
	"os"

	"github.com/luishsr/filesys"
)

func main() {

	// List running Nodes
	nodes := []filesys.Node{
		{"localhost", 8000}, {"localhost", 8001},
	}

	// Spins up the distributed file system's component
	fs := filesys.NewSimpleDistributedFileSystem(nodes)

	// Checks the local file before submitting
	testFile, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer testFile.Close()

	// Uploads a file
	err = fs.Put(testFile)
	if err != nil {
		fmt.Println("Failed to upload file:", err)
		return
	}
	fmt.Println("File uploaded successfully")

	// Downloads the file
	downloadedFile, err := fs.Get("test.txt")
	if err != nil {
		fmt.Println("Failed to download file:", err)
		return
	}
	defer downloadedFile.Close()
	defer os.Remove(downloadedFile.Name())

	fmt.Println("File downloaded successfully:", downloadedFile.Name())
}
