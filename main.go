package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/user/siftlog/filter"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: siftlog <field=value> [field=value ...]")
		fmt.Fprintln(os.Stderr, "Example: siftlog level=error service=api")
		os.Exit(1)
	}

	queries, err := filter.ParseQueries(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid query: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if filter.Match(line, queries) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}
}
