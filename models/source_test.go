package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	var bs []BookSource
	// bookSource, err := ioutil.ReadFile("54good.json")
	bookSource, err := ioutil.ReadFile("bs_test.json")

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bookSource, &bs)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range bs {
		BSCache.Add(b.BookSourceURL, b, 0)
	}
	fmt.Printf("total %d book sources.\n", BSCache.ItemCount())
}

var sr = make(SearchOutput)

func TestSource(t *testing.T) {
	for i, _ := range BSCache.Items() {
		if b, ok := BSCache.Get(i); ok {
			bs, ok := b.(BookSource)
			if ok {
				fmt.Printf("searching with %s\n", bs)
				fmt.Println(bs.SearchBook("明朝败家子"))
			}
		}
	}
}

func TestBook(t *testing.T) {
	book := Book{}
	fmt.Println("===========Book Start===========")
	book.FromURL("http://www.wzzw.la/33/33705/")
	fmt.Println(book.GetChapterList())
	fmt.Println(book.GetTitle())
	fmt.Println(book.GetIntroduce())
	fmt.Println(book.GetAuthor())
	fmt.Println("===========Book End=============")
}

func TestChapter(t *testing.T) {
	c := Chapter{}
	c.FromURL("http://www.b5200.net/96_96421/154221199.html")
	fmt.Println(c.GetContent())
}
