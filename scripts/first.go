package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	// "os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Schedule struct {
	Date time.Time
	Name string
	Room string
}

type EnglishLocale struct {
	CurrencySymbol string
	DateFormat     string
}

func BuildArrayFromString(strlist string, seperator string) [][]string {
	result := strings.Split(strlist, seperator)
	var arr [][]string
	for _, r := range result {
		arr = append(arr, strings.Split(r, ";"))
	}
	return arr
}

// func PrintCommandOutput(args ...string) {
// 	cmd := exec.Command("echo")
// 	cmd.Args = args
// 	bytesOut := cmd.CombinedOutput
// 	fmt.Printf("\n%q\n", bytesOut)
// }

// func ParseCSVToSchedule(csvfile string) *[]Schedule {
// 	filename := fmt.Sprintf("%s/%s.csv", csvfile, time.Now().Format("20060102150405"))
// 	// schedules := []Schedule{}
// 	f, _ := os.Open(filename)
// 	defer f.Close()

// }

func checkLine(line string) *string {
	if strings.Contains(line, ",") {
		return &line
	}
	return nil
}

func main() {

	filePath := "file.csv"

	// locale := &EnglishLocale{
	// 	CurrencySymbol: "euro",
	// 	DateFormat:     "",
	// }

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		headerLineReader := checkLine(line)
		fmt.Println(headerLineReader)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	csvReader := csv.NewReader(file)
	// csvReader.Comma = ',' // Use comma as the field delimiter, also default

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return
		}

		// Process the record. can use the locale variable to format the data according to needs.
		fmt.Println(record)
	}
}

//	traverse the specified directory, and for each file or directory in the tree,
//
// it prints whether it's a file or a directory along with its path
func ListFilesInDir(dir string) []string {
	var filenames []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println("File:", path)
			filenames = append(filenames, path)
		}
		// fmt.Println("Directory:", path)
		return nil
	})
	if err != nil {
		// fmt.Println("Error during filepath.Walk:", err)
		panic(fmt.Errorf("error listing files in %s: %v", dir, err))
	}
	return filenames
}
