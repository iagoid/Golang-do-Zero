package blog

import "generics/blog/entities"

func GetPosts() []entities.Post {
	return []entities.Post{
		{ID: 1, Category: &GetCategories()[0], Title: "Tutorial de Go", Text: "Tutorial simples sobre Golang", Slug: "tutorial-de-go"},
		{ID: 2, Category: &GetCategories()[1], Title: "Aprendendo generics", Text: "Video simples sobre generics", Slug: "aprendendo-generics"},
		{ID: 3, Category: &GetCategories()[2], Title: "Mais um exemplo de Go", Text: "Outro exemplo de go", Slug: "mais-um-exemplo-de-go"},
	}
}
