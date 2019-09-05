package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	// "github.com/766b/mobi"
	"github.com/peterbn/mobi"

	"github.com/idalin/govel/models"
	"github.com/idalin/govel/utils"
)

type MobiStorage struct {
	M        mobi.Builder
	BasePath string
}

func (m *MobiStorage) SaveBook(book *models.Book) error {
	fileName, err := m.GetBookFile(book)
	if err != nil {
		return err
	}
	if utils.IsExist(fileName) {
		return errors.New(fmt.Sprintf("file %s exists.", fileName))
	}
	if m.M == nil {

		m.M = mobi.NewBuilder()
		// if err != nil {
		// 	return err
		// }
	}

	m.M.Title(book.GetName())
	m.M.Compression(mobi.CompressionNone)
	m.M.NewExthRecord(mobi.EXTH_DOCTYPE, "EBOK")
	if book.CoverURL != "" {
		m.M.AddCover(book.CoverURL, book.CoverURL)
	}
	if book.GetAuthor() != "" {
		m.M.NewExthRecord(mobi.EXTH_AUTHOR, book.GetAuthor())
	}
	m.M.CSS("p{ text-indent:2em; padding:0px; margin:0px; }")
	if len(book.GetChapterList()) > 0 {
		for _, c := range book.GetChapterList() {
			m.SaveChapter(book, c)
		}
	}
	// m.M.Write()
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = m.M.WriteTo(f)
	return err
}

func (m *MobiStorage) GetBookFile(book *models.Book) (string, error) {
	if m.BasePath == "" {
		return "", errors.New("No book store path.")
	}
	if book.GetName() == "" {
		return "", errors.New("book name or book source null.")
	}
	basePath, err := filepath.Abs(m.BasePath)
	if err == nil {
		m.BasePath = basePath
	}
	return filepath.Join(m.BasePath, fmt.Sprintf("%s.mobi", book.GetName())), nil
}

func (m *MobiStorage) SaveChapter(book *models.Book, chapter *models.Chapter) error {
	log.DebugF("saving chapter %v.", chapter)
	if m.M == nil {
		return errors.New("Mobi builder not found.")
	}
	content := chapter.GetContent()
	// re := regexp.MustCompile("(\n)+")
	// content = re.ReplaceAllString(content, "\n　　")
	s := strings.Split(content, "\n")
	for i, v := range s {
		s[i] = fmt.Sprintf("<p>%s</p>", strings.TrimSpace(v))
	}
	content = strings.Join(s, "")
	m.M.NewChapter(chapter.GetTitle(), []byte(content))
	return nil
}
