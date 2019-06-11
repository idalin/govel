package models

import (
	"sort"
	"testing"
	"time"
)

func TestShelf(t *testing.T) {
	shelfFile := "/home/dalin/go/src/govel/YueDu/autoSave/myBookShelf.json"

	InitShelf(shelfFile)
	for k, v := range Shelf.Books {
		log.DebugF("book[%d] is %s, author:%s, upadte time:%s\n", k, v.BookInfoBean.Name, v.BookInfoBean.Author, time.Time(v.FinalDate))
	}
	log.InfoF("Is shelf sorted?%v", sort.IsSorted(Shelf.Books))
	Shelf.Sort()
	for k, v := range Shelf.Books {
		log.DebugF("book[%d] is %s, author:%s, upadte time:%s\n", k, v.BookInfoBean.Name, v.BookInfoBean.Author, time.Time(v.FinalDate))
	}
	if err := Shelf.Save(); err != nil {
		log.Error(err.Error())
	}
	return
}
