package main

import (
	"fmt"

	"github.com/therecipe/qt/core"

	"github.com/idalin/govel/models"
)

func init() {
	SearchListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "SearchListModel")
}

type SearchListModel struct {
	core.QAbstractListModel

	_ func()                   `constructor:"init"`
	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Book                  `property:"books"`

	_ func(*Book)                          `slot:"addBook"`
	_ func(row int, key int, value string) `slot:"editBook"`

	// _ func()                                  `signal:"remove,auto"`
	// _ func(book *models.Book) `signal:"add,auto"`
	// _ func(firstName string, lastName string) `signal:"edit,auto"`
	_ func(key string) `signal:"doSearch,auto"`
	// _ func()           `signal:"clear,auto"`
}

func (s *SearchListModel) init() {
	s.SetRoles(map[int]*core.QByteArray{
		Name:        core.NewQByteArray2("name", len("name")),
		Author:      core.NewQByteArray2("author", len("author")),
		NoteURL:     core.NewQByteArray2("noteUrl", len("noteUrl")),
		CoverURL:    core.NewQByteArray2("coverUrl", len("CoverUrl")),
		ChapterURL:  core.NewQByteArray2("chapterUrl", len("chapterUrl")),
		Tag:         core.NewQByteArray2("tag", len("tag")),
		Origin:      core.NewQByteArray2("origin", len("origin")),
		Kind:        core.NewQByteArray2("kind", len("kind")),
		LastChapter: core.NewQByteArray2("lastChapter", len("lastChapter")),
		Introduce:   core.NewQByteArray2("introduce", len("introduce")),
	})
	s.ConnectData(s.data)

	s.ConnectRowCount(s.rowCount)
	s.ConnectRoleNames(s.roleNames)
	s.ConnectAddBook(s.addBook)
	s.ConnectEditBook(s.editBook)
	log.Debug("Start SearchListModel init.")
}

func (s *SearchListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	// fmt.Printf("data called. index:%d, role:%d\n", index.Row(), role)
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(s.Books()) {
		return core.NewQVariant()
	}

	var b = s.Books()[index.Row()]

	switch role {
	case Name:
		{
			// fmt.Println("data Name called.")
			return core.NewQVariant14(b.Name())
		}
	case Author:
		{
			// fmt.Println("data Author called.")
			return core.NewQVariant14(b.Author())
		}
	case NoteURL:
		{
			// fmt.Println("data NoteURL called.")
			return core.NewQVariant14(b.NoteUrl())
		}
	case CoverURL:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant14(b.CoverUrl())
		}
	case ChapterURL:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant14(b.ChapterUrl())
		}
	case Tag:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant14(b.Tag())
		}
	case Origin:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant14(b.Origin())
		}
	case Kind:
		{
			// fmt.Println("data Kind called.")
			return core.NewQVariant14(b.Kind())
		}
	case LastChapter:
		{
			// fmt.Println("data LastChapter called.")
			return core.NewQVariant14(b.LastChapter())
		}

	case Introduce:
		{
			// fmt.Println("data Intro called.")
			return core.NewQVariant14(b.Introduce())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (s *SearchListModel) doSearch(key string) {
	if key == "" {
		return
	}
	go func() {
		for i, _ := range models.BSCache.Items() {
			if b, ok := models.BSCache.Get(i); ok {
				bs, ok := b.(models.BookSource)
				if ok {
					searchResult := bs.SearchBook(key)
					if searchResult != nil {
						for _, sr := range searchResult {
							s.add(sr)
						}
					}
				}
			}
		}
	}()
}

func (s *SearchListModel) rowCount(*core.QModelIndex) int {
	return len(s.Books())
}

func (s *SearchListModel) addBook(book *Book) {
	s.BeginInsertRows(core.NewQModelIndex(), len(s.Books()), len(s.Books()))
	s.SetBooks(append(s.Books(), book))
	s.EndInsertRows()
}

func (s *SearchListModel) editBook(row int, key int, value string) {
	var b = s.Books()[row]
	switch key {
	case Origin:
		b.SetOrigin(fmt.Sprintf("%s %s", b.Origin(), value))
	case CoverURL:
		b.SetCoverUrl(value)
	}
	var bIndex = s.Index(row, 0, core.NewQModelIndex())
	s.DataChanged(bIndex, bIndex, []int{key})
}

func (s *SearchListModel) roleNames() map[int]*core.QByteArray {
	return s.Roles()
}

func (s *SearchListModel) add(item *models.Book) {
	index := s.find(item.Name)
	if index == -1 {
		book := NewBook(nil)
		// if item.Title != "" {
		book.SetName(item.Name)
		// }
		// if item.Author != "" {
		book.SetAuthor(item.Author)
		// }
		// if item.BookURL != "" {
		book.SetNoteUrl(item.NoteURL)
		// }
		if item.CoverURL != "" {
			book.SetCoverUrl(item.CoverURL)
		}
		if item.Kind != "" {
			book.SetKind(item.Kind)
		}
		if item.LastChapter != "" {
			book.SetLastChapter(item.LastChapter)
		}
		if item.Introduce != "" {
			book.SetIntroduce(item.Introduce)
		}
		book.SetTag(item.GetBookSource().BookSourceName)
		s.AddBook(book)
	} else {
		s.editBook(index, Origin, item.GetBookSource().BookSourceName)
		if item.CoverURL != "" {
			fmt.Printf("modified cover of %s to %s\n", item.Name, item.CoverURL)
			s.editBook(index, CoverURL, item.CoverURL)
		}
	}
}

func (s *SearchListModel) find(name string) int {
	for k, v := range s.Books() {
		if v.Name() == name {
			return k
		}
	}
	return -1
}
