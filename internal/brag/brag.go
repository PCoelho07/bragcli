package brag

import "fmt"

type Brag struct {
	fs  FileSystem
	opt Options
}

type Options struct {
	Name        string
	Description string
}

type FileSystem interface {
	Mkdir(path string) error
    Exists(dirname string) (bool, error)
}

func New(fs FileSystem, opt Options) *Brag {
	return &Brag{
        fs: fs,
        opt: opt,
    }
}

func (b *Brag) Init() error {
    if err := b.validateOpts(); err != nil { 
        return fmt.Errorf("invalid name or description: %v", err)
    }

    exists, err := b.fs.Exists(b.opt.Name)
    if err != nil { 
        return fmt.Errorf("could not check if %s exists: %v", b.opt.Name, err)
    }

    if exists { 
        return nil
    }

    if err := b.fs.Mkdir(b.opt.Name); err != nil { 
        return fmt.Errorf("could not init brag: %v", err)
    }

	return nil
}

func (b *Brag) validateOpts() error { 
    if b.opt.Name == "" { 
        return fmt.Errorf("name must be set")
    }

    return nil
}
