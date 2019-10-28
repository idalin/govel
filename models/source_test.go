package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func init() {
	var bs []BookSource
	// bookSource, err := ioutil.ReadFile("bs_test.json")
	// bookSource, err := ioutil.ReadFile("bs.json")
	bookSource, err := ioutil.ReadFile("bs_ok.json")

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

// func TestSearchBooks(t *testing.T) {
// 	log.Info("start testing searchbooks.")
// 	sr := SearchBooks("诡秘之主")
// 	for k, v := range sr {
// 		log.Infof("%s:\n%v\n", k, v)
// 	}
// }

// func TestSource(t *testing.T) {
// 	log.Info("Start testing of source.")
// 	sf, _ := os.Create("../source.txt")
// 	defer sf.Close()
// 	for i, _ := range BSCache.Items() {
// 		if b, ok := BSCache.Get(i); ok {
// 			bs, ok := b.(BookSource)
// 			if ok {
// 				log.DebugF("searching with %v\n", bs)
// 				result := bs.SearchBook("一品修仙")
// 				log.InfoF("result of search:  %v\n", result)
// 				if len(result) > 0 {
// 					for _, book := range result {
// 						if book.GetName() == "一品修仙" {
// 							fmt.Printf("%s:%s\n%s\n%s\n%s\n%s\n", book.GetName(), book.GetAuthor(), book.GetIntroduce(), book.LastChapter, book.GetCoverURL(), book.NoteURL)
// 						}
// 					}
// 				}

// 			}
// 		}
// 	}
// }

func TestBook(t *testing.T) {
	book := Book{}
	log.Info("===========Book Start===========")
	// book.FromURL("https://www.ymoxuan.com/text_147581.html") // content 为text的情况
	book.FromURL("http://www.b5200.net/96_96421/")
	// book.FromURL("https://www.zwdu.com/book/32642/") // content为textNodes的情况
	log.InfoF("%v\n", book.GetChapterList())
	log.InfoF("%v\n", book.GetName())
	log.InfoF("%v\n", book.GetIntroduce())
	log.InfoF("%v\n", book.GetAuthor())
	log.Info("===========Book End=============")
}

// func TestChapter(t *testing.T) {
// 	log.Info("===========Chapter Start===========")
// 	c := Chapter{}
// 	// c.FromURL("http://www.b5200.net/96_96421/168403600.html") // content 为text的情况
// 	c.FromURL("https://www.zwdu.com/book/32642/17845126.html") // content为textNodes的情况
// 	log.InfoF("content:\n%s\n", c.GetContent())
// 	log.Info("===========Chapter End=============")
// }
