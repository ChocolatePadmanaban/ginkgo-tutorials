package books

import "strings"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryByLength() string {

	if b.Pages >= 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}

func (b *Book) AuthorLastName() string {
	nameList := strings.Fields(b.Author)

	return nameList[len(nameList)-1]
}

func (b *Book) AuthorFirstName() string {
	nameList := strings.Fields(b.Author)
	if len(nameList) < 2 {
		return ""
	}
	return nameList[0]
}
