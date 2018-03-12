package main

import(
	"fmt"
)

func intoPost(infix string)string{
	specials := map[rune]int{'*': 10, '.': 9 ,'|':8}
	postfix:=[]rune{}
	temp := []rune{}
	
	for _, r := range infix{
		switch{
			case r == '(':
			temp = append(temp,r)
			
			case r == ')':
				for temp[len(temp)-1] != '('{
					postfix = append(postfix,temp[len(temp)-1])
					temp = temp[:len(temp)-1]
				}
				
				temp = temp[:len(temp)-1]
			
			case specials[r] > 0:
				for len(temp) > 0 && specials[r] <= specials[temp[len(temp)-1]]{
					postfix = append(postfix,temp[len(temp)-1])
					temp = temp[:len(temp)-1]
				}
				temp = append(temp,r)
			
			default:
			postfix = append(postfix,r)
		
		}
	}
	
	for len(temp) > 0 {
		postfix = append(postfix,temp[len(temp)-1])
		temp = temp[:len(temp)-1]
	}
	
	return string(postfix)

}

func main(){

	fmt.Println("Infix: a.b.c*")
	fmt.Println("Profix: ", intoPost("a.b.c*"))

}