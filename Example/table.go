package main

import (
	"fmt"
	"os"

	"github.com/rmasci/script"
)

func main() {
	var render string
	if len(os.Args) <= 1 {
		render = "render=grid"
	} else {
		render = fmt.Sprintf("Render=%s", os.Args[1])
	}
	//err := script.File("cities.csv").Fields(",", ",", 9, 10, 6, 5, 4, 3, 2, 1).Table(render)
	err := script.File("cities.csv").Table(render)
	if err != nil {
		fmt.Println(err)
	}
}
