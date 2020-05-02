package repo

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"gopkg.in/yaml.v2"
)

//GetRepos gets the repositories for the organization specified
func GetRepos(organization string, httpClient *http.Client) (map[string]string, error) {
	client := github.NewClient(httpClient)

	repoMap := map[string]string{}

	repos, _, err := client.Repositories.ListByOrg(context.Background(), organization, &github.RepositoryListByOrgOptions{Type: "public"})
	if err != nil {
		return repoMap, err
	}

	for _, repo := range repos {
		repoMap[*repo.Name] = repo.GetCloneURL()
	}

	return repoMap, err
}

type RepoConfig struct {
	Name   string
	README struct {
		Template string
		Update   bool
		Sections []struct {
			Title string
			Body  string
		}
	}
}

func GetRepoConfig(file []byte) (RepoConfig, error) {
	cfg := RepoConfig{}
	err := yaml.Unmarshal(file, &cfg)
	return cfg, err
}
