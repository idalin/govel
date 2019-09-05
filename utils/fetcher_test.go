package utils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetPage(t *testing.T) {
	// p, err := GetPage("https://m.zhaishuyuan.com")
	p, err := GetPage("http://www.b5200.net/96_96421/154165223.html", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	page, err := ioutil.ReadAll(p)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(page))
	// doc, err := goquery.NewDocumentFromReader(p)
	// doc.Find("#author").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the band and title
	// 	band := s.Find("a").Text()
	// 	title := s.Find("i").Text()
	// 	fmt.Printf("Review %d: %s - %s\n", i, band, title)
	// })
	// c := doc.Find(".booktitle").Find("h1").Text()
	// fmt.Println(c)
}
