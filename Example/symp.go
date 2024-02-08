package main

import (
	"fmt"

	"github.com/rmasci/script"
)

func main() {
	pipe := script.Exec(`curl -s https://bespinwerks.web.att.com/todos/todo.json`)
	isTrue, _ := pipe.Exec(`jq '.[] | select(.date | contains("2023"))'`).Exec(`grep true`).Exec(`wc -l`).String()
	pipe = script.Exec(`curl -s https://bespinwerks.web.att.com/todos/todo.json`)
	isFalse, _ := pipe.Exec(`jq '.[] | select(.date | contains("2023"))'`).Exec(`grep false`).Exec(`wc -l`).String()
	fmt.Println("True:", isTrue)
	fmt.Println("False:", isFalse)
}
