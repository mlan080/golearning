package main

import (
	"school"
	"fmt"
)

func main(){
s := school.New()

	s.Add("Tomas", 3)
	fmt.Println(s)	
	s.Add("Tommy", 3)
	fmt.Println(s)	
	fmt.Println(s.Grade(3))
}