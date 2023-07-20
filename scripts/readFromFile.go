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

func readFromFile(filename string) ([]Schedule, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	schedules, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	var schedulesList []Schedule
	for _, line := range schedules {
		if len(line) == 4 {
			schedulesList = append(schedulesList, Schedule{
				Name:        line[0],
				Date:        line[1],
				Room:        line[2],
				Location:    line[3],
				Description: line[4],
				Subject:     line[5],
			})
		}
	}
	return schedulesList, nil
}
