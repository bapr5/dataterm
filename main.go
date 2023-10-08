package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var entrys []BlogEntry = generateEntrysList(fillBlogs())

type BlogEntry struct {
	ID   int    `json:"id"`
	Name string `name:"name"`
	Date string `json:"date"`
}

func fillBlogs() []string {
	var allblogs []string
	files, err := os.ReadDir("./blog")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			allblogs = append(allblogs, file.Name())
		}
	}
	return allblogs
}

func generateEntrysList(myFileList []string) []BlogEntry {
	var entrys []BlogEntry
	for i, file := range myFileList {
		entrys = append(entrys, BlogEntry{ID: i, Name: file, Date: getFileCreationDate(file)})
	}
	return entrys
}
func getFileCreationDate(fileName string) string {
	fileInfo, err := os.Stat("./blog/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.ModTime().String()
}

func getBlogs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, entrys)
}

func main() {
	//fillBlogs()
	router := gin.Default()

	router.GET("/blogs", getBlogs)
	router.Run("localhost:8000")
}
