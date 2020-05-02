package cmd

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "penguin",
	Short: "Generate pages for repositories with ease",
}

var au aurora.Aurora

var organization string
var colors bool
var repoSaveDir string

func init() {
	rootCmd.PersistentFlags().StringVarP(&organization, "organization", "o", "funkypenguins-geek-cookbook", "Organization to scan through")
	rootCmd.PersistentFlags().StringVarP(&repoSaveDir, "repoSaveDir", "s", "repos", "Directory to save repos to when modifying")
	rootCmd.PersistentFlags().BoolVarP(&colors, "colors", "c", true, "Enable color in terminal output")

	au = aurora.NewAurora(colors)
}

//Execute executes the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(au.Red(err))
		os.Exit(1)
	}
}
