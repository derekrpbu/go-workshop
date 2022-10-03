package main

import (
	"books/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {

	server = gin.Default()

}

func main() {
	basepath := server.Group("/books")

	basepath.GET("/getTitle/:title", getBookByTitle)
	basepath.GET("/getKey/:key", getBookByKey)

	log.Fatal(server.Run(":9090"))
}

func getBookByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	resp, err := http.Get("https://openlibrary.org/works/" + key + ".json")
	if err != nil {
		log.Fatalln(err)
	}

	var result models.Key
	respbody, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(respbody, &result); err != nil { // Parse []byte to the go struct pointer
		log.Fatalln(err)
	}

	ctx.JSON(resp.StatusCode, result)
}

func getBookByTitle(ctx *gin.Context) {
	title := ctx.Param("title")
	resp, err := http.Get("https://openlibrary.org/search.json?title=" + title)
	if err != nil {
		log.Fatalln(err)
	}

	var result models.Book
	respbody, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(respbody, &result); err != nil { // Parse []byte to the go struct pointer
		log.Fatalln(err)
	}
	ctx.JSON(resp.StatusCode, result)
}
