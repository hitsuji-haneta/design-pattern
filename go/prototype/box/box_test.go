package box_test

import (
	"testing"

	"github.com/hitsuji-haneta/design-pattern/go/prototype/box"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("Box", func() {
	Context(".Use()", func() {

		It("should return a wrapped message box", func() {
			sut := box.New("*")
			expected := `
*********
* Hello *
*********
`
			Expect(sut.Use("Hello")).To(Equal(expected))
		})

		It("should return a message box on various messages", func() {
			sut := box.New("*")
			expected := `
****************************
* Hello, my name is Alice. *
****************************
`

			Expect(sut.Use("Hello, my name is Alice.")).To(Equal(expected))
		})

		It("should return a message box on various decorators", func() {
			sut := box.New("=")
			expected := `
=========
= Hello =
=========
`
			Expect(sut.Use("Hello")).To(Equal(expected))
		})
	})
})
