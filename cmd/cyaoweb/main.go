package main

import (
    "os"
    "fmt"
    "net/http"
    "flag"
    "cyoa"
    "log"
    // "html/template" // need if we pass template
)


func main() {
    port := flag.Int("port", 8080, "the  port start the CYOA web application on")
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

    // way to pass custom template inside
    // tpl := template.Must(template.New("").Parse("Hello!"))
    // h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))

    h := cyoa.NewHandler(story) // keep the default options
    fmt.Printf("Starting the server on http://127.0.0.1:%d\n", *port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
