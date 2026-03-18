package brag

import (
	"errors"
	"strings"
	"time"
)

type BragItem struct {
	Title       string
	Description string
	CreatedAt   time.Time
	Path        string
}

func NewBragItem(title, description string) (*BragItem, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("title cannot be blank")
	}

	return &BragItem{
		Title:       title,
		Description: description,
		Path:        "",
		CreatedAt:   time.Now(),
	}, nil
}
