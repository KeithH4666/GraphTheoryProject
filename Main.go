package main

import(
	"fmt"
)

func intoPost(infix string)string{
	specials := map[rune]int{'*': 10, '.': 9 ,'|':8}
	postfix:=[]rune{}
	temp := []rune{}
	
	return string(postfix)

}

func main(){

	fmt.Println("Infix: a.b.c*")
	fmt.Println("Profix: ", intoPost("a.b.c*"))

}