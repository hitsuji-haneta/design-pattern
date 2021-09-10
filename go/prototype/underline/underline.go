package underline

type Underline struct {
	decorator string
}

func New(deco string) *Underline {
	return &Underline{deco}
}

func (u *Underline) Use(msg string) string {
	var line string

	length := len(msg)
	for i := 0; i < length; i++ {
		line += u.decorator
	}

	result := "\n" +
		msg + "\n" +
		line + "\n"

	return result
}
