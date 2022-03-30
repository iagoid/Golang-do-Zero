package blog

import "generics/blog/entities"

func GetCategories() []entities.Category {
	return []entities.Category{
		{ID: 1, Name: "Generics", Slug: "generics"},
		{ID: 2, Name: "Tutoriais", Slug: "tutorias"},
		{ID: 3, Name: "Videos", Slug: "videos"},
	}
}
