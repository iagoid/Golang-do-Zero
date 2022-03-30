package main

import (
	"fmt"
	"generics/blog"
	"generics/blog/entities"
	"generics/cache"
)

// https://youtu.be/CWczc9F0hBc

func main() {
	// Passa iniciar uma struct eu tenho que passar o tipo dela
	cc := cache.New[entities.Category]()
	categories := blog.GetCategories
	for _, category := range categories() {
		cc.Set(category.Slug, category)
	}

	cp := cache.New[entities.Post]()
	posts := blog.GetPosts
	for _, post := range posts() {
		cp.Set(post.Slug, post)
	}

	fmt.Println(cc.Get("generics"))
	fmt.Println("----------------")
	fmt.Println(cp.Get("tutorial-de-go"))

}
