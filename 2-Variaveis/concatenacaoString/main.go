package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

const MAX = 1000

func main() {
	normal()
	arrayAppend()
	writeString() // Melhor Performance
}

func normal() {
	var b string
	start := time.Now()

	for i := 0; i < MAX; i++ {
		b += strconv.Itoa(i)
	}
	// fmt.Println(b)
	fmt.Println("Normal", time.Since(start))
}

func arrayAppend() {
	var strs []string
	start := time.Now()

	for i := 0; i < MAX; i++ {
		strs = append(strs, strconv.Itoa(i))
	}

	// fmt.Println(strings.Join(strs, ""))
	fmt.Println("Array append", time.Since(start))
}

// Melhor Performance
func writeString() {
	var b bytes.Buffer
	start := time.Now()

	for i := 0; i < MAX; i++ {
		b.WriteString(strconv.Itoa(i))
	}
	// fmt.Println(b.String())
	fmt.Println("Write string", time.Since(start))
}
