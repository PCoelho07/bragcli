package fs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Disk struct{}

func NewDisk() *Disk {
	return &Disk{}
}


func (d *Disk) Mkdir(path string) error {
    home, err := os.UserHomeDir()
    if err != nil { 
        return fmt.Errorf("cannot get user home dir: %v", err)
    }

	if err := os.MkdirAll(filepath.Join(home, ".brag", path), 0755); err != nil {
		return fmt.Errorf("cannot create %s: %v", path, err)
	}

	return nil
}

func (d *Disk) Exists(dirname string) (bool, error) {
    home, err := os.UserHomeDir()
    if err != nil { 
        return false, fmt.Errorf("cannot get user home dir: %v", err)
    }

	info, err := os.Stat(filepath.Join(home, ".brag", dirname))
	if err == nil {
		return info.IsDir(), nil
	}

    if errors.Is(err, os.ErrNotExist) { 
        return false, nil
    }

    return false, fmt.Errorf("could not get any info from %s: %v", dirname, err)
}
