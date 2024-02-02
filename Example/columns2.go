package main

import (
	"flag"

	"github.com/rmasci/script"
)

// usage: go run columns2.go -f one2ten.csv -d "," -c "1,4,3,5,2,5,3,5,9"
func main() {
	var file string
	var columns string
	var delimeter string
	flag.StringVar(&file, "f", "", "File")
	flag.StringVar(&columns, "c", "1", "Columns - comma separate.")
	flag.StringVar(&delimeter, "d", " ", "Delimeter")
	flag.Parse()
	script.File("one2ten.txt").Column(delimeter, 6, 1, 8, 3).Match("six").Stdout()
}
