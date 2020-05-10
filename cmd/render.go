package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/geek-cookbook/autopenguin/pkg/render"
	"github.com/geek-cookbook/autopenguin/pkg/repo"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Updates README files in repositories",
	Long:  "Preview which repos will get updated, before writing to them",
	Run: func(cmd *cobra.Command, args []string) {
		tp := github.BasicAuthTransport{
			Username: ghUsername,
			Password: ghPassword,
		}
		repos, err := repo.GetRepos(organization, tp.Client())
		fatalErrorCheck(err)

		i := 0
		for name, url := range repos {
			fmt.Printf("Cloning Repositories [%d/%d]: %s/%s\n", i, len(repos), au.Bold(au.Yellow(organization)), au.Bold(au.Green(name)))
			directory := path.Join(repoSaveDir, organization, name)
			_, err := git.PlainClone(directory, false, &git.CloneOptions{
				URL: url,
			})

			if err != nil {
				if err != git.ErrRepositoryAlreadyExists {
					fatalErrorCheck(err)
				}
			}

			i++
		}
		fmt.Printf("Cloned Repositories\n")

		skippedRepos := []string{}
		noChanges := []string{}

		i = 0
		for name := range repos {
			fmt.Printf("Rendering READMEs [%d/%d]: %s/%s\n", i, len(repos), au.Bold(au.Yellow(organization)), au.Bold(au.Green(name)))
			i++
			directory := path.Join(repoSaveDir, organization, name)
			readme := path.Join(directory, "README.md")
			repocfg := path.Join(directory, ".funkypenguin", "repo.yaml")

			cfgf, err := os.Open(repocfg)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("%s %s does not exist on %s/%s. Skipping.\n", au.Yellow(au.Bold("WARN")), au.Blue(".funkypenguin/repo.yaml"), au.Yellow(organization), au.Green(name))
					skippedRepos = append(skippedRepos, name)
					continue

				} else {
					fatalErrorCheck(err)
				}
			}
			cfgb, err := ioutil.ReadAll(cfgf)
			fatalErrorCheck(err)
			cfg, err := repo.GetRepoConfig(cfgb)
			fatalErrorCheck(err)

			ctx := render.GetREADMEContext(cfg)
			tpl, err := render.GetREADMETemplate(cfg)
			fatalErrorCheck(err)

			prev, err := ioutil.ReadFile(readme)
			if err != nil {
				if !os.IsNotExist(err) {
					fatalErrorCheck(err)
				}
			}

			file, err := os.Create(readme)
			fatalErrorCheck(err)

			err = tpl.Execute(file, ctx)
			fatalErrorCheck(err)

			err = file.Close()
			fatalErrorCheck(err)

			current, err := ioutil.ReadFile(readme)
			fatalErrorCheck(err)

			if prev == current {
				fmt.Printf("%s %s/%s already has an up-to-date README.\n", au.Blue(au.Bold("INFO")), au.Yellow(organization), au.Green(name))
				noChanges = append(noChanges, name)
			}

		}
		if wetRun {
			i = 0
			for name := range repos {
				fmt.Printf("Pushing Updates [%d/%d]: %s/%s\n", i, len(repos), au.Bold(au.Yellow(organization)), au.Bold(au.Green(name)))
				i++
				skip := false
				for _, skipped := range skippedRepos {
					if skipped == name {
						fmt.Printf("%s %s/%s was previously skipped. Skipping (Again).\n", au.Blue(au.Bold("INFO")), au.Yellow(organization), au.Green(name))
						skip = true
					}
				}
				for _, skipped := range noChanges {
					if skipped == name {
						fmt.Printf("%s %s/%s had no changes. Skipping.\n", au.Blue(au.Bold("INFO")), au.Yellow(organization), au.Green(name))
						skip = true
					}
				}
				if skip {
					continue
				}
				directory := path.Join(repoSaveDir, organization, name)
				repo, err := git.PlainOpen(directory)
				fatalErrorCheck(err)
				wt, err := repo.Worktree()
				fatalErrorCheck(err)

				err = wt.AddGlob("README.md")
				fatalErrorCheck(err)
				_, err = wt.Commit("Update README (via .funkypenguin)", &git.CommitOptions{
					Author: &object.Signature{
						Email: "cookbook@funkypenguin.co.nz",
						Name:  "Penguin Tools",
						When:  time.Now(),
					},
				})
				fatalErrorCheck(err)
				err = repo.Push(&git.PushOptions{
					Auth: &http.BasicAuth{
						Username: ghUsername,
						Password: ghPassword,
					},
				})
				fatalErrorCheck(err)

			}
		} else {
			fmt.Printf("%s Wet Run was not specified, not committing (try with -w to wet run)", au.Bold(au.Blue("INFO")))
		}
	},
}

func fatalErrorCheck(e error) {
	if e != nil {
		notifyErrorCheck(e)
		os.Exit(1)
	}
}

func notifyErrorCheck(e error) {
	if e != nil {
		fmt.Printf("%s %v\n", au.Red(au.Bold("ERROR")), e)
	}
}

var wetRun bool
var ghUsername string
var ghPassword string

func init() {
	renderCmd.Flags().BoolVarP(&wetRun, "wet-run", "w", false, "Create and merge pull requests")

	renderCmd.Flags().StringVarP(&ghUsername, "username", "u", "none", "Username for GitHub")
	renderCmd.Flags().StringVarP(&ghPassword, "token", "t", "none", "Access token or Password for GitHub")
	renderCmd.MarkFlagRequired("token")
	renderCmd.MarkFlagRequired("username")
	rootCmd.AddCommand(renderCmd)
}
