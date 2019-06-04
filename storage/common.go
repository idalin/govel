package storage

import (
	"github.com/idalin/govel/models"
)

type Storage interface {
	SaveBook(book *models.Book) error
	SaveChapter(book *models.Book, chapter *models.Chapter) error
}
