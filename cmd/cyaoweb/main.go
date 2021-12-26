package main

import (
    "os"
    "fmt"
    "html/template"
    "net/http"
    "strings"
    "flag"
    "cyoa"
)

const tmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Choose Tour Own Adventure</title>
    </head>
    <body>
        <section class="page">
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
        </section>
    <style>
        body {
           font-family: helvetica, arial;
        }
        h1 {
            text-align:center;
            position:relative;
        }
        .page {
            width: 80%;
            max-width: 500px;
            margin: auto;
            margin-top: 40px;
            margin-bottom: 40px;
            padding: 80px;
            background: #FFFCF6;
            border: 1px solid #eee;
            box-shadow: 0 10px 6 px -6px #777;
        }
        ul {
            border-top: 1px dotted #ccc
            padding: 10px 0 0 0;
            -webkit-padding-start: 0;
        }
        li {
            padding-top: 10px;
        }
        a,
        a:visited {
            text-decoration: none:
            color: #6295b5;
        }
        a:active,
        a:hover {
            color: #7792a2;
        }
        p {
            text-indent: 1em;
        }
    </style>
    </body>
</html>`

var t = template.New("fieldname example")

func MapHandler(pathsToUrls cyoa.Story, fallback http.Handler) http.HandlerFunc {
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
    filename := flag.String("file", "gopher.json", "The JSON file with CYOA story")
    flag.Parse()
    fmt.Printf("Using the story in %s.\n", *filename)

    f, err := os.Open(*filename)
    if err != nil {
        fmt.Println(err)
    }

    story, err := cyoa.JsonStory(f)
    if err != nil {
        fmt.Println(err)
    }

    t, _ = t.Parse(tmpl)

    mux := defaultMux()
    mapHandler := MapHandler(story, mux)
    fmt.Println("Starting the server on :8080")
    http.ListenAndServe(":8080", mapHandler)
}
