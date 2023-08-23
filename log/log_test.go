package log

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"testing"
	"time"
)

func TestLogCount(t *testing.T) {
	file, err := os.Open("/Users/juzhang/Downloads/appemail-puller 4.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Regular expression to match the timestamp and relevant information
	// Modify this regex according to your actual log format
	regex := regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) - INFO - ([^\s]+) - (.*)`)

	// Map to store the count of logs for each timestamp
	logCount := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if len(matches) >= 4 {
			//timestamp := matches[1]
			//module := matches[2]
			//activity := matches[3]

			// Process the timestamp as needed, e.g., extract hours and minutes
			// ...

			// Parse the timestamp into a time.Time object
			timestamp, err := time.Parse("2006/01/02 15:04:05", matches[1])
			if err != nil {
				fmt.Println("Error parsing timestamp:", err)
				continue
			}

			// Truncate timestamp to minute level
			truncatedTimestamp := timestamp.Truncate(time.Minute)

			// Format truncated timestamp for grouping
			formattedTimestamp := truncatedTimestamp.Format("2006/01/02 15:04")

			// Formatted timestamp for grouping
			//formattedTimestamp := timestamp // Modify this part

			logCount[formattedTimestamp]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	keys := make([]string, 0, len(logCount))
	for key := range logCount {
		keys = append(keys, key)
	}

	// Sort keys
	sort.Strings(keys)

	// Access map values using sorted keys
	for _, key := range keys {
		fmt.Printf("Key: %s, Value: %d\n", key, logCount[key])
	}
}
