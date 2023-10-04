package main

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//get_cache_entries()
	if _, err := os.Stat("./blog.cache"); err == nil {
		readFile, err := os.Open("./blog.cache")
		defer readFile.Close()
		if err != nil {
			log.Fatal("Error opening blog cache file.")
		}
		scanner := bufio.NewScanner(readFile)
		for scanner.Scan() {
			//fmt.Println(scanner.Text())
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Print("There is no blog cache file! Generating...")
		_, err := os.Create("./blog.cache")
		if err != nil {
			log.Fatal(err)
		} else {

		}
	} else {
		log.Print(err)
	}
	fs := http.FileServer(http.Dir("./site"))
	go test_blog_dir()
	http.Handle("/", fs)
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func test_blog_dir() {
	entries, err := os.ReadDir("./blog")
	var blogs []fs.DirEntry
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		s := strings.Split(entry.Name(), ".")
		if s[len(s)-1] == "blog" {
			file2hash(entry)
			blogs = append(blogs, entry)
		}
	}

	for _, e := range blogs {
		log.Print("Found:", e.Name())
	}
}
func get_cache_entries() map[string]string {
	file, err := os.Open("./blog.cache")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cached := map[string]string{}
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ":")
		a, b := res[0], res[1]
		cached[a] = b
	}

	//fmt.Print(cached)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return cached
}

func file2hash(myFile fs.DirEntry) string {
	f, err := os.Open("./blog/" + myFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hash := md5.New()
	_, err = io.Copy(hash, f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s MD5 checksum is %x \n", f.Name(), hash.Sum(nil))
	return string(hash.Sum(nil))

}
