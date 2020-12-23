package adapter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hitsuji-haneta/design-pattern/go/adapter/printbanner"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

type Print interface {
	PrintWeak() string
	PrintStrong() string
}

var _ = Describe("PrintBanner", func() {
	var print Print

	Context("on Print interface", func() {
		print = printbanner.New("Hello, world")

		It("returns weak sting with parenthesis", func() {
			Expect(print.PrintWeak()).To(Equal("(Hello, world)"))
		})

		It("returns strong sting with asterisks", func() {
			Expect(print.PrintStrong()).To(Equal("*Hello, world*"))
		})
	})
})
