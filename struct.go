package main

import (
	"fmt"
)

type Movie struct {
	Title    string
	Released bool
	Length   int
}

func main() {
	movie := Movie{}
	movie.Title = "Wizard of Oz"
	movie.Released = true
	movie.Length = 125

	fmt.Println("Movie is \n", movie)

}
