package types

// Book: class
type Book struct {
	BookId         string
	P              Program
	BookTitle      string
	BookCurriculum string
	BookLabel      string
	BookStorage    string
	C              Condition
}

// BookSummary: class
type BookSummary struct {
	B            Book
	BookTotal    int
	BookAssigned int
}
