package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/rmasci/script"
)

// usage: go run columns2.go -f one2ten.csv -d "," -c "1,4,3,5,2,5,3,5,9"
func main() {
	var file string
	var columns string
	var delimeter string
	flag.StringVar(&file, "f", "one2test.txt", "File")
	flag.StringVar(&columns, "c", "1", "Columns - comma separate.")
	flag.StringVar(&delimeter, "d", " ", "Delimeter")
	flag.Parse()

	// Split the columns string into a slice of strings
	columnsStr := strings.Split(columns, ",")

	// Create a slice to hold the integers
	columnsInt := make([]int, len(columnsStr))

	// Iterate over the columnsStr slice and convert each string to an integer
	for i, colStr := range columnsStr {
		colInt, err := strconv.Atoi(colStr)
		if err != nil {
			panic(err)
		}
		columnsInt[i] = colInt
	}

	// Use the columnsInt slice in the Fields function
	script.Cat(file).Fields(delimeter, delimeter, columnsInt...).Stdout()
}
