package presenter

import (
	"brag/internal/brag"
	"fmt"
)

type TextPresenter struct{}

// ── ANSI codes ────────────────────────────────────────────────────────────────
const (
	ansiReset  = "\033[0m"
	ansiBold   = "\033[1m"
	ansiDim    = "\033[2m"
	ansiYellow = "\033[33m"
	ansiBlue   = "\033[34m"
)

var graphColors = []string{
	"\033[31m", // red
	"\033[32m", // green
	"\033[33m", // yellow
	"\033[34m", // blue
	"\033[35m", // magenta
	"\033[36m", // cyan
}

func NewTextPresenter() *TextPresenter {
	return &TextPresenter{}
}

func (bi *TextPresenter) Present(item brag.BragItem) error {
	fmt.Println()
	fmt.Printf("%spath %s%s\n", ansiYellow+ansiBold, item.Path, ansiReset)
	fmt.Printf("%sTitle:        %s%s\n", ansiBlue+ansiBold, item.Title, ansiReset)
	if item.Description != "" {
		fmt.Printf("Description:  %s\n", item.Description)
	}
	fmt.Printf("Created At:   %s\n", item.CreatedAt.Format("02/01/2006 15:04:05"))
	fmt.Println()

	return nil
}
