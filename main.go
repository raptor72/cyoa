package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
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
    for idx, value := range(story) {
        fmt.Println(idx, value)
        fmt.Println("####################")
    }
}