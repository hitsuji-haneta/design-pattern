package bookshelf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hitsuji-haneta/design-pattern/go/iterator/book"
	"github.com/hitsuji-haneta/design-pattern/go/iterator/bookshelf"
)

func TestVisualizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Suite")
}

var _ = Describe("Bookshelf", func() {
	var (
		bs  *bookshelf.BookShelf
		itr bookshelf.Iterator
	)

	Context("Iterator with next items", func() {
		sampleBook := book.New("sample book")

		BeforeEach(func() {
			bs = bookshelf.New(5)
			bs.AppendBook(sampleBook)
			itr = bs.CreateIterator()
		})

		It("HasNext() should return true", func() {
			Expect(itr.HasNext()).To(Equal(true))
		})
		It("Next() should return the next book", func() {
			Expect(itr.Next()).To(Equal(sampleBook))
		})
	})

	Context("Iterator without next items", func() {
		BeforeEach(func() {
			bs = bookshelf.New(5)
			itr = bs.CreateIterator()
		})

		It("HasNext() should return false", func() {
			Expect(itr.HasNext()).To(Equal(false))
		})
		It("Next() should return nil", func() {
			Expect(itr.Next()).To(BeNil())
		})
	})
})
