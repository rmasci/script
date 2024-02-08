package main

import (
	"github.com/rmasci/script"
)

func main() {
	script.File("cities.csv").Fields(",", ",", 10, 9, 8, 1, 2, 3, 4, 5, 6, 7).Table()
}
