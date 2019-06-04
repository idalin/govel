package models

type BookShelf struct {
	AllowUpdate      bool      `json:"allowUpdate"`
	BookInfoBean     *BookInfo `json:"bookInfoBean"`
	ChapterListSize  int       `json:"chapterListSize"`
	DurChapter       int       `json:"durChapter"`
	DurChapterName   string    `json:"durChapterName"`
	DurChapterPage   int       `json:"durChapterPage"`
	FinalDate        int64     `json:"finalDate"`
	FinalRefreshData int64     `json:"finalRefreshData"`
	Group            int       `json:"group"`
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
	FinalRefreshData int64         `json:"finalRefreshData"`
	Introduce        string        `json:"introduce"`
	Name             string        `json:"name"`
	NoteURL          string        `json:"noteUrl"`
	Origin           string        `json:"origin"`
	Tag              string        `json:"tag"`
}
