package render

import (
	"html/template"
	"io/ioutil"

	"github.com/Masterminds/sprig"
	"github.com/funkypenguins-geek-cookbook/penguin/pkg/repo"
)

type READMEContext struct {
	Repo struct {
		Name string
	}
	Sections []struct {
		Title string
		Body  string
	}
}

func GetREADMEContext(repository repo.RepoConfig) READMEContext {
	return READMEContext{
		Repo: struct{ Name string }{
			Name: repository.Name,
		},
		Sections: repository.README.Sections,
	}
}

func GetREADMETemplate(repository repo.RepoConfig) (*template.Template, error) {
	tpl := template.New(repository.Name).Funcs(sprig.FuncMap())

	t, err := ioutil.ReadFile("templates/readme/" + repository.README.Template + ".md")
	if err != nil {
		return tpl, err
	}
	t2 := string(t)
	tpl, err = tpl.Parse(t2)

	tpl, err = tpl.ParseGlob("templates/readme/*tpl*")
	return tpl, err
}
