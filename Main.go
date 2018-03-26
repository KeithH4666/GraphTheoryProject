package main

import(
	"fmt"
)

type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}

type nfa struct{
	initial *state
	accept *state
}

func regtonfa(postfix string)*nfa {
	nfastack := []*nfa{}
	
	for _, r := range postfix {
		switch r {
		case '.':
			frag2:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
				
			frag1.accept.edge1 = frag2.initial
				
			nfastack = append(nfastack, &nfa{initial:frag1.initial, accept: frag2.accept})
		case '|':
			frag2:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
				
			accept := state{}
			initial := state{edge1: frag1.initial,edge2:frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
				
			nfastack = append(nfastack, &nfa{initial:&initial, accept: &accept})
		case '*':
			frag:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
				
			accept:= state{}
			initial:= state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
				
			nfastack = append(nfastack, &nfa{initial:&initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r,edge1: &accept}
				
			nfastack = append(nfastack, &nfa{initial:&initial,accept:&accept})
		}
	}
	
	return nfastack[0]
}

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
	
	fmt.Println("Profix: ab.c*")
	fmt.Println("NFA: ", regtonfa("ab.c*|"))

}