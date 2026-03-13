package brag

type Store interface {
	Init() error
	Save(item *BragItem) error
}

type Brag struct {
	store Store
}

func New(s Store) *Brag {
	return &Brag{store: s}
}

func (b *Brag) Init() error {
	return b.store.Init()
}

func (b *Brag) Create(title, description string) error {
	item, err := NewBragItem(title, description)
	if err != nil {
		return err
	}
	return b.store.Save(item)
}
