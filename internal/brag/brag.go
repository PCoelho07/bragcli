package brag

import "errors"

type Store interface {
	IsInitialized() (bool, error)
	Initialize() error
	Save(item *BragItem) error
	ReadAll() ([]BragItem, error)
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

func (b *Brag) List(name string) ([]BragItem, error) {

	if err := b.checkInitialized(); err != nil {
		return nil, err
	}

	return b.store.ReadAll()
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
