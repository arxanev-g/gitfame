package statistics

import (
	"sort"
)

type Stat struct {
	AuthorName string
	NumLines   int
	NumCommits int
	NumFile    int
}

type Stats struct {
	Statistics []Stat
}

func (s *Stats) SortResults(sortKey string) {
	switch sortKey {
	case "lines":
		sort.SliceStable(s.Statistics, func(i, j int) bool {
			if s.Statistics[i].NumLines == s.Statistics[j].NumLines {
				if s.Statistics[i].NumCommits == s.Statistics[j].NumCommits {
					if s.Statistics[i].NumFile == s.Statistics[j].NumFile {
						return s.Statistics[i].AuthorName < s.Statistics[j].AuthorName
					}
					return s.Statistics[i].NumFile > s.Statistics[j].NumFile
				}
				return s.Statistics[i].NumCommits > s.Statistics[j].NumCommits
			}
			return s.Statistics[i].NumLines > s.Statistics[j].NumLines
		})
	case "commits":
		sort.SliceStable(s.Statistics, func(i, j int) bool {
			if s.Statistics[i].NumCommits == s.Statistics[j].NumCommits {
				if s.Statistics[i].NumLines == s.Statistics[j].NumLines {
					if s.Statistics[i].NumFile == s.Statistics[j].NumFile {
						return s.Statistics[i].AuthorName < s.Statistics[j].AuthorName
					}
					return s.Statistics[i].NumFile > s.Statistics[j].NumFile
				}
				return s.Statistics[i].NumLines > s.Statistics[j].NumLines
			}
			return s.Statistics[i].NumCommits > s.Statistics[j].NumCommits
		})
	case "files":
		sort.SliceStable(s.Statistics, func(i, j int) bool {
			if s.Statistics[i].NumFile == s.Statistics[j].NumFile {
				if s.Statistics[i].NumLines == s.Statistics[j].NumLines {
					if s.Statistics[i].NumCommits == s.Statistics[j].NumCommits {
						return s.Statistics[i].AuthorName < s.Statistics[j].AuthorName
					}
					return s.Statistics[i].NumCommits > s.Statistics[j].NumCommits
				}
				return s.Statistics[i].NumLines > s.Statistics[j].NumLines
			}
			return s.Statistics[i].NumFile > s.Statistics[j].NumFile
		})
	}
}
