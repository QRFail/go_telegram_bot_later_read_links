package storage

import (
	"crypto/sha1"
	"fmt"
	"go_telegram_bot_later_read_links/lib/e"
	"io"
)

type Storage interface {
	save(p *Page) error
	pickRandom(userName string) (p *Page, err error)
	delete(p *Page) error
	isExist(p *Page) bool
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("error write URL in hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("error write UserName in hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
