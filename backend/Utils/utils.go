package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func LoadCSVData(filename string) ([][]any, error) {
	// Open the CSV filename
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result [][]interface{}
	for index, record := range records {
		if index == 0 {
			continue
		}
		row := make([]any, 0, len(record))
		for _, value := range record {
			row = append(row, value)
		}
		result = append(result, row)
	}

	return result, nil
}

func FormatValues(values []any) string {
	base := "("
	for _, val := range values {
		base = base + ", '" + fmt.Sprint(val) + "'"
	}
	base += ")"
	return strings.Replace(base, ",", " ", 1)
}

func getCurrentTime() string {
	t := time.Now().UTC().Format("2006-01-02 15:04:05.000")
	return t
}

func Log(endp string) {
	log.Printf("Requested : <%s>\n", endp)
}
