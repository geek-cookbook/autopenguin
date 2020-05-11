{{ define "header" }}
[cookbookurl]: https://geek-cookbook.funkypenguin.co.nz
[discourseurl]: https://discourse.geek-kitchen.funkypenguin.co.nz
[discordurl]: http://chat.funkypenguin.co.nz
[patreonurl]: https://patreon.com/funkypenguin
[blogurl]: https://www.funkypenguin.co.nz
[twitchurl]: https://www.twitch.tv/funkypenguinz
[twitterurl]: https://twitter.com/funkypenguin
[dockerurl]: https://geek-cookbook.funkypenguin.co.nz/ha-docker-swarm/design
[k8surl]: https://geek-cookbook.funkypenguin.co.nz/kubernetes/start
[![geek-cookbook](https://raw.githubusercontent.com/geek-cookbook/autopenguin/master/images/readme_header.png)][cookbookurl]
[![Discord](https://img.shields.io/discord/396055506072109067?color=black&label=Hot%20Sweaty%20Geeks&logo=discord&logoColor=white&style=for-the-badge)][discordurl]
[![Forums](https://img.shields.io/discourse/topics?color=black&label=Forums&logo=discourse&logoColor=white&server=https%3A%2F%2Fdiscourse.geek-kitchen.funkypenguin.co.nz&style=for-the-badge)][discourseurl]
[![Cookbook](https://img.shields.io/badge/Recipes-44-black?style=for-the-badge&color=black)][cookbookurl]
[![Twitch Status](https://img.shields.io/twitch/status/funkypenguinnz?style=for-the-badge&label=LiveGeeking&logoColor=white&logo=twitch)][twitchurl]

:wave: Welcome, traveller!
> The [Geek Cookbook][cookbookurl] is a collection of geek-friendly "recipes" to run popular applications on [Docker Swarm][dockerurl] or [Kubernetes][k8surl], in a progressive, easy-to-follow format.  ***Come and [join us][discordurl], fellow geeks!*** :neckbeard:
{{ end }}

{{ define "contents" }} 
# Contents
{{ range $i, $section := .Sections }}
{{ add $i 1 }}. [{{ $section.Title }}](#{{ regexReplaceAll " " (regexReplaceAll "[^\\w\\- ]" ($section.Title | lower) "") "-" }})
{{- end }}
{{ end }}

{{ define "how_to_use_chart" }} 
# How to use it?

Use helm to add the repo:

```
helm repo add geek-cookbook https://geek-cookbook.github.io/charts/
```

Then simply install using helm, for example

```
kubectl create namespace {{ .Repo.Name }}
helm upgrade --install --namespace {{ .Repo.Name }} geek-cookbook/{{ .Repo.Name }}
```

{{ end }}

{{ define "sections" }}
{{ range $i, $section := .Sections }}
# {{ $section.Title }}

{{ $section.Body}}
{{- end }}
{{ end }}
