package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
)

type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

var elasticClient *elastic.Client

func main() {
	var err error
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://elasticsearch:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	r := gin.Default()
	r.POST("/documents", createDocumentsEndpoint)
	if err = r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func createDocumentsEndpoint(c *gin.Context) {

}
