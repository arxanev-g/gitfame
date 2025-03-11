package getfiles

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/arxanev-g/gitfame/gitfame/internal/input"
)

type language struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Extensions []string `json:"extensions"`
}

func GetAllFiles(commands *input.CommandLineArgs) []string {
	cmd := exec.Command("git", "ls-tree", "-r", "--name-only", commands.CommitPointer)
	cmd.Dir = commands.Repository
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Error in ls-tree")
	}
	files := strings.Split(string(output), "\n")

	var languages []language
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			break
		}
		dir = filepath.Dir(dir)
	}

	data, _ := os.ReadFile(dir + "/gitfame/configs/language_extensions.json")
	json.Unmarshal(data, &languages)

	set := make(map[string]int, len(languages))
	for i, j := range languages {
		set[strings.ToLower(j.Name)] = i
	}
	for _, i := range commands.Languages {
		if itr, ok := set[strings.ToLower(i)]; ok {
			commands.Extensions = append(commands.Extensions, languages[itr].Extensions...)
		}
	}
	var fileList []string
	for _, file := range files {
		if file != "" {
			if !HasExtension(file, commands.Extensions) {
				continue
			}
			if len(commands.Exclude) > 0 && MatchesPatterns(file, commands.Exclude) || len(commands.Restricted) > 0 && !MatchesPatterns(file, commands.Restricted) {
				continue
			}
			fileList = append(fileList, file)
		}
	}
	return fileList
}

func HasExtension(path string, excludedExtensions []string) bool {
	if len(excludedExtensions) == 0 {
		return true
	}
	ext := filepath.Ext(path)
	for _, e := range excludedExtensions {
		if strings.EqualFold(ext, e) {
			return true
		}
	}
	return false
}

func IsAcceptableLanguage(fileLanguage string, languages []string) bool {
	if len(languages) == 0 {
		return true
	}
	for _, language := range languages {
		if strings.EqualFold(language, fileLanguage) {
			return true
		}
	}
	return false
}

func GetLanguage(path string, mapping []language) string {
	fileExtension := filepath.Ext(path)
	for _, mappingEntity := range mapping {
		for _, extension := range mappingEntity.Extensions {
			if strings.EqualFold(fileExtension, extension) {
				return mappingEntity.Name
			}
		}
	}
	return ""
}

func MatchesPatterns(filename string, patterns []string) bool {
	for _, pattern := range patterns {
		match, _ := filepath.Match(pattern, filename)
		if match {
			return true
		}
	}
	return false
}
