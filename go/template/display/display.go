package display

type Printer interface {
	Open() string
	Print(bool) string
	Close() string
}

type AbstractDisplay struct {
	Printer Printer
}

func (ad *AbstractDisplay) Display(num int) string {

	result := ad.Printer.Open()
	for i:=0; i<num; i++ {
		result += ad.Printer.Print(i == num-1)
	}
	result += ad.Printer.Close()

	return result
}