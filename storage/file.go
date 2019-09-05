package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/idalin/govel/models"
	"github.com/idalin/govel/utils"
)

type FileStorage struct {
	BasePath string
}

func (f *FileStorage) SaveBook(book *models.Book) error {
	log.DebugF("saving book: %v.", book)
	bookDir, err := f.GetBookDir(book)
	if err != nil {
		return err
	}

	if !utils.IsExist(bookDir) {
		err := os.MkdirAll(bookDir, 0755)
		if err != nil {
			return err
		}
	}

	if len(book.GetChapterList()) > 0 {
		for _, c := range book.GetChapterList() {
			// fmt.Printf("saving chapter %s\n", c)
			f.SaveChapter(book, c)
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

	return nil
}

func (f *FileStorage) GetBookDir(book *models.Book) (string, error) {
	if f.BasePath == "" {
		return "", errors.New("No book store path.")
	}
	if book.Tag == "" || book.GetName() == "" {
		return "", errors.New("book name or book source null.")
	}
	re := regexp.MustCompile("[./:]")
	site := re.ReplaceAllString(book.Tag, "")
	bookDirName := fmt.Sprintf("%s-%s", book.GetName(), site)
	if !utils.IsExist(f.BasePath) {
		return "", errors.New(fmt.Sprintf("book cache path: %s not exists.", f.BasePath))
	}
	bookDir := filepath.Join(f.BasePath, bookDirName)
	return bookDir, nil
}

func (f *FileStorage) GetBook(book *models.Book) error {
	bookDir, err := f.GetBookDir(book)
	if err != nil {
		return err
	}
	bookInfoFile, err := f.GetBookInfoFile(book)
	if err != nil {
		return err
	}
	bookInfo, err := ioutil.ReadFile(bookInfoFile)
	if err != nil {
		return err
	}
	// var b models.Book
	err = json.Unmarshal(bookInfo, &book)
	log.Debug(book.Name)
	log.Debug(book.Introduce)
	chapters, err := ioutil.ReadDir(bookDir)
	if err != nil {
		return err
	}
	for _, cFile := range chapters {
		if strings.HasSuffix(cFile.Name(), "nb") {
			log.Debug(cFile.Name())
			cArray := strings.Split(cFile.Name(), "-")
			index, err := strconv.Atoi(cArray[0])
			if err != nil {
				log.Error(err.Error())
			}
			title := strings.Replace(strings.Join(cArray[1:], ""), ".nb", "", 1)
			log.DebugF("index:%d, title:%s\n", index, title)
			chapter := models.Chapter{
				Index:        index,
				ChapterTitle: title,
			}
			err = f.GetChapter(book, &chapter)
			if err != nil {
				log.Error(err.Error())
			}
			log.Debug(chapter.Content)
		}
	}
	return nil
}

func (f *FileStorage) GetBookInfoFile(book *models.Book) (string, error) {
	bookDir, err := f.GetBookDir(book)
	if err != nil {
		return "", err
	}
	return filepath.Join(bookDir, "info.json"), nil
}

func (f *FileStorage) SaveChapter(book *models.Book, chapter *models.Chapter) error {
	// log.DebugF("saving chapter %v.", chapter)
	chapterFileName := f.GetChapterFile(book, chapter)
	if utils.IsExist(chapterFileName) {
		return errors.New("file exists,pass.")
	}
	content := chapter.GetContent()
	// re := regexp.MustCompile("(\n)+")
	// content = re.ReplaceAllString(content, "\n　　")
	content = fmt.Sprintf("%s\n\n　　%s", chapter.GetTitle(), content)
	err := ioutil.WriteFile(chapterFileName, []byte(content), 0644)
	return err
}

func (f *FileStorage) GetChapterFile(book *models.Book, chapter *models.Chapter) string {
	bookDir, err := f.GetBookDir(book)
	if err != nil {
		return ""
	}
	if chapter.GetIndex() == -1 || chapter.GetTitle() == "" {
		return ""
	}
	return filepath.Join(bookDir, fmt.Sprintf("%05d-%s.nb", chapter.GetIndex(), chapter.GetTitle()))
}

func (f *FileStorage) GetChapter(book *models.Book, chapter *models.Chapter) error {
	chapterFileName := f.GetChapterFile(book, chapter)
	if chapterFileName == "" {
		return errors.New("invalid chapter.")
	}
	c, err := ioutil.ReadFile(chapterFileName)
	if err != nil {
		return err
	}
	cArray := strings.Split(string(c), "\n\n")
	if len(cArray) >= 2 {
		chapter.ChapterTitle = cArray[0]
		chapter.Content = strings.Join(cArray[1:], "")
	}

	return nil
}
