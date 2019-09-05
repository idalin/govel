package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func init() {
	var bs []BookSource
	bookSource, err := ioutil.ReadFile("bs_test.json")
	// bookSource, err := ioutil.ReadFile("bs.json")

	log.Info("Start testing")
	log.Debug("Debug logging.")

	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(bookSource, &bs)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, b := range bs {
		BSCache.Add(b.BookSourceURL, b, 0)
	}
	log.DebugF("total %d book sources.\n", BSCache.ItemCount())
}

var sr = make(SearchOutput)

// func TestSource(t *testing.T) {
// 	log.Info("Start testing of source.")
// 	for i, _ := range BSCache.Items() {
// 		if b, ok := BSCache.Get(i); ok {
// 			bs, ok := b.(BookSource)
// 			if ok {
// 				log.DebugF("searching with %v\n", bs)
// 				log.InfoF("result of search:  %v\n", bs.SearchBook("明朝败家子"))
// 			}
// 		}
// 	}
// }

// func TestBook(t *testing.T) {
// 	book := Book{}
// 	log.Info("===========Book Start===========")
// 	// book.FromURL("http://www.b5200.net/96_96421/") // content 为text的情况
// 	book.FromURL("https://www.zwdu.com/book/32642/") // content为textNodes的情况
// 	log.InfoF("%v\n", book.GetChapterList())
// 	log.InfoF("%v\n", book.GetName())
// 	log.InfoF("%v\n", book.GetIntroduce())
// 	log.InfoF("%v\n", book.GetAuthor())
// 	log.Info("===========Book End=============")
// }

func TestChapter(t *testing.T) {
	log.Info("===========Chapter Start===========")
	c := Chapter{}
	// c.FromURL("http://www.b5200.net/96_96421/168403600.html") // content 为text的情况
	c.FromURL("https://www.zwdu.com/book/32642/17845126.html") // content为textNodes的情况
	log.InfoF("content:\n%s\n", c.GetContent())
	log.Info("===========Chapter End=============")
}
