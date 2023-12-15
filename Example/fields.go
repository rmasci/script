package main

import (
	"github.com/rmasci/script"
)

func main() {
	script.File("cities.csv").Fields(",", ",", 9, 10, 6, 5, 4, 3, 2, 1).Stdout()
}
