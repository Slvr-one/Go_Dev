package scripts

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Schedule struct {
	Name string

	// Day string
	// Time   time.Time
	Date string

	Room     string
	Location string

	Description string
	Subject     string
}

func writeToFile(filename string, schedules []Schedule) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, schedule := range schedules {
		err := writer.Write([]string{schedule.Date, schedule.Name, schedule.Location, schedule.Description})
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	writer.Flush()
}
