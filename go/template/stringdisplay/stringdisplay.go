package stringdisplay

import (
	"github.com/hitsuji-haneta/design-pattern/go/template/display"
)

type stringDisplay struct {
	*display.AbstractDisplay
	text string
}

func New(text string) *stringDisplay {
	sd := &stringDisplay{&display.AbstractDisplay{}, text}
	sd.Printer = sd
	return sd
}

func (sd *stringDisplay) Open() string {
	return "["
}

func (sd *stringDisplay) Print(flag bool) string {
	if flag {
		return sd.text
	}
	return sd.text + ", "
}

func (sd *stringDisplay) Close() string {
	return "]"
}
