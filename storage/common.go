package storage

import (
	"github.com/apsdehal/go-logger"

	"github.com/idalin/govel/models"
	"github.com/idalin/govel/utils"
)

var log *logger.Logger

func init() {
	log = utils.GetLogger()
}

type Storage interface {
	SaveBook(book *models.Book) error
	SaveChapter(book *models.Book, chapter *models.Chapter) error
}
