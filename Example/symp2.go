package main

import (
	"fmt"

	"github.com/rmasci/script"
)

func main() {
	pipe := script.Exec(`curl -s https://bespinwerks.web.att.com/todos/todo.json`)
	pipe.Exec(`jq '.[] | select(.date | contains("2023"))'`)
	pipe.Exec(`grep true`)
	isTrue, _ := pipe.Exec(`wc -l`).String()
	pipe = script.Exec(`curl -s https://bespinwerks.web.att.com/todos/todo.json`)
	fmt.Println("True:", isTrue)
	pipe.Exec(`jq '.[] | select(.date | contains("2023"))'`)
	pipe.Exec(`grep false`)
	isFalse, _ := pipe.Exec(`wc -l`).String()
	fmt.Println("False:", isFalse)
}
