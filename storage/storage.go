package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/psihachina/telegrambot/lib/e"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExist(p *Page) (bool, error)
}

var ErrNoSavedPage = errors.New("no saved page")

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash:", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash:", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
