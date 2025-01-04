package main

import (
	"bufio"
	"fmt"
	"os"

	cs2log "github.com/joao-silva1007/cs2-log-re2"
)

// Usage:
//
// From file:
// go run main.go example.log
//
// From STDIN:
// cat example.log | go run main.go
//
// To File:
// go run main.go > out.txt
//
// Omit errors:
// go run main.go 2>/dev/null

func main() {

	var file *os.File
	var err error

	if len(os.Args) < 2 {
		file = os.Stdin
	} else {
		file, err = os.Open(os.Args[1])
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := bufio.NewReader(file)

	// read first line
	l, _, err := r.ReadLine()

	for err == nil {

		// parse
		m, errParse := cs2log.Parse(string(l))

		if errParse != nil {
			// print parse errors to stderr
			fmt.Fprintf(os.Stderr, "ERROR: %s", cs2log.ToJSON(m))
		} else {
			// print to stdout
			fmt.Fprintf(os.Stdout, "%s", cs2log.ToJSON(m))
		}

		// next line
		l, _, err = r.ReadLine()
	}
}
