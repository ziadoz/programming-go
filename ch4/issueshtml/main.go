// Usage: go run main.go is:open json decoder > issues.html
//        go run main.go is:open laravel eloquent > issues.html
// 		  ./issueshtml repo:golang/go commenter:gopherbot json encoder > issues.html
package main

import (
	"html/template"
	"log"
	"os"

	"gopl.io/ch4/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}}</h1>
<table>
<tr style="text-align: left">
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href="{{.HTMLURL}}">{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href="{{.User.HTMLURL}}">{{.User.Login}}</a></td>
	<td><a href="{{.HTMLURL}}">{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
