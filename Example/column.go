package main

import (
	"github.com/rmasci/script"
)

func main() {
	// Echo a string and select the second column (1-indexed). The output is "two".
	script.Echo("one two three four five").Column(2).Stdout()

	// Echo a string and select the second field (1-indexed). The output is "two".  Same as it is with Column.
	script.Echo("one two three four five").Fields(" ", " ", 2).Stdout()

	// Echo a string with semicolon-separated fields. Select the second, fifth, and third fields (1-indexed) and join them with commas.
	// The output is "two,five,three,seven,six". In this case, 234 is skipped.
	script.Echo("one;two;three;four;five;six;seven").Fields(";", ",", 2, 5, 3, 234, 7, 6).Stdout()
}
