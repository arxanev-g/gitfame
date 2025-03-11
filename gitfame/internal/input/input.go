package input

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type CommandLineArgs struct {
	Repository    string
	CommitPointer string
	SortOrderKey  string
	UseCommiter   bool
	Format        string
	Extensions    []string
	Languages     []string
	Exclude       []string
	Restricted    []string
}

func NewCommandLineArgs() *CommandLineArgs {
	return &CommandLineArgs{}
}

func isValidOrderKey(orderKey string) bool {
	if (strings.Compare(orderKey, "lines") == 0) || (strings.Compare(orderKey, "commits") == 0) || (strings.Compare(orderKey, "files") == 0) {
		return true
	}
	return false
}

func isValidFormat(Format string) bool {
	if (strings.Compare(Format, "tabular") == 0) || (strings.Compare(Format, "csv") == 0) || (strings.Compare(Format, "json") == 0) || (strings.Compare(Format, "json-lines") == 0) {
		return true
	}
	return false
}

func (cla *CommandLineArgs) GetCommandLineArgs() error {

	args := os.Args[1:]

	var rootCmd = &cobra.Command{
		Use: "gitfame",
		Run: func(cmd *cobra.Command, args []string) {
			cla.Repository, _ = cmd.Flags().GetString("repository")
			cla.CommitPointer, _ = cmd.Flags().GetString("revision")
			cla.SortOrderKey, _ = cmd.Flags().GetString("order-by")
			cla.UseCommiter, _ = cmd.Flags().GetBool("use-committer")
			cla.Format, _ = cmd.Flags().GetString("format")
			cla.Extensions, _ = cmd.Flags().GetStringSlice("extensions")
			cla.Languages, _ = cmd.Flags().GetStringSlice("languages")
			cla.Exclude, _ = cmd.Flags().GetStringSlice("exclude")
			cla.Restricted, _ = cmd.Flags().GetStringSlice("restrict-to")
		},
	}

	rootCmd.Flags().StringP("repository", "r", ".", "Path to Git repository")
	rootCmd.Flags().StringP("revision", "", "HEAD", "Git revision")
	rootCmd.Flags().StringP("order-by", "", "lines", "Sort results by 'lines', 'commits', or 'files'")
	rootCmd.Flags().BoolP("use-committer", "", false, "Use committer instead of author in calculations")
	rootCmd.Flags().StringP("format", "", "tabular", "Output format: 'tabular', 'csv', 'json', 'json-lines'")
	rootCmd.Flags().StringSliceVar(&cla.Extensions, "extensions", []string{}, "extensions")
	rootCmd.Flags().StringSliceVar(&cla.Languages, "languages", []string{}, "languages")
	rootCmd.Flags().StringSliceVar(&cla.Exclude, "exclude", []string{}, "exclude")
	rootCmd.Flags().StringSliceVar(&cla.Restricted, "restrict-to", []string{}, "restrict-to")

	rootCmd.SetArgs(args)
	rootCmd.Execute()

	if !isValidOrderKey(cla.SortOrderKey) {
		return fmt.Errorf("invalid sort order key: %s. Should be one of: 'lines', 'commits', 'files'", cla.SortOrderKey)
	}
	if !isValidFormat(cla.Format) {
		return fmt.Errorf("invalid output format: %s. Should be one of: 'tabular', 'csv', 'json', 'json-lines'", cla.Format)
	}
	return nil
}
