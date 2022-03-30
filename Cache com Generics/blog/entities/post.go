package entities

type Post struct {
	ID       int
	Category *Category
	Title    string
	Text     string
	Slug     string
}
