package adapter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hitsuji-haneta/design-pattern/go/template/chardisplay"
	"github.com/hitsuji-haneta/design-pattern/go/template/stringdisplay"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("CharDisplay", func() {
	cd := chardisplay.New('A')
	Context(".Display()", func() {
		actual := cd.Display(5)
		It("should return five characters with <>", func() {
			expected := "<AAAAA>"
			Expect(actual).To(Equal(expected))
		})
	})
})

var _ = Describe("StringDisplay", func() {
	cd := stringdisplay.New("Hello")
	Context(".Display()", func() {
		actual := cd.Display(5)
		It("should return five sentences with []", func() {
			expected := "[Hello, Hello, Hello, Hello, Hello]"
			Expect(actual).To(Equal(expected))
		})
	})
})
