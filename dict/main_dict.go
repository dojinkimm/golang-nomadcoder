package main

import (
	"fmt"
	"github.com/dojinkimm/golang_study/dict"
)

// Dict main Function
func main() {
	dictionary := dict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil{
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}

