package storage

import (
	"encoding/json"
	"fmt"
	"govel/models"
	"io/ioutil"
	"log"
	"testing"
)

var fs = &FileStorage{}

func init() {
	fs.BasePath = "./cache"
	var bs []models.BookSource
	// bookSource, err := ioutil.ReadFile("54good.json")
	bookSource, err := ioutil.ReadFile("../bs.json")

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bookSource, &bs)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range bs {
		models.BSCache.Add(b.BookSourceURL, b, 0)
	}
	fmt.Printf("total %d book sources.\n", models.BSCache.ItemCount())
}

func TestFileStorage(t *testing.T) {
	book := models.Book{}
	book.FromURL("http://www.b5200.net/96_96421/")
	// fmt.Println(book.GetChapterList())
	err := fs.SaveBook(&book)
	if err != nil {
		fmt.Println(err.Error())
	}
}
