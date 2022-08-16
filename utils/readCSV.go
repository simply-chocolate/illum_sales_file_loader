package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

// TODO: We need a way to get the data from the FTP. We could probably do this first.

type empData struct {
	Name string
	Age  string
	City string
}

func readCSV() {

	csvFile, err := os.Open("emp.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		emp := empData{
			Name: line[0],
			Age:  line[1],
			City: line[2],
		}
		fmt.Println(emp.Name + " " + emp.Age + " " + emp.City)
	}
}
