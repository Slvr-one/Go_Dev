package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	// test files
	fileOne, fileTwo, fileThree, root = "words.txt", "data.txt", "/tmp/dat", "/home/dvir/Documents"

	// source & destination for copy code
	src  = fileOne
	dest = fileTwo

	files []string

	// test slice of strings
	words = []string{"sky", "falcon", "rock", "hawk"}

	// a string from which we create a slice of bytes
	val  = "old\nfalcon\nsky\ncup\nforest\n"
	data = []byte(val)
)

func main() {

	// Walk - walks the file tree, calling the specified function for each file or directory in the tree including root
	// recursively walking all subdirectories.
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return nil
		}

		// append the file to files slice if isn't a dir & has a .txt ext
		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}
		return nil
	})
	check(err)

	// go over files slice & print out all matching files
	for _, file := range files {
		fmt.Println(file)
	}

	// Stat - Check if file exist, get file info
	fInfo, statErr := os.Stat(fileOne)

	// can check(err), but err type is known here
	if errors.Is(statErr, os.ErrNotExist) {
		fmt.Println("file does not exist")
	} else {
		fmt.Println("file exists")
	}

	// get the size of the file in bytes
	fsize := fInfo.Size()
	fmt.Printf("The file size is %d bytes\n", fsize)

	// get the last modification time
	mTime := fInfo.ModTime()
	fmt.Printf("The file last mod time is: %d \n", mTime) // fmt.Println(mTime)

	// Create - creates a file with mode 0666
	file, touchErr := os.Create(fileTwo)
	defer file.Close()
	check(touchErr)

	fmt.Println("file created")

	// Remove - deletes the given file
	deleteErr := os.Remove(fileThree)
	check(deleteErr)
	fmt.Println("file deleted")

	// copies a file
	bytesRead, err := ioutil.ReadFile(src)
	check(err)

	err = ioutil.WriteFile(dest, bytesRead, 0644)
	check(err)

	// ///////////////////////////////////////////////////////////////////////////////

	// ReadFile - slurping a file’s entire contents into memory
	dat, err := os.ReadFile(fileOne)
	check(err)
	fmt.Print(string(dat))

	// reads the whole file in one go; should not be used for large files #
	content, readErr := ioutil.ReadFile(fileOne)
	check(readErr)
	fmt.Println(string(content))

	// WriteFile - write the 'slice of bytes' (data), to the specified file with 0644 permissions
	// If the file does't exist, it creates it; otherwise - truncates it before writing
	writeErr := ioutil.WriteFile(fileTwo, data, 0644)
	check(writeErr)
	fmt.Println("done")

	// write a slice of strings to a file
	for _, word := range words {
		_, err := file.WriteString(word + "\n")
		check(err)
	}
	fmt.Println("done")

	// ////////////////////////////////////////////////////////////////////////////////

	// to append to file, include the os.O_APPEND flag in flags of os.OpenFile function
	f, err := os.OpenFile(fileOne, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer f.Close()

	if _, err := f.WriteString("cloud\n"); err != nil {
		log.Fatal(err)
	}

	// // Open - Opens the specified file to obtain an os.File value, for reading
	// // the associated file descriptor has mode O_RDONLY
	// f, err := os.Open(fileOne)
	// check(err)
	// defer f.Close()

	// Read some bytes from the beginning of the file. Allow up to 5 to be read but also note how many actually were read
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	//Seek - check a known location in the file and Read there
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// ReadAtLeast - more robustly implemented
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// rewind
	_, err = f.Seek(0, 0)
	check(err)

	// /////////////////////////////////////////////////////////////////////////////////

	// return a acanner to read from; line by line, more appropriate for a large file #
	scanner := bufio.NewScanner(f)

	// Scan - advance to the next token
	// get the advancement with the Text function
	// In the default mode, .Scan advances by lines
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	scannerErr := scanner.Err()
	check(scannerErr)

	// implements a buffered reader, may be useful both for:
	// - its efficiency with many small reads
	// - and because of the additional reading methods it provides
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close - (should be scheduled after Opening, with defer)
	f.Close()
}
