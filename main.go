package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "html/template"
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
        <a href="/cyoa/{{ .Arc }}">
        {{ .Text }}
        </a>
        </li>
        {{end}}
        </ul>
    </body>
</html>`

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

    t := template.New("fieldname example")
    t, _ = t.Parse(tmpl)


    for _, value := range(story) {
//        fmt.Println(idx, value)
//        fmt.Println("####################")
//      fmt.Println(value.Options)
//        Options := value.Options
        t.Execute(os.Stdout, value)
    }

//    fmt.Println(story["debate"])

}