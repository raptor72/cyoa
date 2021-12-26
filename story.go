package cyoa

import (
    "io"
	"encoding/json"
)

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