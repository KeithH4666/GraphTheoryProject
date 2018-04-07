package main

//Imports
import(
	"fmt"
)

//Structs
//Can only have two edges
type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}
//Helper struct
type nfa struct{
	initial *state
	accept *state
}
//Converts regular expression to an NFA (Address)
func regtonfa(postfix string)*nfa {
	nfastack := []*nfa{}
	
	//Switch statment to handle each character in the regular expresion
	for _, r := range postfix {
		switch r {
		case '.':
			//Get whats on the stack
			frag2:=nfastack[len(nfastack)-1]
			//store up to to the position
			nfastack = nfastack[:len(nfastack)-1]
			frag1:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			
			//Join both states together
			frag1.accept.edge1 = frag2.initial
			nfastack = append(nfastack, &nfa{initial:frag1.initial, accept: frag2.accept})
		case '|':
			//Similar to above
			frag2:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1:=nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			
			//New accept state and initial state
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
	//Return the only thing left on the stack
	return nfastack[0]
}

// Converts regular expression (infix) to postfix
func intoPost(infix string)string{
	//Order characters by precedence using a map
	specials := map[rune]int{'*': 10, '.': 9 ,'|':8}
	//Hold the postfix expression
	postfix:=[]rune{}
	//Temp holds characters before being put into postfix
	temp := []rune{}
	
	//Switch statement which orders the expression according to their precedence , Goes through each character
	for _, r := range infix{
		switch{
			//Add to temp stack
			case r == '(':
			temp = append(temp,r)
			
			//Append everything to post fix before the character ")"
			case r == ')':
				for temp[len(temp)-1] != '('{
					postfix = append(postfix,temp[len(temp)-1])
					temp = temp[:len(temp)-1]
				}
				
				temp = temp[:len(temp)-1]
			
			//Handles all of the special characters and orders them according to their precedence
			case specials[r] > 0:
				for len(temp) > 0 && specials[r] <= specials[temp[len(temp)-1]]{
					postfix = append(postfix,temp[len(temp)-1])
					temp = temp[:len(temp)-1]
				}
				temp = append(temp,r)
			
			default:
			//Adds to end of postfix array if no character matches
			postfix = append(postfix,r)
		
		}
	}
	
	for len(temp) > 0 {
		postfix = append(postfix,temp[len(temp)-1])
		temp = temp[:len(temp)-1]
	}
	
	//Return postfix as a string
	return string(postfix)

}

func addState(l []*state,s *state, a *state)[]*state{
	l  = append(l,s)
	
	if s != a && s.symbol == 0 {
		l = addState(l,s.edge1,a)
		if s.edge2 != nil{
			l = addState(l,s.edge2,a)
		}
	}
	
	return l
}

//Checks for a match
func pomatch(po string, s string) bool{

	ismatch := false
	ponfa := regtonfa(po)
	
	current := []*state{}
	next := []*state{}
	
	current = addState(current[:],ponfa.initial,ponfa.accept)
	
	for  _, r := range s {
		for _, c := range current{
			if c.symbol == r{
				next = addState(next[:],c.edge1,ponfa.accept)
			}
		}
		current, next = next, []*state{}
		
		
	}
	
	for _, c := range current{
		if c == ponfa.accept{
			ismatch = true
			break
		}
	}
	
	return ismatch

}

func main(){
	
	//Variables to hold user input
	var userInputReg string 
	var userInputString string 
	var postReg string 
	//var nfa string
	
	//Infix "a.b.c*"
	fmt.Println("Please enter regular expresion:")
	fmt.Scanln(&userInputReg)
	postReg = intoPost(userInputReg)
	fmt.Println("Converting into PostFix: Which is: ", postReg)
	
	//Regex Match
	fmt.Println("Regex match - Please enter string you want to test against regex:")
	fmt.Scanln(&userInputString)
	fmt.Println(pomatch(postReg,userInputString))

}