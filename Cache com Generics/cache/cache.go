package cache

import "generics/blog/entities"

type Cacheable interface {
	entities.Category | entities.Post
}

type cache[T Cacheable] struct {
	data map[string]T
}

// Passa o T(Tipo) mass ele não é definido
func (c *cache[T]) Set(key string, value T) {
	c.data[key] = value
}

// Pego um arquivo de memso T(Tipo) passado
func (c *cache[T]) Get(key string) (v T){
	if v, ok := c.data[key]; ok {
		return v
	}
	return 
}

func New[T Cacheable]() *cache[T] {
	c := cache[T]{}
	c.data = make(map[string]T)

	return &c
}