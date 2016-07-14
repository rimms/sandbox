package main

import (
	"flag" // For command-line parser
	"os"
)

var omitNewLine *bool = flag.Bool("n", false, "Do not print the trailing newline character")

func main() {
	flag.Parse()

	output := ""

	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			output += " "
		}
		output += flag.Arg(i)
	}

	// Add the tailing newline charactor
	if !*omitNewLine {
		output += "\n"
	}

	os.Stdout.WriteString(output)
}
