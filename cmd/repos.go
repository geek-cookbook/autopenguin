package cmd

import (
	"fmt"
	"os"

	"github.com/funkypenguins-geek-cookbook/penguin/pkg/repo"
	"github.com/spf13/cobra"
)

var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Get a list of repositories on the organization",
	Long:  "Preview which repos will get updated, before writing to them",
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := repo.GetRepos(organization, nil)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		for name, url := range repos {
			fmt.Printf("%s: %s\n", au.Bold(au.Yellow(name)), au.Green(url))
		}

	},
}

func init() {
	rootCmd.AddCommand(reposCmd)
}
