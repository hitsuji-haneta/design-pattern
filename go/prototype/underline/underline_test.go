package underline_test

import (
	"testing"

	"github.com/hitsuji-haneta/design-pattern/go/prototype/underline"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("Box", func() {
	Context(".Use()", func() {

		It("should return a message with underline", func() {
			sut := underline.New("*")
			expected := `
Hello
*****
`
			Expect(sut.Use("Hello")).To(Equal(expected))
		})

		It("should return a message with underline on various messages", func() {
			sut := underline.New("*")
			expected := `
Hello, my name is Alice.
************************
`

			Expect(sut.Use("Hello, my name is Alice.")).To(Equal(expected))
		})

		It("should return a message with underline on various decorators", func() {
			sut := underline.New("=")
			expected := `
Hello
=====
`
			Expect(sut.Use("Hello")).To(Equal(expected))
		})
	})
})
