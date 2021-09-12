package prototype_test

import (
	"testing"

	"github.com/hitsuji-haneta/design-pattern/go/prototype/box"
	"github.com/hitsuji-haneta/design-pattern/go/prototype/manager"
	"github.com/hitsuji-haneta/design-pattern/go/prototype/underline"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("Protorype", func() {
	It("should create prototypes", func() {
		defer GinkgoRecover()

		proto1 := box.New("*")
		proto2 := box.New("=")
		proto3 := underline.New("~")
		proto4 := underline.New("=")

		m := manager.New()
		m.Register("star box", proto1)
		m.Register("double box", proto2)
		m.Register("wave line", proto3)
		m.Register("double line", proto4)

		expected1 := `
*********
* Hello *
*********
`
		expected2 := `
=========
= Hello =
=========
`
		expected3 := `
Hello
~~~~~
`
		expected4 := `
Hello
=====
`

		Expect(m.Create("star box").Use("Hello")).To(Equal(expected1))
		Expect(m.Create("double box").Use("Hello")).To(Equal(expected2))
		Expect(m.Create("wave line").Use("Hello")).To(Equal(expected3))
		Expect(m.Create("double line").Use("Hello")).To(Equal(expected4))
	})
})
