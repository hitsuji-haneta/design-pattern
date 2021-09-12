package box

import "fmt"

type Box struct {
	decorator string
}

func New(deco string) *Box {
	return &Box{deco}
}

func (b *Box) Use(msg string) string {
	var line string

	length := len(msg) + 4
	for i := 0; i < length; i++ {
		line += b.decorator
	}

	box := "\n" +
		line + "\n" +
		b.decorator + " %s " + b.decorator + "\n" +
		line + "\n"

	return fmt.Sprintf(box, msg)
}
