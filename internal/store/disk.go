package store

import (
	"brag/internal/brag"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Disk struct {
	baseDir string
}

func NewDiskStore(name string) (*Disk, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot get user home dir: %w", err)
	}
	return &Disk{baseDir: filepath.Join(home, ".brag", name)}, nil
}

func (d *Disk) IsInitialized() (bool, error) {
	return d.exists()
}

func (d *Disk) Initialize() error {
	if err := os.MkdirAll(d.baseDir, 0755); err != nil {
		return fmt.Errorf("cannot create bragdoc: %w", err)
	}
	return nil
}

func (d *Disk) Save(item *brag.BragItem) error {
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

func (d *Disk) exists() (bool, error) {
	info, err := os.Stat(d.baseDir)
	if err == nil {
		return info.IsDir(), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, fmt.Errorf("could not stat %s: %w", d.baseDir, err)
}

func (d *Disk) ReadAll() ([]brag.BragItem, error) {
	files, err := os.ReadDir(d.baseDir)
	if err != nil {
		return nil, fmt.Errorf("cannot read from %s: %v", d.baseDir, err)
	}

	var bragItem brag.BragItem
	var filePaths []string
	var bragItemList []brag.BragItem

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filePaths = append(filePaths, filepath.Join(d.baseDir, f.Name()))
	}

	for _, filePath := range filePaths {
		fContent, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("cannot read from file %s: %v", filePath, err)
		}

		if err := json.Unmarshal(fContent, &bragItem); err != nil {
			return nil, fmt.Errorf("cannot parse file %s: %v", filePath, err)
		}

		bragItemList = append(bragItemList, bragItem)
	}

	return bragItemList, nil
}
