package main

import (
	"fmt"

	"github.com/hitsuji-haneta/design-pattern/go/iterator"
)

func main() {
	fmt.Println("===start===")
	if err := iterator.Run(); err != nil {
		fmt.Println("The main process got an exception:", err)
	}
	fmt.Println("====end====")
}
