{{ define "header" }}
[cookbookurl]: https://geek-cookbook.funkypenguin.co.nz
[kitchenurl]: https://discourse.kitchen.funkypenguin.co.nz
[discordurl]: http://chat.funkypenguin.co.nz
[patreonurl]: https://patreon.com/funkypenguin
[blogurl]: https://www.funkypenguin.co.nz
[hub]: https://hub.docker.com/r/funkypenguin/munin-node/

[![geek-cookbook](https://raw.githubusercontent.com/funkypenguin/www.funkypenguin.co.nz/master/images/geek-kitchen-banner.png)][cookbookurl]
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
helm repo add funkypenguins-geek-cookbook-{{ .Repo.Name }} \
  https://funkypenguins-geek-cookbook.github.io/{{ .Repo.Name }}/
```

Then simply install using helm:

```
helm upgrade --install --namespace {{ .Repo.Name }} {{ .Repo.Name }} \
  funkypenguins-geek-cookbook-{{ .Repo.Name }}
```

<aside class="notice">
srsly bro? Why such a long name? Because we've moved from a mono-repo for helm charts, to a repo per-chart. This simplifies PR dependencies, and make it easier to track only the charts you're interested
in, by watching each repo
</aside>

{{ end }}

{{ define "sections" }}
{{ range $i, $section := .Sections }}
# {{ $section.Title }}

{{ $section.Body}}
{{- end }}
{{ end }}