package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/idalin/govel/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type FileStorage struct {
	BasePath string
}

func (f *FileStorage) SaveBook(book *models.Book) error {
	bookDir, err := f.GetBookDir(book)
	if err != nil {
		return err
	}

	if _, err := os.Stat(bookDir); os.IsNotExist(err) {
		err := os.MkdirAll(bookDir, 0755)
		if err != nil {
			return err
		}
	}

	bookInfoFile, err := f.GetBookInfoFile(book)
	if err != nil {
		return err
	}
	bookInfo, _ := json.MarshalIndent(book, "", " ")
	err = ioutil.WriteFile(bookInfoFile, bookInfo, 0644)
	if err != nil {
		return err
	}

	if len(book.GetChapterList()) > 0 {
		for _, c := range book.GetChapterList() {
			// fmt.Printf("saving chapter %s\n", c)
			f.SaveChapter(book, c)
		}
	}
	return nil
}

func (f *FileStorage) GetBookDir(book *models.Book) (string, error) {
	if f.BasePath == "" {
		return "", errors.New("No book store path.")
	}
	if book.BookSourceSite == "" || book.GetTitle() == "" {
		return "", errors.New("book name or book source null.")
	}
	re := regexp.MustCompile("[./:]")
	site := re.ReplaceAllString(book.BookSourceSite, "")
	bookDirName := fmt.Sprintf("%s-%s", book.GetTitle(), site)
	if _, err := os.Stat(f.BasePath); os.IsNotExist(err) {
		return "", errors.New(fmt.Sprintf("book cache path: %s not exists.", f.BasePath))
	}
	bookDir := filepath.Join(f.BasePath, bookDirName)
	return bookDir, nil
}

func (f *FileStorage) GetBookInfoFile(book *models.Book) (string, error) {
	bookDir, err := f.GetBookDir(book)
	if err != nil {
		return "", err
	}
	return filepath.Join(bookDir, "info.json"), nil
}

func (f *FileStorage) SaveChapter(book *models.Book, chapter *models.Chapter) error {
	bookDir, err := f.GetBookDir(book)
	if err != nil {
		return err
	}
	chapterFileName := filepath.Join(bookDir, fmt.Sprintf("%05d-%s.nb", chapter.GetIndex(), chapter.GetTitle()))
	if _, err := os.Stat(chapterFileName); os.IsExist(err) {
		return errors.New("file exists,pass.")
	}
	content := chapter.GetContent()
	re := regexp.MustCompile("(\n)+")
	content = re.ReplaceAllString(content, "\n  ")
	content = fmt.Sprintf("%s\n\n%s", chapter.GetTitle(), content)
	err = ioutil.WriteFile(chapterFileName, []byte(content), 0644)
	return err
}
