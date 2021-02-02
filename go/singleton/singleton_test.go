package singleton_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hitsuji-haneta/design-pattern/go/singleton"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("singleton", func() {
	Context(".GetInstanse()", func() {
		It("should return the same instances.", func() {
			obj1 := singleton.GetInstanse()
			obj2 := singleton.GetInstanse()
			Expect(obj1 == obj2).To(BeTrue())
		})
	})
})
