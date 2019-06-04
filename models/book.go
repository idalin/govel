package models

import (
	"errors"
	"fmt"
	"github.com/idalin/govel/utils"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Book struct {
	BookSourceSite string            `json:"source"`
	Title          string            `json:"name"`
	Author         string            `json:"author"`
	BookURL        string            `json:"book_url"`
	CoverURL       string            `json:"cover_url"`
	Kind           string            `json:"kind"`
	LastChapter    string            `json:"last_chapter"`
	NoteURL        string            `json:"note_url"`
	Introduce      string            `json:"intro"`
	ChapterList    []*Chapter        `json:"-"`
	BookSourceInst *BookSource       `json:"-"`
	Page           *goquery.Document `json:"-"`
}

func (b Book) String() string {
	return fmt.Sprintf("%s( %s )", b.Title, b.BookURL)
}

func (b *Book) GetBookSource() *BookSource {
	if b.BookSourceInst != nil {
		return b.BookSourceInst
	}
	if b.BookSourceSite == "" {
		if b.BookURL == "" {
			return nil
		}
		b.BookSourceSite = utils.GetHostByURL(b.BookURL)
	}
	if bsItem, ok := BSCache.Get(b.BookSourceSite); ok {
		if bs, ok := bsItem.(BookSource); ok {
			b.BookSourceInst = &bs
			return &bs
		} else {
			return nil
		}
	}
	return nil
}

func (b *Book) FromURL(bookURL string) error {
	if bookURL == "" {
		return errors.New("no url.")
	}
	_, err := url.ParseRequestURI(bookURL)
	if err != nil {
		return err
	}
	b.BookURL = bookURL
	b.BookSourceSite = utils.GetHostByURL(b.BookURL)
	b.GetAuthor()
	b.GetIntroduce()
	return nil
}

func (b *Book) FromCache(bookPath string) error {
	if _, err := os.Stat(bookPath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("book path: %s not exists.", bookPath))
	}
	bookName := filepath.Base(bookPath)
	fmt.Printf("book name is: %s.\n", bookName)
	return nil
}

func (b *Book) getBookPage() (*goquery.Document, error) {
	if b.Page != nil {
		return b.Page, nil
	}
	bs := b.GetBookSource()
	if b.BookURL != "" && bs != nil {
		p, err := utils.GetPage(b.BookURL, b.GetBookSource().HTTPUserAgent)
		if err == nil {
			doc, err := goquery.NewDocumentFromReader(p)
			if err == nil {
				b.Page = doc

				return b.Page, err
			}
		}
		return nil, err
	}
	return nil, errors.New("can't get book page.")
}

func (b *Book) GetChapterList() []*Chapter {
	b.UpdateChapterList(len(b.ChapterList))
	return b.ChapterList
}

func (b *Book) UpdateChapterList(startFrom int) error {
	doc, err := b.getBookPage()
	if err != nil {
		return err
	}
	sel, _ := utils.ParseRules(doc, b.BookSourceInst.RuleChapterList)
	if sel != nil {
		sel.Each(func(i int, s *goquery.Selection) {
			if i < startFrom {
				return
			}
			_, name := utils.ParseRules(s, b.BookSourceInst.RuleChapterName)
			_, url := utils.ParseRules(s, b.BookSourceInst.RuleContentURL)
			if strings.HasPrefix(url, "/") {
				url = fmt.Sprintf("%s%s", b.BookSourceInst.BookSourceURL, url)
			}
			b.ChapterList = append(b.ChapterList, &Chapter{
				ChapterTitle: name,
				ChapterURL:   url,
				BelongToBook: b,
				Index:        i,
			})
		})
	}
	return nil
}

func (b *Book) GetTitle() string {
	if b.Title != "" {
		return b.Title
	}
	doc, err := b.getBookPage()
	if err == nil {
		_, title := utils.ParseRules(doc, b.BookSourceInst.RuleBookName)
		if title != "" {
			b.Title = title
		}
	} else {
		fmt.Printf("get title error:%s\n", err.Error()) // for debug
	}
	return b.Title
}

func (b *Book) GetIntroduce() string {
	if b.Introduce != "" {
		return b.Introduce
	}
	doc, err := b.getBookPage()
	if err == nil {
		_, intro := utils.ParseRules(doc, b.BookSourceInst.RuleIntroduce)
		if intro != "" {
			b.Introduce = intro
		}
	} else {
		fmt.Printf("get introduce error:%s\n", err.Error()) // for debug
	}
	return b.Introduce
}

func (b *Book) GetAuthor() string {
	if b.Author == "" {

		doc, err := b.getBookPage()
		if err == nil {
			_, intro := utils.ParseRules(doc, b.BookSourceInst.RuleBookAuthor)
			if intro != "" {
				b.Author = intro
			}
		} else {
			fmt.Printf("get author error:%s\n", err.Error()) // for debug
		}
	}
	return b.Author
}
