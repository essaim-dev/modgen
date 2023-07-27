# `modgen`
[![Go Reference](https://pkg.go.dev/badge/go.essaim.dev/modgen.svg)](https://pkg.go.dev/go.essaim.dev/modgen)

Static generator for Go module import URLs. 

## Description
`modgen` generates webpages to serve an Go module import URLs (also known as Vanity URLs).
Unlike similar projects (e.g. [govanityurls](https://github.com/GoogleCloudPlatform/govanityurls/tree/master)), `modgen` was designed to be static. This makes it possible to use static web hosting providers — like Github Pages, for example.
> We are actually using Github Pages to host our own Go module import URLs,
> you can check the Github Actions workflow we use [here](https://github.com/essaim-dev/go.essaim.dev/blob/main/.github/workflows/modgen-gh-pages.yml)!

For more informations on Go module imports, please refer to the official [documentation](https://go.dev/ref/mod#serving-from-proxy).

## Installation
```
go install go.essaim.dev/modgen/cmd/modgen@latest
```

## Usage

```
Usage of modgen:
  -config string
        path of the configuration file (default "modgen.yaml")
  -index-tmpl string
        path of an optional custom index template
  -module-tmpl string
        path of an optional custom module template
  -target string
        path where the site should be generated (default "gen/")
```
### Config file example
```yaml
host: go.essaim.dev

modules:
  - path: /modgen
    vcs: git
    repo-url: https://github.com/essaim-dev/modgen
```
**Note:** All fields are mandatory.

### Index template example
```html
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<title>{{.Host}}</title>
</head>
<body>
	<h1>{{.Host}}</h1>
	<ul>
		{{range .Modules}}<li><a href="{{.Path}}">{{$.Host}}{{.Path}}</a></li>{{end}}
	</ul>
</body>
</html>
```

### Module template example
```html
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<meta name="go-import" content="{{.Path}} {{.VCS}} {{.RepoURL}}">
	<meta http-equiv="refresh" content="1; url=https://pkg.go.dev/{{.Path}}">
</head>
<body>
	<h1>{{.Path}}</h1>
	<p><a href="https://pkg.go.dev/{{.Path}}">See the package on pkg.go.dev</a>.</p>
</body>
</html>
```
