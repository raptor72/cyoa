package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "html/template"
    "net/http"
    "strings"
)

type chapter struct {
    Title   string `json:"title"`
    Story   []string `json:"story"`
    Options []chapterOption `json:"options"`
}

type chapterOption struct {
   Text string `json:"text"`
   Arc string `json:"arc"`
}

const tmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Choose Tour Own Adventure</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Story}}<p>{{ . }}</p>{{else}}<p><strong>no rows</strong></p>{{end}}
        <ul>
        {{ range .Options}}
        <li>
        <a href="/{{ .Arc }}">
        {{ .Text }}
        </a>
        </li>
        {{end}}
        </ul>
    </body>
</html>`

var t = template.New("fieldname example")

func MapHandler(pathsToUrls map[string]chapter, fallback http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        clear_path := strings.Replace(r.URL.Path, "/", "", 1)
        if r.URL.Path == "/" || r.URL.Path == "/info" {
            clear_path = "intro"
        }
        if val, ok := pathsToUrls[clear_path]; ok {
            t.Execute(w, val)
            return
        }
        fallback.ServeHTTP(w, r)
    }
}

func defaultMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
    return
}


func main() {
    story := make(map[string]chapter)
    dataJSON, err := ioutil.ReadFile("gopher.json")
    if err != nil {
        fmt.Println(err)
    }
    if err := json.Unmarshal(dataJSON, &story); err != nil {
        fmt.Println(err)
        return
    }

    t, _ = t.Parse(tmpl)

    mux := defaultMux()
    mapHandler := MapHandler(story, mux)
    fmt.Println("Starting the server on :8080")
    http.ListenAndServe(":8080", mapHandler)
}
