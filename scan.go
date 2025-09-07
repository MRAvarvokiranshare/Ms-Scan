package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Site struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run scan.go <username>")
		return
	}
	username := os.Args[1]

	file, _ := os.ReadFile("sites.json")
	var sites []Site
	_ = json.Unmarshal(file, &sites)

	for _, site := range sites {
		url := site.URL
		url = replace(url, "{u}", username)
		resp, err := http.Head(url)
		if err == nil && resp.StatusCode == 200 {
			fmt.Printf("%s: ✅ %s\n", site.Name, url)
		} else {
			fmt.Printf("%s: ❌ %s\n", site.Name, url)
		}
	}
}

func replace(str, old, new string) string {
	return string([]byte(fmt.Sprintf("%s", str)))
}
