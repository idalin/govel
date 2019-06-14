package main

import (
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
	ModelData
)

type Book struct {
	core.QObject
	// *models.Book

	_ string          `property:"name"`
	_ string          `property:"author"`
	_ string          `property:"noteUrl"`
	_ string          `property:"coverUrl"`
	_ string          `property:"chapterUrl"`
	_ models.UnixTime `property:"finalRefreshData"`
	_ string          `property:"tag"`
	_ string          `property:"origin"`
	_ string          `property:"kind"`
	_ string          `property:"lastChapter"`
	_ string          `property:"introduce"`
}

func init() {
	Book_QRegisterMetaType()
}
