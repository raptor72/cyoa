package cyoa

import (
    "net/http"
    "strings"
    "io"
    "log"
    "encoding/json"
    "html/template"
)

const tmpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Choose Your Own Adventure</title>
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

func NewHandler(s Story) http.Handler {
    return handler{s}
}

type handler struct {
    s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.New("").Parse(tmpl))

    clear_path := strings.TrimSpace(r.URL.Path)

    if clear_path == "" || clear_path == "/" {
        clear_path = "/intro"
    }

    // "/intro" => "intro"
    clear_path = clear_path[1:]

    if chapter, ok := h.s[clear_path]; ok {
        err := tpl.Execute(w, chapter)
        if err != nil {
            log.Printf("%v", err)
            http.Error(w, "Something went wrong...", http.StatusInternalServerError)
        }
        return
    }
    http.Error(w, "Chapter not found.", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
    var story Story
 	if err := d.Decode(&story); err != nil {
	    return nil, err
    }
    return story, nil
}

type Story map[string]Chapter

type Chapter struct {
    Title   string `json:"title"`
    Story   []string `json:"story"`
    Options []ChapterOption `json:"options"`
}

type ChapterOption struct {
   Text string `json:"text"`
   Arc string `json:"arc"`
}