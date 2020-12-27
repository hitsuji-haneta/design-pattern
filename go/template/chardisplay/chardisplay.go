package chardisplay

import (
	"github.com/hitsuji-haneta/design-pattern/go/template/display"
)

type charDisplay struct {
	*display.AbstractDisplay
	char rune
}

func New(char rune) *charDisplay {
	cd := &charDisplay{&display.AbstractDisplay{}, char}
	cd.Printer = cd
	return cd
}

func (cd *charDisplay) Open() string {
	return "<"
}

func (cd *charDisplay) Print(flag bool) string {
	return string(cd.char)
}

func (cd *charDisplay) Close() string {
	return ">"
}