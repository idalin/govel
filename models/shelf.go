package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/idalin/govel/utils"
)

var Shelf = &BookShelf{}

func InitShelf(fileName string) {
	Shelf.FilePath = fileName
	if !utils.IsExist(fileName) {
		_, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		items, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.Unmarshal(items, &Shelf.Books)
	}
}

type UnixTime time.Time

func (t *UnixTime) UnmarshalJSON(data []byte) (err error) {
	r := strings.Replace(string(data), `"`, ``, -1)
	ti, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(ti/1000, 0)
	return
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	ts := fmt.Sprintf("%v", time.Time(t).Unix()*1000)
	return []byte(ts), nil
}

type BookShelf struct {
	FilePath string
	Books    Items
}

type Items []*ShelfItem

type ShelfItem struct {
	AllowUpdate      bool      `json:"allowUpdate"`      // 是否允许更新
	BookInfoBean     *BookInfo `json:"bookInfoBean"`     // 书籍信息
	ChapterListSize  int       `json:"chapterListSize"`  // 章节数
	DurChapter       int       `json:"durChapter"`       // 在读章节index
	DurChapterName   string    `json:"durChapterName"`   // 在读章节名字
	DurChapterPage   int       `json:"durChapterPage"`   // 在读章节页数
	FinalDate        UnixTime  `json:"finalDate"`        // 最后阅读时间？
	FinalRefreshDate UnixTime  `json:"finalRefreshData"` // 最后更新时间 "阅读"有typo,这里应该是 finalRefreshDate
	Group            int       `json:"group"`            // 分组
	HasUpdate        bool      `json:"hasUpdate"`
	IsLoading        bool      `json:"isLoading"`
	LastChapterName  string    `json:"lastChapterName"`
	NewChapters      int       `json:"newChapters"`
	NoteURL          string    `json:"noteUrl"`
	SerialNumber     int       `json:"serialNumber"`
	Tag              string    `json:"tag"`
	UseReplaceRule   bool      `json:"useReplaceRule"`
}

type BookInfo struct {
	Author           string        `json:"author"`
	BookmarkList     []interface{} `json:"bookmarkList"`
	ChapterURL       string        `json:"chapterUrl"`
	CoverURL         string        `json:"coverUrl"`
	FinalRefreshData UnixTime      `json:"finalRefreshData"`
	Introduce        string        `json:"introduce"`
	Name             string        `json:"name"`
	NoteURL          string        `json:"noteUrl"`
	Origin           string        `json:"origin"`
	Tag              string        `json:"tag"`
}

func (s Items) Len() int {
	return len(s)
}

// 实现sort的接口，用来给书架排序，最后阅读时间近的排前面
func (s Items) Less(i, j int) bool {
	return time.Time(s[i].FinalDate).Unix() > time.Time(s[j].FinalDate).Unix()
}

func (s Items) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (b *BookShelf) Sort() {
	sort.Sort(b.Books)
}

func (b *BookShelf) Save() error {
	b.Sort()
	shelves, err := json.MarshalIndent(b.Books, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(b.FilePath, shelves, 0644)
	return err
}
