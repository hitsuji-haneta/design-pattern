package banner

import (
	"fmt"
)

type banner struct {
	text string
}

func New(text string) *banner {
	return &banner{text}
}

func (b *banner) ShowWithParen() string {
	return fmt.Sprintf("(%s)", b.text)
}

func (b *banner) ShowWithAster() string {
	return fmt.Sprintf("*%s*", b.text)
}
