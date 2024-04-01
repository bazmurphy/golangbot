package main

import (
	"fmt"
)

// we define an author struct
type author struct {
	firstName string
	lastName  string
	bio       string
}

// a method on the author struct
func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

// we define a blogPost struct
type blogPost struct {
	title   string
	content string
	// this uses the author struct above
	author
}

// a details method on the blogPost
func (p blogPost) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

// we define a website struct
type website struct {
	// this uses the blogPost struct above
	blogPosts []blogPost
}

// a contents method on the website
func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	// we create a website
	// that has blogPosts (a slice of blogPost)
	// and the blogPost have title, content and author (and details method)
	// and the author has firstname, lastname, bio (and fullName method)
	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
}

// Contents of Website

// Title:  Inheritance in Go
// Content:  Go supports composition instead of inheritance
// Author:  Naveen Ramanathan
// Bio:  Golang Enthusiast

// Title:  Struct instead of Classes in Go
// Content:  Go does not support classes but methods can be added to structs
// Author:  Naveen Ramanathan
// Bio:  Golang Enthusiast

// Title:  Concurrency
// Content:  Go is a concurrent language and not a parallel one
// Author:  Naveen Ramanathan
// Bio:  Golang Enthusiast
