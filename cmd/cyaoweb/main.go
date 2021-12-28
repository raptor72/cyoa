package main

import (
    "os"
    "fmt"
    "net/http"
    "flag"
    "cyoa"
    "log"
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

    h := cyoa.NewHandler(story)
    fmt.Printf("Starting the server on http://127.0.0.1:%d\n", *port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
