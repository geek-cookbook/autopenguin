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

{{ define "sections" }}
{{ range $i, $section := .Sections }}
# {{ $section.Title }}

{{ $section.Body}}
{{- end }}
{{ end }}