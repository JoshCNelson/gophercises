package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type StoryMap map[string]StoryArc

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func loadStoryFromJson() StoryMap {
	jsonFile, _ := os.Open("./story.json")

	storyJson, _ := ioutil.ReadAll(jsonFile)

	var storymap StoryMap

	json.Unmarshal(storyJson, &storymap)

	return storymap
}

func StoryHandler(s StoryMap) http.Handler {
	h := handler{s}
	return h
}

type handler struct {
	s StoryMap
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	markup := `
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{end}}

		{{range .Options}}
			<p>
				<a href="{{.Arc}}">{{.Text}}
			</p>
		{{end}}
	`

	t, _ := template.New("webpage").Parse(markup)

	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		t.Execute(w, chapter)
	}
}

func main() {
	story := loadStoryFromJson()

	storyHandler := StoryHandler(story)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", storyHandler)
}
