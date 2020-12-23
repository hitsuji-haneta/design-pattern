package printbanner

import (
	"github.com/hitsuji-haneta/design-pattern/go/adapter/banner"
)

type Banner interface {
	ShowWithParen() string
	ShowWithAster() string
}

type printBanner struct {
	banner Banner
}

func New(text string) *printBanner {
	return &printBanner{banner.New(text)}
}

func (pb *printBanner) PrintWeak() string {
	return pb.banner.ShowWithParen()
}

func (pb *printBanner) PrintStrong() string {
	return pb.banner.ShowWithAster()
}
