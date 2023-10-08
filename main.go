package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	go startup_routine()
	fs := http.FileServer(http.Dir("./site"))
	http.Handle("/", fs)
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func startup_routine() {

	var dir_blogs []string = get_dir_blogs()
}

func get_dir_blogs() []string {
	entries, err := os.ReadDir("./blog")
	var blogs []string
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		s := strings.Split(entry.Name(), ".")
		if s[len(s)-1] == "blog" {
			blogs = append(blogs, entry.Name())
		}
	}
	for _, e := range blogs {
		log.Print("Found:", e)
	}
	return blogs
}
