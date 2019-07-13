package types

// Book struct represents collection Book
type Book struct {
	Id       string  `json:"id"`
	Isbn     string  `json:"isbn"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Author   *Author `json:"author"`
}

// Author represents author of book
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
