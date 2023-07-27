package modgen

import "html/template"

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
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
`))

var moduleTmpl = template.Must(template.New("module").Parse(`<!DOCTYPE html>
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
</html>`))
