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
}

func NewBragItem(title, description string) (*BragItem, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("title cannot be blank")
	}

	if strings.TrimSpace(description) == "" {
		return nil, errors.New("description cannot be blank")
	}

	return &BragItem{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
