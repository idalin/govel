package storage

import (
	"encoding/json"
	"io/ioutil"

	// "log"

	"github.com/idalin/govel/models"
)

var fs = &FileStorage{}

func init() {
	// fs.BasePath = "/home/dalin/go/src/github.com/idalin/govel/storage/cache"
	fs.BasePath = "./cache"
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

// func TestFileStorage(t *testing.T) {
// 	book := models.Book{}
// 	// book.FromURL("https://www.zwdu.com/book/39025/")
// 	// book.FromURL("https://www.zwdu.com/book/41228/")
// 	book.FromURL("http://www.b5200.net/96_96421/")
// 	// fmt.Println(book.GetChapterList())
// 	err := fs.SaveBook(&book)
// 	if err != nil {
// 		log.Error(err.Error())
// 	}

// }
