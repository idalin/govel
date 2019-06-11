package storage

import (
	"encoding/json"
	"io/ioutil"
	// "log"
	"testing"

	"github.com/idalin/govel/models"
)

var fs = &FileStorage{}

func init() {
	fs.BasePath = "/home/dalin/go/src/github.com/idalin/govel/storage/cache"
	var bs []models.BookSource
	// bookSource, err := ioutil.ReadFile("54good.json")
	bookSource, err := ioutil.ReadFile("../bs.json")

	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(bookSource, &bs)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, b := range bs {
		models.BSCache.Add(b.BookSourceURL, b, 0)
	}
	log.DebugF("total %d book sources.\n", models.BSCache.ItemCount())
}

func TestFileStorage(t *testing.T) {
	book := models.Book{}
	book.FromURL("https://www.zwdu.com/book/32642/")
	// fmt.Println(book.GetChapterList())
	err := fs.SaveBook(&book)
	if err != nil {
		log.Error(err.Error())
	}

}
