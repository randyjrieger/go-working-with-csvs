package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func createCoolCarScale(data [][]string) []CoolCarScale {
	// convert csv lines to array of structs
	var coolCarScale []CoolCarScale
	for i, line := range data {
		if i > 0 {
			var rec CoolCarScale
			for j, field := range line {
				if j == 0 {
					rec.Make = field
				} else if j == 1 {
					rec.Model = field
				} else if j == 2 {
					var err error
					rec.Coolness, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				}
			}
			coolCarScale = append(coolCarScale, rec)
		}
	}

	return coolCarScale
}

func main() {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// 2. Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 3. Assign successive lines of raw CSV data to fields of the created structs
	coolCarScale := createCoolCarScale(data)

	// 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonData, err := json.MarshalIndent(coolCarScale, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
