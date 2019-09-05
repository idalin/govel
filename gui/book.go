package main

import (
	"fmt"

	"github.com/therecipe/qt/core"

	"github.com/idalin/govel/models"
)

const (
	Name = int(core.Qt__UserRole) + 1<<iota
	Author
	NoteURL
	CoverURL
	ChapterURL
	Tag
	Origin
	Kind
	LastChapter
	Introduce
)

type Book struct {
	core.QObject
	// *models.Book

	_ string          `property:"name"`
	_ string          `property:"author"`
	_ string          `property:"noteUrl"`
	_ string          `property:"coverUrl"`
	_ string          `property:"chapterUrl"`
	_ models.UnixTime `property:"finalRefreshDate"`
	_ string          `property:"tag"`
	_ string          `property:"origin"`
	_ string          `property:"kind"`
	_ string          `property:"lastChapter"`
	_ string          `property:"introduce"`
}

func init() {
	Book_QRegisterMetaType()
}

type AbstractBookListModel struct {
	core.QAbstractListModel
	_ func()                               `constructor:"init"`
	_ map[int]*core.QByteArray             `property:"roles"`
	_ []*Book                              `property:"books"`
	_ func(*Book)                          `slot:updateShelf`
	_ func(*Book)                          `slot:"addBook"`
	_ func(row int, key int, value string) `slot:"editBook"`
}

func (a *AbstractBookListModel) init() {
	a.SetRoles(map[int]*core.QByteArray{
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
	a.ConnectData(a.data)

	a.ConnectRowCount(a.rowCount)
	a.ConnectRoleNames(a.roleNames)
	a.ConnectAddBook(a.addBook)
	a.ConnectEditBook(a.editBook)
	log.Debug("Start AbstractBookListModel init.")
}

func (a *AbstractBookListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	// fmt.Printf("data called. index:%d, role:%d\n", index.Row(), role)
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(a.Books()) {
		return core.NewQVariant()
	}

	var b = a.Books()[index.Row()]

	switch role {
	case Name:
		{
			// fmt.Println("data Name called.")
			return core.NewQVariant1(b.Name())
		}
	case Author:
		{
			// fmt.Println("data Author called.")
			return core.NewQVariant1(b.Author())
		}
	case NoteURL:
		{
			// fmt.Println("data NoteURL called.")
			return core.NewQVariant1(b.NoteUrl())
		}
	case CoverURL:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant1(b.CoverUrl())
		}
	case ChapterURL:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant1(b.ChapterUrl())
		}
	case Tag:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant1(b.Tag())
		}
	case Origin:
		{
			// fmt.Println("data Cover_url called.")
			return core.NewQVariant1(b.Origin())
		}
	case Kind:
		{
			// fmt.Println("data Kind called.")
			return core.NewQVariant1(b.Kind())
		}
	case LastChapter:
		{
			// fmt.Println("data LastChapter called.")
			return core.NewQVariant1(b.LastChapter())
		}

	case Introduce:
		{
			// fmt.Println("data Intro called.")
			return core.NewQVariant1(b.Introduce())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}
func (a *AbstractBookListModel) rowCount(*core.QModelIndex) int {
	return len(a.Books())
}

func (a *AbstractBookListModel) roleNames() map[int]*core.QByteArray {
	return a.Roles()
}

func (a *AbstractBookListModel) addBook(book *Book) {
	a.BeginInsertRows(core.NewQModelIndex(), len(a.Books()), len(a.Books()))
	a.SetBooks(append(a.Books(), book))
	a.EndInsertRows()
}

func (a *AbstractBookListModel) editBook(row int, key int, value string) {
	var b = a.Books()[row]
	switch key {
	case Origin:
		b.SetOrigin(fmt.Sprintf("%s %s", b.Origin(), value))
	case CoverURL:
		b.SetCoverUrl(value)
	}
	var bIndex = a.Index(row, 0, core.NewQModelIndex())
	a.DataChanged(bIndex, bIndex, []int{key})
}

func (a *AbstractBookListModel) add(item *models.Book) {
	book := NewBook(nil)
	book.SetName(item.Name)
	book.SetAuthor(item.Author)
	book.SetNoteUrl(item.NoteURL)
	book.SetCoverUrl(item.CoverURL)
	book.SetKind(item.Kind)
	book.SetChapterUrl(item.ChapterURL)
	book.SetFinalRefreshDate(item.FinalRefreshDate)
	book.SetOrigin(item.Origin)
	book.SetLastChapter(item.LastChapter)
	book.SetIntroduce(item.Introduce)
	a.AddBook(book)
}
