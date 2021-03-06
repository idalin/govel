package models

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/idalin/govel/utils"
)

var (
	getChapterTried int = 0
)

type Chapter struct {
	BookSourceSite string            `json:"source"`
	BookSourceInst *BookSource       `json:"-"`
	Content        string            `json:"content"`
	ChapterTitle   string            `json:"title"`
	Read           bool              `json:"is_read"`
	ChapterURL     string            `json:"url"`
	Index          int               `json:"index"`
	BelongToBook   *Book             `json:"-"`
	Page           *goquery.Document `json:"-"`
}

func (c Chapter) String() string {
	return fmt.Sprintf("%s( %s )", c.ChapterTitle, c.ChapterURL)
}
func (c *Chapter) FromURL(chapterURL string) error {
	if chapterURL == "" {
		return errors.New("no url.")
	}
	_, err := url.ParseRequestURI(chapterURL)
	if err != nil {
		return err
	}
	c.ChapterURL = chapterURL
	c.BookSourceSite = utils.GetHostByURL(c.ChapterURL)
	return nil
}

func (c *Chapter) GetBookSource() *BookSource {
	if c.BookSourceInst != nil {
		return c.BookSourceInst
	}
	if c.BookSourceSite == "" {
		if c.ChapterURL == "" {
			return nil
		}
		c.BookSourceSite = utils.GetHostByURL(c.ChapterURL)
	}
	bs := GetBookSourceByURL(c.BookSourceSite)
	c.BookSourceInst = bs
	return bs
}

func (c *Chapter) getChapterPage() (*goquery.Document, error) {
	if c.Page != nil {
		return c.Page, nil
	}
	bs := c.GetBookSource()
	if c.ChapterURL != "" && bs != nil {
		p, err := utils.GetPage(c.ChapterURL, c.GetBookSource().HTTPUserAgent)
		if err == nil {
			doc, err := goquery.NewDocumentFromReader(p)
			if err == nil {
				c.Page = doc

				return c.Page, err
			}
		}
		return nil, err
	}
	return nil, errors.New("can't get chapter page.")
}

func (c *Chapter) GetContent() string {
	if c.Content != "" {
		return c.Content
	}
	doc, err := c.getChapterPage()
	if err == nil {
		_, content := utils.ParseRules(doc, c.BookSourceInst.RuleBookContent)
		if content != "" {
			// re := regexp.MustCompile("(\b)+")
			// content = re.ReplaceAllString(content, "\n    ")
			c.Content = content
			getChapterTried = 0
			return c.Content
		}
	} else {
		if getChapterTried < 5 {
			getChapterTried++
			time.Sleep(time.Duration(1) * time.Second)
			return c.GetContent()
		} else {
			getChapterTried = 0
		}
		log.DebugF("get content error:%s,retry %d.\n", err.Error(), getChapterTried)
	}
	return c.Content
}

func (c *Chapter) GetTitle() string {
	if c.ChapterTitle != "" {
		return c.ChapterTitle
	}
	return ""
}

func (c *Chapter) GetBook() *Book {
	if c.BelongToBook != nil {
		return c.BelongToBook
	}
	return nil
}

func (c *Chapter) GetIndex() int {
	if c.Index != -1 {
		return c.Index
	}
	return -1
}
