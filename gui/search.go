package main

import (
	"fmt"
	"govel/models"

	"github.com/therecipe/qt/core"
)

const (
	Name = int(core.Qt__UserRole) + 1<<iota
	Author
	Book_url
	Cover_url
	Kind
	Last_chapter
	Note_url
	Intro
	Book_source
)

func init() {
	SearchListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "SearchListModel")
	BookItem_QRegisterMetaType()
}

type SearchListModel struct {
	core.QAbstractListModel

	_ func()                   `constructor:"init"`
	_ map[int]*core.QByteArray `property:"roles"`
	_ []*BookItem              `property:"books"`

	_ func(*BookItem)                      `slot:"addBook"`
	_ func(row int, key int, value string) `slot:"editBook"`

	// _ func()                                  `signal:"remove,auto"`
	// _ func(book *models.Book) `signal:"add,auto"`
	// _ func(firstName string, lastName string) `signal:"edit,auto"`
	_ func(key string) `signal:"doSearch,auto"`
	// _ func()           `signal:"clear,auto"`
}

type BookItem struct {
	core.QObject

	_ string `property:"name"`
	_ string `property:"author"`
	_ string `property:"book_url"`
	_ string `property:"cover_url"`
	_ string `property:"kind"`
	_ string `property:"last_chapter"`
	_ string `property:"note_url"`
	_ string `property:"intro"`
	_ string `property:"book_source"`
}

func (s *SearchListModel) init() {
	s.SetRoles(map[int]*core.QByteArray{
		Name:         core.NewQByteArray2("name", len("name")),
		Author:       core.NewQByteArray2("author", len("author")),
		Book_url:     core.NewQByteArray2("book_url", len("book_url")),
		Cover_url:    core.NewQByteArray2("cover_url", len("cover_url")),
		Kind:         core.NewQByteArray2("kind", len("kind")),
		Last_chapter: core.NewQByteArray2("last_chapter", len("last_chapter")),
		Note_url:     core.NewQByteArray2("note_url", len("note_url")),
		Intro:        core.NewQByteArray2("intro", len("intro")),
		Book_source:  core.NewQByteArray2("book_source", len("book_source")),
	})
	s.ConnectData(s.data)

	s.ConnectRowCount(s.rowCount)
	s.ConnectRoleNames(s.roleNames)
	s.ConnectAddBook(s.addBook)
	s.ConnectEditBook(s.editBook)
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
	case Book_url:
		{
			// fmt.Println("data BookURL called.")
			return core.NewQVariant14(b.Book_url())
		}
	case Cover_url:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant14(b.Cover_url())
		}
	case Kind:
		{
			// fmt.Println("data Kind called.")
			return core.NewQVariant14(b.Kind())
		}
	case Last_chapter:
		{
			// fmt.Println("data LastChapter called.")
			return core.NewQVariant14(b.Last_chapter())
		}
	case Note_url:
		{
			// fmt.Println("data NoteURL called.")
			return core.NewQVariant14(b.Note_url())
		}
	case Intro:
		{
			// fmt.Println("data Intro called.")
			return core.NewQVariant14(b.Intro())
		}
	case Book_source:
		{
			// fmt.Println("data BookSource called.")
			return core.NewQVariant14(b.Book_source())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

// func (s *SearchListModel) clear() {
// 	// s.SetBooks()
// 	s.BeginRemoveRows(core.NewQModelIndex(), row, row)
// 	s.SetBooks(append())
// 	s.EndRemoveRows()
// }
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

func (s *SearchListModel) addBook(book *BookItem) {
	s.BeginInsertRows(core.NewQModelIndex(), len(s.Books()), len(s.Books()))
	s.SetBooks(append(s.Books(), book))
	s.EndInsertRows()
}

func (s *SearchListModel) editBook(row int, key int, value string) {
	var b = s.Books()[row]
	switch key {
	case Book_source:
		b.SetBook_source(fmt.Sprintf("%s %s", b.Book_source(), value))
	case Cover_url:
		b.SetCover_url(value)
	}
	var bIndex = s.Index(row, 0, core.NewQModelIndex())
	s.DataChanged(bIndex, bIndex, []int{key})
}

func (s *SearchListModel) roleNames() map[int]*core.QByteArray {
	return s.Roles()
}

func (s *SearchListModel) add(item *models.Book) {
	index := s.find(item.Title)
	if index == -1 {
		book := NewBookItem(nil)
		// if item.Title != "" {
		book.SetName(item.Title)
		// }
		// if item.Author != "" {
		book.SetAuthor(item.Author)
		// }
		// if item.BookURL != "" {
		book.SetBook_url(item.BookURL)
		// }
		if item.CoverURL != "" {
			book.SetCover_url(item.CoverURL)
		}
		if item.Kind != "" {
			book.SetKind(item.Kind)
		}
		if item.LastChapter != "" {
			book.SetLast_chapter(item.LastChapter)
		}
		if item.NoteURL != "" {
			book.SetNote_url(item.NoteURL)
		}
		if item.Introduce != "" {
			book.SetIntro(item.Introduce)
		}
		book.SetBook_source(item.GetBookSource().BookSourceName)
		s.AddBook(book)
	} else {
		s.editBook(index, Book_source, item.GetBookSource().BookSourceName)
		if item.CoverURL != "" {
			fmt.Printf("modified cover to %s\n", item.CoverURL)
			s.editBook(index, Cover_url, item.CoverURL)
		}
	}
}

func (s *SearchListModel) find(title string) int {
	for k, v := range s.Books() {
		if v.Name() == title {
			return k
		}
	}
	return -1
}
