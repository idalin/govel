package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"
	// "log"
	"net/url"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/apsdehal/go-logger"
	"github.com/patrickmn/go-cache"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/idalin/govel/utils"
)

var BSCache *cache.Cache = cache.New(0, 0)
var log *logger.Logger

func init() {
	log = utils.GetLogger()
}

type SearchOutput map[string][]*Book

func InitBS(fileName string) {
	var bs []BookSource
	bookSource, err := ioutil.ReadFile(fileName)

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
}

func SortSearchOutput(so SearchOutput) []string {
	sortedResult := make(map[string]int, len(so))
	// var keys = make([]int, len(so))
	var newKeys = make([]string, len(so))
	// var result = &SearchOutput{}
	for k, v := range so {
		sortedResult[k] = len(v)
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range sortedResult {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	for _, kv := range ss {
		// fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		newKeys = append(newKeys, kv.Key)
	}
	return newKeys
}

type BookSource struct {
	BookSourceGroup       string `json:"bookSourceGroup"`
	BookSourceName        string `json:"bookSourceName"`
	BookSourceURL         string `json:"bookSourceUrl"`
	CheckURL              string `json:"checkUrl"`
	Enable                bool   `json:"enable"`
	HTTPUserAgent         string `json:"httpUserAgent"`
	RuleBookAuthor        string `json:"ruleBookAuthor"`
	RuleBookContent       string `json:"ruleBookContent"`
	RuleBookName          string `json:"ruleBookName"`
	RuleChapterList       string `json:"ruleChapterList"`
	RuleChapterName       string `json:"ruleChapterName"`
	RuleChapterURL        string `json:"ruleChapterUrl"`
	RuleChapterURLNext    string `json:"ruleChapterUrlNext"`
	RuleContentURL        string `json:"ruleContentUrl"`
	RuleCoverURL          string `json:"ruleCoverUrl"`
	RuleFindURL           string `json:"ruleFindUrl"`
	RuleIntroduce         string `json:"ruleIntroduce"`
	RuleSearchAuthor      string `json:"ruleSearchAuthor"`
	RuleSearchCoverURL    string `json:"ruleSearchCoverUrl"`
	RuleSearchKind        string `json:"ruleSearchKind"`
	RuleSearchLastChapter string `json:"ruleSearchLastChapter"`
	RuleSearchList        string `json:"ruleSearchList"`
	RuleSearchName        string `json:"ruleSearchName"`
	RuleSearchNoteURL     string `json:"ruleSearchNoteUrl"`
	RuleSearchURL         string `json:"ruleSearchUrl"`
	SerialNumber          int    `json:"serialNumber"`
	Weight                int    `json:"weight"`
}

func (bs BookSource) String() string {
	return fmt.Sprintf("%s( %s )", bs.BookSourceName, bs.BookSourceURL)
}

type SearchResult struct {
	BookSourceSite string `json:"source"`
	BookTitle      string `json:"name"`
	Author         string `json:"author"`
	BookURL        string `json:"book_url"`
	CoverURL       string `json:"cover_url"`
	Kind           string `json:"kind"`
	LastChapter    string `json:"last_chapter"`
	NoteURL        string `json:"note_url"`
}

/*
例:http://www.gxwztv.com/search.htm?keyword=searchKey&pn=searchPage-1
- ?为get @为post
- searchKey为关键字标识,运行时会替换为搜索关键字,
- searchPage,searchPage-1为搜索页数,从0开始的用searchPage-1,
- page规则还可以写成
{index（第一页）,
indexSecond（第二页）,
indexThird（第三页）,
index-searchPage+1 或 index-searchPage-1 或 index-searchPage}
- 要添加转码编码在最后加 |char=gbk
- |char=escape 会模拟js escape方法进行编码
如果搜索结果可能会跳到简介页请填写简介页url正则
*/
func (b *BookSource) SearchBook(title string) []*Book {
	if b.RuleSearchURL == "" || b.RuleSearchURL == "-" {
		return nil
	}
	searchUrl := b.RuleSearchURL

	// Process encoding transform
	if strings.Contains(searchUrl, "|char") {
		charParam := strings.Split(searchUrl, "|")[1]
		searchUrl = strings.Replace(searchUrl, fmt.Sprintf("|%s", charParam), "", -1)
		charEncoding := strings.Split(charParam, "=")[1]
		charEncoding = strings.ToLower(charEncoding)
		if charEncoding == "gbk" || charEncoding == "gb2312" || charEncoding == "gb18030" {
			data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(title)), simplifiedchinese.GBK.NewEncoder()))
			title = string(data)
		}
	}

	var err error
	var p io.Reader
	searchUrl = strings.Replace(searchUrl, "=searchKey", fmt.Sprintf("=%s", url.QueryEscape(title)), -1)
	searchUrl = strings.Replace(searchUrl, "searchPage-1", "0", -1)
	searchUrl = strings.Replace(searchUrl, "searchPage", "1", -1)
	// if searchUrl contains "@", searchKey should be post, not get.
	if b.searchMethod() == "post" {
		data := strings.Split(searchUrl, "@")[1]
		params := strings.Replace(data, "=searchKey", fmt.Sprintf("=%s", url.QueryEscape(title)), -1)
		p, err = utils.PostPage(strings.Split(searchUrl, "@")[0], params, b.HTTPUserAgent)
	} else {
		log.Debug(searchUrl)
		p, err = utils.GetPage(searchUrl, b.HTTPUserAgent)
	}

	if err != nil {
		log.ErrorF("searching book error:%s\n", err.Error())
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(p)
	if err != nil {
		log.ErrorF("searching book error:%s\n", err.Error())
		return nil
	}
	if doc == nil {
		log.Error("doc is nil.")
		return nil
	}
	return b.extractSearchResult(doc)

}

func (b *BookSource) searchMethod() string {
	if strings.Contains(b.RuleSearchURL, "@") {
		return "post"
	}
	return "get"
}

func (b *BookSource) searchPage() int {
	if !strings.Contains(b.RuleSearchURL, "searchPage") {
		return -1
	}
	if strings.Contains(b.RuleSearchURL, "searchPage-1") {
		return 0
	}
	return 1
}

func (b *BookSource) extractSearchResult(doc *goquery.Document) []*Book {
	var srList []*Book
	sel, str := utils.ParseRules(doc, b.RuleSearchList)
	if sel != nil {
		sel.Each(func(i int, s *goquery.Selection) {
			_, title := utils.ParseRules(s, b.RuleSearchName)
			if title != "" {
				_, url := utils.ParseRules(s, b.RuleSearchNoteURL)
				_, author := utils.ParseRules(s, b.RuleSearchAuthor)
				_, kind := utils.ParseRules(s, b.RuleSearchKind)
				_, cover := utils.ParseRules(s, b.RuleSearchCoverURL)
				_, lastChapter := utils.ParseRules(s, b.RuleSearchLastChapter)
				_, noteURL := utils.ParseRules(s, b.RuleSearchNoteURL)
				if strings.HasPrefix(url, "/") {
					url = fmt.Sprintf("%s%s", b.BookSourceURL, url)
				}
				if strings.HasPrefix(cover, "/") {
					cover = fmt.Sprintf("%s%s", b.BookSourceURL, cover)
				}
				if strings.HasPrefix(noteURL, "/") {
					noteURL = fmt.Sprintf("%s%s", b.BookSourceURL, noteURL)
				}
				sr := &Book{
					Tag:         b.BookSourceURL,
					Name:        title,
					Author:      author,
					Kind:        kind,
					CoverURL:    cover,
					LastChapter: lastChapter,
					NoteURL:     noteURL,
					Origin:      b.BookSourceName,
				}
				srList = append(srList, sr)

			}
		})

	} else {
		log.ErrorF("No search result found. string:%s\n", str)
	}

	return srList
}

func SearchBooks(title string) SearchOutput {
	c := make(chan *Book, 5)
	defer close(c)
	result := make(SearchOutput)

	for i, _ := range BSCache.Items() {
		go func(i string) {
			if b, ok := BSCache.Get(i); ok {
				bs, ok := b.(BookSource)
				if ok {
					searchResult := bs.SearchBook(title)
					if searchResult != nil {
						for _, sr := range searchResult {
							c <- sr
						}
					}
				} else {
					log.ErrorF("not book source.")
				}
			}
		}(i)
	}

	for {
		select {
		case i, ok := <-c:
			if ok {
				log.DebugF("Got result:%s, %s\n", i, i.Tag)
				if _, ok := result[i.Name]; !ok {
					result[i.Name] = []*Book{i}
				} else {
					result[i.Name] = append(result[i.Name], i)
				}
			}
		case <-time.After(5 * time.Second):
			log.DebugF("Timeout,exiting...\n")
			goto end
		}
	}
end:
	log.DebugF("exited!\n")

	for _, key := range SortSearchOutput(result) {
		if key != "" {
			resultJson, _ := json.MarshalIndent(result[key], "", "    ")
			log.DebugF("%s:\n %s\n", key, resultJson)
		}
	}
	return result
}
