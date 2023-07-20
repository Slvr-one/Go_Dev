package scripts

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
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

func ParseCSVToSchedule(csvfile string) (*[]Schedule, error) {

	schedules := []Schedule{}

	filename := fmt.Sprintf("%s/%s.csv", csvfile, time.Now().Format("20060102150405"))
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer f.Close()
	reader := csv.NewReader(f)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return nil, err
		}
		if len(line) == 0 { // for exact finish on last one
			continue
		}
		schedule := Schedule{
			Date:        line[0],
			Name:        line[1],
			Location:    line[2],
			Description: line[3],
		}
		schedules = append(schedules, schedule)
	}
	return &schedules, nil
}
