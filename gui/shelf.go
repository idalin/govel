package main

import (
	"github.com/therecipe/qt/core"

	"github.com/idalin/govel/models"
)

const (
	AllowUpdate = Introduce + 1<<iota
	ChapterListSize
	DurChapter
	DurChapterName
	DurChapterPage
	FinalDate
	FinalRefreshDate
	Group
	HasUpdate
	IsLoading
	LastChapterName
	NewChapters
	SerialNumber
	UseReplaceRule
)

func init() {
	ShelfItem_QRegisterMetaType()
	ShelfListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "ShelfListModel")
}

type ShelfItem struct {
	core.QObject

	_ bool            `property:"allowUpdate"`      // 是否允许更新
	_ *Book           `property:"bookInfoBean"`     // 书籍信息
	_ int             `property:"chapterListSize"`  // 章节数
	_ int             `property:"durChapter"`       // 在读章节index
	_ string          `property:"durChapterName"`   // 在读章节名字
	_ int             `property:"durChapterPage"`   // 在读章节页数
	_ models.UnixTime `property:"finalDate"`        // 最后阅读时间？
	_ models.UnixTime `property:"finalRefreshData"` // 最后更新时间 "阅读"有typo,这里应该是 finalRefreshDate
	_ int             `property:"group"`            // 分组
	_ bool            `property:"hasUpdate"`
	_ bool            `property:"isLoading"`
	_ string          `property:"lastChapterName"`
	_ int             `property:"newChapters"`
	_ string          `property:"noteUrl"`
	_ int             `property:"serialNumber"`
	_ string          `property:"tag"`
	_ bool            `property:"useReplaceRule"`
}

type ShelfListModel struct {
	core.QAbstractListModel
	_ func()                               `constructor:"init"`
	_ map[int]*core.QByteArray             `property:"roles"`
	_ []*ShelfItem                         `property:"books"`
	_ func(*ShelfItem)                     `slot:"updateShelf"`
	_ func(*ShelfItem)                     `slot:"addBook"`
	_ func(row int, key int, value string) `slot:"editBook"`
}

func (s *ShelfListModel) init() {
	log.Debug("Start ShelfListModel init.")
	s.SetRoles(map[int]*core.QByteArray{
		Name:             core.NewQByteArray2("name", len("name")),
		Author:           core.NewQByteArray2("author", len("author")),
		NoteURL:          core.NewQByteArray2("noteUrl", len("noteUrl")),
		CoverURL:         core.NewQByteArray2("coverUrl", len("CoverUrl")),
		ChapterURL:       core.NewQByteArray2("chapterUrl", len("chapterUrl")),
		Tag:              core.NewQByteArray2("tag", len("tag")),
		Origin:           core.NewQByteArray2("origin", len("origin")),
		Kind:             core.NewQByteArray2("kind", len("kind")),
		LastChapter:      core.NewQByteArray2("lastChapter", len("lastChapter")),
		Introduce:        core.NewQByteArray2("introduce", len("introduce")),
		AllowUpdate:      core.NewQByteArray2("allowUpdate", len("allowUpdate")),
		ChapterListSize:  core.NewQByteArray2("chapterListSize", len("chapterListSize")),
		DurChapter:       core.NewQByteArray2("durChapter", len("durChapter")),
		DurChapterName:   core.NewQByteArray2("durChapterName", len("durChapterName")),
		DurChapterPage:   core.NewQByteArray2("durChapterPage", len("durChapterPage")),
		FinalDate:        core.NewQByteArray2("finalDate", len("finalDate")),
		FinalRefreshDate: core.NewQByteArray2("finalRefreshDate", len("finalRefreshDate")),
		Group:            core.NewQByteArray2("group", len("group")),
		HasUpdate:        core.NewQByteArray2("hasUpdate", len("hasUpdate")),
		IsLoading:        core.NewQByteArray2("isLoading", len("isLoading")),
		LastChapterName:  core.NewQByteArray2("lastChapterName", len("lastChapterName")),
		NewChapters:      core.NewQByteArray2("newChapters", len("newChapters")),
		SerialNumber:     core.NewQByteArray2("serialNumber", len("serialNumber")),
		UseReplaceRule:   core.NewQByteArray2("useReplaceRule", len("useReplaceRule")),
	})
	s.ConnectData(s.data)

	s.ConnectRowCount(s.rowCount)
	s.ConnectRoleNames(s.roleNames)
	s.ConnectAddBook(s.addBook)
	s.ConnectEditBook(s.editBook)
	for _, v := range models.Shelf.Books {
		s.add(v)
	}
	// s.AbstractBookListModel.init()
}

func (s *ShelfListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	log.DebugF("data called. index:%d, role:%d\n", index.Row(), role)
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(s.Books()) {
		return core.NewQVariant()
	}
	var shelf = s.Books()[index.Row()]
	var b = shelf.BookInfoBean()

	switch role {
	case Name:
		{
			// log.Debug("data Name called.")
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
	// case AllowUpdate:
	// 	{
	// 		return core.NewQVariant9(shelf.AllowUpdate())
	// 	}
	case ChapterListSize:
		{
			return core.NewQVariant5(shelf.ChapterListSize())
		}
	case DurChapter:
		{
			return core.NewQVariant5(shelf.DurChapter())
		}
	case DurChapterName:
		{
			return core.NewQVariant1(shelf.DurChapterName())
		}
	case DurChapterPage:
		{
			return core.NewQVariant5(shelf.DurChapterPage())
		}
	case FinalDate:
		{
			return core.NewQVariant1(shelf.FinalDate())
		}
	// case FinalRefreshDate:
	// 	{
	// 		return core.NewQVariant1(shelf.FinalRefreshDate())
	// 	}
	case Group:
		{
			return core.NewQVariant5(shelf.Group())
		}
	// case HasUpdate:
	// 	{
	// 		return core.NewQVariant9(shelf.HasUpdate())
	// 	}
	case IsLoading:
		{
			return core.NewQVariant9(shelf.IsLoading())
		}
	case LastChapterName:
		{
			return core.NewQVariant1(shelf.LastChapterName())
		}
	case NewChapters:
		{
			return core.NewQVariant5(shelf.NewChapters())
		}
	case SerialNumber:
		{
			return core.NewQVariant5(shelf.SerialNumber())
		}
	// case UseReplaceRule:
	// 	{
	// 		return core.NewQVariant9(shelf.UseReplaceRule())
	// 	}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (s *ShelfListModel) rowCount(*core.QModelIndex) int {
	log.DebugF("roleCount Called.role Count:%d", len(s.Books()))
	return len(s.Books())
}

func (s *ShelfListModel) roleNames() map[int]*core.QByteArray {
	log.Debug("roleNames Called.")
	return s.Roles()
}

func (s *ShelfListModel) add(item *models.ShelfItem) {
	shelf := NewShelfItem(nil)
	book := NewBook(nil)
	b := item.BookInfoBean
	book.SetName(b.Name)
	book.SetAuthor(b.Author)
	book.SetNoteUrl(b.NoteURL)
	book.SetCoverUrl(b.CoverURL)
	book.SetKind(b.Kind)
	book.SetChapterUrl(b.ChapterURL)
	book.SetFinalRefreshData(b.FinalRefreshData)
	book.SetOrigin(b.Origin)
	book.SetLastChapter(b.LastChapter)
	book.SetIntroduce(b.Introduce)
	shelf.SetBookInfoBean(book)
	shelf.SetAllowUpdate(item.AllowUpdate)
	shelf.SetChapterListSize(item.ChapterListSize)
	shelf.SetDurChapter(item.DurChapter)
	shelf.SetDurChapterName(item.DurChapterName)
	shelf.SetDurChapterPage(DurChapterPage)
	shelf.SetFinalDate(item.FinalDate)
	// shelf.SetFinalRefreshDate(item.FinalRefreshDate)
	shelf.SetGroup(item.Group)
	shelf.SetHasUpdate(item.HasUpdate)
	shelf.SetIsLoading(item.IsLoading)
	shelf.SetLastChapterName(item.LastChapterName)
	shelf.SetNewChapters(item.NewChapters)
	shelf.SetSerialNumber(item.SerialNumber)
	shelf.SetUseReplaceRule(item.UseReplaceRule)
	s.AddBook(shelf)
}

func (s *ShelfListModel) addBook(shelf *ShelfItem) {
	s.BeginInsertRows(core.NewQModelIndex(), len(s.Books()), len(s.Books()))
	s.SetBooks(append(s.Books(), shelf))
	s.EndInsertRows()
}

func (s *ShelfListModel) editBook(row int, key int, value string) {
	// var b = s.Books()[row]
	// switch key {
	// case Origin:
	// 	b.SetOrigin(fmt.Sprintf("%s %s", b.Origin(), value))
	// case CoverURL:
	// 	b.SetCoverUrl(value)
	// }
	// var bIndex = s.Index(row, 0, core.NewQModelIndex())
	// s.DataChanged(bIndex, bIndex, []int{key})
}
