package main

import (
	"log"

	"github.com/arxanev-g/gitfame/gitfame/internal/counting"
	"github.com/arxanev-g/gitfame/gitfame/internal/getfiles"
	"github.com/arxanev-g/gitfame/gitfame/internal/input"
	"github.com/arxanev-g/gitfame/gitfame/internal/output"
	"github.com/arxanev-g/gitfame/gitfame/internal/statistics"
)

func PrintResults(stats *statistics.Stats, commands *input.CommandLineArgs) {
	stats.SortResults(commands.SortOrderKey)
	switch commands.Format {
	case "csv":
		output.PrintCSV(stats)
	case "json":
		output.PrintJSON(stats)
	case "json-lines":
		output.PrintJSONLines(stats)
	default:
		output.PrintTabular(stats)
	}
}

func main() {
	commands := input.NewCommandLineArgs()
	if err := commands.GetCommandLineArgs(); err != nil {
		log.Fatal(err)
	}
	allFile := getfiles.GetAllFiles(commands)
	stats := counting.Count(allFile, commands)
	PrintResults(stats, commands)
}
