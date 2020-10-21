package main

import (
	"fmt"
	"os"
	"time"

	excelreports "example.com/emailreports"
	"example.com/emailreports/utils"
	_ "github.com/jessevdk/go-flags"
)

const usage = `
Usage:
	excel-report COMMAND [OPTIONS, ...]
Commands:
	list        List the reports
	generate 		Generates the provided report
Options:
	-r					Report Name
`

func runReport(ctx *utils.Context) {
	report := ctx.Options.ReportName
	today := time.Now()
	switch report {
	case "Race Wise":
		excelreports.GenDemoReport()
		fmt.Println("Report Generated.  Time taken: ", time.Since(today))
	default:
		fmt.Print("No such report exists, Kindly check the report name")
	}
}

func list() {
	list := "Report List :\n\n"
	for _, k := range utils.ReportList {
		list += k + "\n"
	}
	fmt.Println(list)
}

func main() {
	ctx, _ := utils.Init()

	arg := ctx.NextArg()
	if arg == "" {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch arg {
	case "list":
		list()
	case "generate":
		runReport(ctx)
	default:
		fmt.Printf("%s is not a valid command\n", arg)
		fmt.Println(usage)
		os.Exit(1)
	}
}
