package idcard

import "fmt"

type idCard struct {
	owner string
}

func (ic *idCard) Use() string {
	return fmt.Sprintf("This is %s's ID.", ic.owner)
}

func (ic *idCard) GetOwner() string {
	return ic.owner
}
