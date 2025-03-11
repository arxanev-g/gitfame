package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/arxanev-g/gitfame/gitfame/internal/statistics"
)

func PrintTabular(stats *statistics.Stats) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	_, err := fmt.Fprintln(w, "Name\tLines\tCommits\tFiles")

	if err != nil {
		fmt.Println(err)
	}
	for _, line := range stats.Statistics {
		fmt.Fprintln(w, line.AuthorName+"\t"+strconv.Itoa(line.NumLines)+"\t"+strconv.Itoa(line.NumCommits)+"\t"+strconv.Itoa(line.NumFile))
	}

	w.Flush()
}

func PrintCSV(stats *statistics.Stats) {
	header := []string{"Name", "Lines", "Commits", "Files"}
	w := csv.NewWriter(os.Stdout)
	var buff [][]string
	buff = append(buff, header)
	for _, line := range stats.Statistics {
		buff = append(buff, []string{line.AuthorName, strconv.Itoa(line.NumLines), strconv.Itoa(line.NumCommits), strconv.Itoa(line.NumFile)})
	}
	w.WriteAll(buff)
}

func PrintJSON(stats *statistics.Stats) {
	var buff []map[string]interface{}
	for _, line := range stats.Statistics {

		buff = append(buff, map[string]interface{}{
			"name":    line.AuthorName,
			"lines":   line.NumLines,
			"commits": line.NumCommits,
			"files":   line.NumFile,
		})
	}
	jsonData, _ := json.Marshal(buff)

	fmt.Println(string(jsonData))
}

func PrintJSONLines(stats *statistics.Stats) {
	for _, line := range stats.Statistics {

		jsonLine, _ := json.Marshal(map[string]interface{}{
			"name":    line.AuthorName,
			"lines":   line.NumLines,
			"commits": line.NumCommits,
			"files":   line.NumFile,
		})

		fmt.Println(string(jsonLine))
	}
}
