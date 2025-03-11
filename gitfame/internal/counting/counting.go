package counting

import (
	"strings"

	"github.com/arxanev-g/gitfame/gitfame/internal/gitcommands"
	"github.com/arxanev-g/gitfame/gitfame/internal/input"
	"github.com/arxanev-g/gitfame/gitfame/internal/statistics"
)

type Counts struct {
	ULines      map[string]int
	UCommits    map[string]map[string]struct{}
	UNumCommits map[string]int
	UFiles      map[string]map[string]struct{}
	UNumFiles   map[string]int
}

func Count(files []string, commands *input.CommandLineArgs) *statistics.Stats {
	counts := &Counts{
		ULines:      make(map[string]int),
		UCommits:    make(map[string]map[string]struct{}),
		UNumCommits: make(map[string]int),
		UFiles:      make(map[string]map[string]struct{}),
		UNumFiles:   make(map[string]int),
	}

	for _, path := range files {
		ProcessFile(path, commands, counts)
	}

	for author, commits := range counts.UCommits {
		counts.UNumCommits[author] += len(commits)
	}
	for author, files := range counts.UFiles {
		counts.UNumFiles[author] += len(files)
	}
	var stats statistics.Stats
	for name, numCommits := range counts.UNumCommits {
		numLines := 0

		if actualNumLines, ok := counts.ULines[name]; ok {
			numLines = actualNumLines
		}
		stat := statistics.Stat{
			AuthorName: name,
			NumLines:   numLines,
			NumCommits: numCommits,
			NumFile:    counts.UNumFiles[name],
		}
		stats.Statistics = append(stats.Statistics, stat)
	}
	return &stats

}

func ProcessFile(path string, commands *input.CommandLineArgs, counts *Counts) {
	statLines, _ := gitcommands.GitBlame(commands.CommitPointer, path, commands.Repository)

	author := ""
	commit := ""

	if len(statLines) == 1 {
		logLines, _ := gitcommands.GitLog(commands.CommitPointer, path, commands.Repository)

		words := strings.Split(logLines[1], " ")
		author = strings.Join(words[1:len(words)-1], " ")
		if _, ok := counts.UCommits[author]; !ok {
			counts.UCommits[author] = make(map[string]struct{})
			counts.UNumCommits[author]++
		}
		counts.UNumFiles[author]++
	}

	for i := 0; i < len(statLines); i++ {
		words := strings.Split(statLines[i], " ")

		if commands.UseCommiter && words[0] == "committer" {
			commit = strings.Split(statLines[i-5], " ")[0]
		} else if !commands.UseCommiter && words[0] == "author" {
			commit = strings.Split(statLines[i-1], " ")[0]
		} else {
			continue
		}
		author = strings.Join(words[1:], " ")
		counts.ULines[author]++

		if _, ok := counts.UCommits[author]; !ok {
			counts.UCommits[author] = make(map[string]struct{})
		}
		if _, ok := counts.UCommits[author][commit]; !ok {
			counts.UCommits[author][commit] = struct{}{}
		}

		if _, ok := counts.UFiles[author]; !ok {
			counts.UFiles[author] = make(map[string]struct{})
		}
		if _, ok := counts.UFiles[author][path]; !ok {
			counts.UFiles[author][path] = struct{}{}
		}
	}
}
