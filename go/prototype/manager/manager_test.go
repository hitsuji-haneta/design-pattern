package manager_test

import (
	"testing"

	"github.com/hitsuji-haneta/design-pattern/go/prototype/box"
	"github.com/hitsuji-haneta/design-pattern/go/prototype/manager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("Manager", func() {
	Context(".Create", func() {

		It("should return a clone of product", func() {
			sut := manager.New()

			starbox := box.New("*")
			sut.Register("star box", starbox)
			actual := sut.Create("star box")
			Expect(actual).To(Equal(starbox))
		})
	})
})
