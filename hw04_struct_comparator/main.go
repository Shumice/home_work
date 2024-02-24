package main

type Book struct {
	id     int
	title  string
	author string
	year   uint16
	size   int
	rate   float32
}

// Set методы
func (b *Book) setId(id int) {
	b.id = id
}
func (b *Book) setTitle(title string) {
	b.title = title
}
func (b *Book) setAuthor(author string) {
	b.author = author
}
func (b *Book) setYear(year uint16) {
	b.year = year
}
func (b *Book) setSize(size int) {
	b.size = size
}
func (b *Book) setRate(rate float32) {
	b.rate = rate
}

// Get методы
func (b *Book) Id() int {
	return b.id
}
func (b *Book) Title() string {
	return b.title
}
func (b *Book) Author() string {
	return b.author
}
func (b *Book) Year() uint16 {
	return b.year
}
func (b *Book) Size() int {
	return b.size
}
func (b *Book) Rate() float32 {
	return b.rate
}

type Enums int

const (
	Year Enums = iota
	Size
	Rate
)

type Comparator struct {
	CompareFunc Enums
}

func (c *Comparator) Compare(book *Book, book2 *Book) bool {
	switch c.CompareFunc {
	case Year:
		return book.year > book2.year
	case Size:
		return book.size > book2.size
	case Rate:
		return book.rate > book2.rate
	default:
		return false
	}
}

func main() {

}
