package main

import (
	"fmt"
	"os"
	"time"

	"example.com/emailreports/reports/shared/utils"

	excelreports "example.com/emailreports"
)

var reportList = []string{
	"Race Wise",
}

func main() {
	args := os.Args[1:]
	fmt.Println()
	argParser(args)
	// args := flag.String("run", "Race Wise", "Runs the provided report")
	// flag.Parse()
	// runReport(*args)
}

var commands = map[string]cmd{
	"-run":  cmd{desc: "Runs the provided report", usage: "-run \"report name\"", command: "-run"},
	"-list": cmd{desc: "Lists all the available reports", usage: "-list", command: "list"},
	"-help": cmd{desc: "Lists all the commands available with description", usage: "-help", command: "-help"},
}

func runReport(report string) {
	today := time.Now()
	switch report {
	case string(utils.RaceWise):
		excelreports.GenDemoReport()
		fmt.Println("Report Generated.  Time taken: ", time.Since(today))
	}
}

func argParser(args []string) {
	switch args[0] {
	case "-list":
		fmt.Println(getList())
	case "-help":
		fmt.Print(getHelp())
	case "-run":
		if len(args) > 1 {
			runReport(args[1])
		} else {
			fmt.Println("No report name provided to run, please provide a report name or use h for help")
		}
	default:
		fmt.Println("No argument Provided, please use -h for help on commands")
	}
}

type cmd struct {
	usage   string
	desc    string
	command string
}

func getHelp() string {
	help := "Command \t\t Usage \t\t\t\t Description\n\n"
	for k, v := range commands {
		help += k + "  \t\t  " + v.usage + "  \t\t\t\t  " + v.desc + "\n"
	}
	return help
}

func getList() string {
	list := "Report List :\n\n"
	for _, k := range reportList {
		list += k + "\n"
	}
	return list
}
