package main

import (
	"github.com/rmasci/script"
)

func main() {
	script.ZFile("../testdata/releases.json.gz").First(10).Stdout()
}
