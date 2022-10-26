package books_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ChocolatePadmanaban/ginkgo-tutorials/books_one"
)

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})
	Context("Time tests", func() {
		It("Check Times", func() {
			time.Sleep(1 * time.Minute)
			first := time.Now()
			fmt.Println("First", first)
		})
	})

	Describe("Categorizing book length", func() {
		fmt.Println("in second describe", time.Now())
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})

	It("can extract the author's last name", func() {
		Expect(longBook.AuthorLastName()).To(Equal("Hugo"))
	})

	It("interprets a single author name as a last name", func() {
		longBook.Author = "Hugo"
		Expect(longBook.AuthorLastName()).To(Equal("Hugo"))
	})

	It("can extract the author's first name", func() {
		Expect(longBook.AuthorFirstName()).To(Equal("Victor"))
	})

	It("returns no first name when there is a single author name", func() {
		longBook.Author = "Hugo"
		Expect(longBook.AuthorFirstName()).To(BeZero()) //BeZero asserts the value is the zero-value for its type.  In this case: ""
	})

})
