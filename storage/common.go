package storage

import (
	"govel/models"
)

type Storage interface {
	SaveBook(book *models.Book) error
	SaveChapter(book *models.Book, chapter *models.Chapter) error
}
