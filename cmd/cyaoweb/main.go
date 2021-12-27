package main

import (
    "os"
    "fmt"
    // "html/template"
    "net/http"
    // "strings"
    "flag"
    "cyoa"
    "log"
)



// var t = template.New("fieldname example")
// func StoryHandler(s cyoa.Story, fallback http.Handler) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         clear_path := strings.Replace(r.URL.Path, "/", "", 1)
//         if r.URL.Path == "/" || r.URL.Path == "/info" {
//             clear_path = "intro"
//         }
//         if val, ok := s[clear_path]; ok {
//             t.Execute(w, val)
//             return
//         }
//         fallback.ServeHTTP(w, r)
//     }
// }
// func defaultMux() *http.ServeMux {
//     mux := http.NewServeMux()
//     mux.HandleFunc("/", FallbackHandler)
//     return mux
// }
// func FallbackHandler(w http.ResponseWriter, r *http.Request) {
//     http.NotFound(w, r)
//     return
// }

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

    // t, _ = t.Parse(tmpl)
    // mux := defaultMux()
    // mapHandler := StoryHandler(story, mux)
    h := cyoa.NewHandler(story)
    fmt.Printf("Starting the server on http://127.0.0.1:%d\n", *port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
