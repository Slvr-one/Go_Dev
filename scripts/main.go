package scripts

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// type EnglishLocale struct {
// 	CurrencySymbol string
// 	DateFormat     string
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
	// csvReader.Comma = '_' // Comma field delimiter is default

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return
		}

		// Process the record.
		// can use the locale variable to format the data according to needs.
		fmt.Println(record)
	}
}
