package gitcommands

import (
	"os/exec"
	"strings"
)

func GitBlame(commit, pathToFile, execDir string) ([]string, error) {
	gitCommand := exec.Command("git", "blame", "--line-porcelain", "-b", commit, pathToFile)
	gitCommand.Dir = execDir

	out, _ := gitCommand.Output()
	var res strings.Builder
	res.Write(out)

	return strings.Split(res.String(), "\n"), nil
}

func GitLog(commit, pathToFile, execDir string) ([]string, error) {
	gitCommand := exec.Command("git", "log", "-p", commit, "--follow", "--", pathToFile)
	gitCommand.Dir = execDir

	out, _ := gitCommand.Output()
	var res strings.Builder
	res.Write(out)

	return strings.Split(res.String(), "\n"), nil
}
