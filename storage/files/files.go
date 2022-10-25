package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"go_telegram_bot_later_read_links/lib/e"
	"go_telegram_bot_later_read_links/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(p *storage.Page) (err error) {
	defer func() {
		err = e.WrapIfErr("can`t save page ", err)
	}()

	fPath := filepath.Join(s.basePath, p.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(p)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	if err := gob.NewEncoder(file).Encode(p); err != nil {
		return err
	}

	return nil
}

func (s Storage) PickRandom(UserName string) (p *storage.Page, err error) {
	defer func() {
		err = e.WrapIfErr("can`t pick random page ", err)
	}()
	fPath := filepath.Join(s.basePath, p.UserName)

	files, err := os.ReadDir(fPath)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.New("Not found files")
	}

	rand.Seed(time.Now().UnixNano())

	file := files[rand.Intn(len(files))]

	return s.decodeFile(filepath.Join(s.basePath, file.Name()))

}

func (s Storage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return e.Wrap("can`t remove page ", err)
	}

	fPath := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(fPath); err != nil {
		msg := fmt.Sprintf("can`t remove page %s", fPath)
		return e.Wrap(msg, err)
	}

	return nil
}

func (s Storage) IsExist(p *storage.Page) bool {
	fileName, err := fileName(p)
	if err != nil {
		return false
	}

	fPath := filepath.Join(s.basePath, p.UserName, fileName)

	_, err = os.Stat(fPath)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func (s Storage) decodeFile(filePath string) (p *storage.Page, err error) {
	defer func() {
		err = e.WrapIfErr("can`t decodeFile ", err)
	}()
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	var page storage.Page

	if err := gob.NewDecoder(f).Decode(&page); err != nil {
		return nil, err
	}

	return &page, nil

}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
