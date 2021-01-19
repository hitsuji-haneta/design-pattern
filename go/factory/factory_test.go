package adapter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hitsuji-haneta/design-pattern/go/factory/idcard"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("IDCard", func() {
	Context(".Use()", func() {
		It("should return a text message.", func() {
			factory := idcard.NewFactory()
			card1 := factory.Create("Alice")
			card2 := factory.Create("Bob")
			Expect(card1.Use()).To(Equal("This is Alice's ID."))
			Expect(card2.Use()).To(Equal("This is Bob's ID."))
		})
	})
})

var _ = Describe("Factory", func() {
	Context(".GetOwners()", func() {
		It("should return owner names.", func() {
			factory := idcard.NewFactory()
			_ = factory.Create("Alice")
			_ = factory.Create("Bob")
			Expect(factory.GetOwners()).To(Equal([]string{"Alice", "Bob"}))
		})
	})
})
