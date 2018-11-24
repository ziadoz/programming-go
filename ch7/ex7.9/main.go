package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Movie struct {
	Title  string
	Year   int
	Colour bool
	Actors []string
}

var tmpl string = `
<h1>Movies</h1>
<table>
	<thead>
		<th>Title</th>
		<th>Year</th>
		<th>Colour</th>
		<th>Actors</th>
	</thead>
	<tbody>
		{{range .}}
			<tr>
				<td>{{.Title}}</td>
				<td>{{.Year}}</td>
				<td>{{if .Colour}}Yes{{else}}No{{end}}</td>
				<td>{{StringsJoin .Actors ", "}}</td>
			</tr>
		{{end}}
	</tbody>
</table>
`

var movietmpl = template.Must(template.New("movies").
	Funcs(template.FuncMap{"StringsJoin": strings.Join}).
	Parse(tmpl))

func main() {
	log.SetFlags(0)
	log.SetPrefix("webserver: ")

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeFunc)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	movies := []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Colour: false,
			Actors: []string{
				"Humphrey Bogart",
				"Ingrid Bergman",
			},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Colour: true,
			Actors: []string{
				"Paul Newman",
			},
		},
		{
			Title:  "Bullitt",
			Year:   1968,
			Colour: true,
			Actors: []string{
				"Steve McQueen",
				"Jacqueline Bisset",
			},
		},
	}

	movietmpl.Execute(w, movies)
}
