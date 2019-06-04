package utils

import (
	"fmt"
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetPage(t *testing.T) {
	// p, err := GetPage("https://m.zhaishuyuan.com")
	p, err := GetPage("http://m.zhaishuyuan.com/book/28361", "")
	if err != nil {
		log.Fatal(err)
	}
	// page, err := ioutil.ReadAll(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(page))
	doc, err := goquery.NewDocumentFromReader(p)
	doc.Find("#author").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
	c := doc.Find(".booktitle").Find("h1").Text()
	fmt.Println(c)
}
