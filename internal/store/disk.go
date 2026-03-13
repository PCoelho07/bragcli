package store

import (
	"brag/internal/brag"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type DiskStore struct {
	baseDir string
}

func NewDiskStore(name string) (*DiskStore, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot get user home dir: %w", err)
	}
	return &DiskStore{baseDir: filepath.Join(home, ".brag", name)}, nil
}

func (d *DiskStore) Init() error {
	exists, err := d.exists()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	if err := os.MkdirAll(d.baseDir, 0755); err != nil {
		return fmt.Errorf("cannot create bragdoc: %w", err)
	}
	return nil
}

func (d *DiskStore) Save(item *brag.BragItem) error {
	exists, err := d.exists()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("bragdoc not initialized, run 'brag init' first")
	}

	filename := item.CreatedAt.Format("20060102T150405") + ".json"
	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("cannot encode brag item: %w", err)
	}
	if err := os.WriteFile(filepath.Join(d.baseDir, filename), data, 0644); err != nil {
		return fmt.Errorf("cannot save brag item: %w", err)
	}
	return nil
}

func (d *DiskStore) exists() (bool, error) {
	info, err := os.Stat(d.baseDir)
	if err == nil {
		return info.IsDir(), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, fmt.Errorf("could not stat %s: %w", d.baseDir, err)
}
