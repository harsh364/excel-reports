package main

import (
	"flag"
	"fmt"
)

func main() {
	// args := os.Args[1:]
	fmt.Println()
	// argParser(args)
	run := flag.String("run", "", "Runs the provided report")
	list := flag.Bool("list", false, "Lists all the Reports")

	flag.Parse()
	if *list {
		fmt.Println(getList())
	}
	runReport(*run)
}
