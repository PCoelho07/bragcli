package brag

import (
	"errors"
	"fmt"
)

type Store interface {
	IsInitialized() (bool, error)
	Initialize() error
	Save(item *BragItem) error
	ReadAll() ([]BragItem, error)
}

type Presenter interface {
	Present(item BragItem) error
}

type Brag struct {
	store Store
}

func New(s Store) *Brag {
	return &Brag{store: s}
}

var ErrNotInitialized = errors.New("bragdoc not initialized, run 'brag init' first")

func (b *Brag) Init() error {
	err := b.checkInitialized()
	if err == nil {
		return nil
	}

	if !errors.Is(err, ErrNotInitialized) {
		return err
	}

	return b.store.Initialize()
}

func (b *Brag) Create(title, description string) error {
	if err := b.checkInitialized(); err != nil {
		return err
	}

	item, err := NewBragItem(title, description)
	if err != nil {
		return err
	}

	return b.store.Save(item)
}

func (b *Brag) List(name string, presenter Presenter) ([]BragItem, error) {
	if err := b.checkInitialized(); err != nil {
		return nil, err
	}

	items, err := b.store.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read all brag items: %v", err)
	}

	for _, i := range items {
		if err := presenter.Present(i); err != nil {
			return nil, fmt.Errorf("cannot present brag item: %v", err)
		}
	}

	return nil, nil
}

func (b *Brag) checkInitialized() error {
	initialized, err := b.store.IsInitialized()
	if err != nil {
		return err
	}

	if !initialized {
		return ErrNotInitialized
	}

	return nil
}
