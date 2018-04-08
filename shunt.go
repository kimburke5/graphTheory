//ID: G00269948 
//Author: Kimberly Burke
//Program: build a non-deterministic finite automaton (NFA) from a regular expression

package main

import  (
    "fmt"
)

//converts infix reg expressions to postfix reg expressions
//takes in a string (infix) returns a string
func intopost (infix string) string {
	//created map for special characters to map them into integers
	//map keeps track of the allowed special character
	// . meaning concatonate, | meaning or, * meaning the Kleene star
	specials := map[rune]int{'*': 10, '+':9, '?': 8, '.': 7, '|': 6}
	
	//rune is an alias for int32 and is equivalent to int32 in all ways. 
	//It is used, by convention, to distinguish character values from integer values. 
	//rune is a built in data type

	//rune arrays used as stacks
	//postfix reg expression
	postfix := []rune{}
	//stack to store operators from infix reg expression
  stack := []rune{}
	
  //loop through infix string convert to postfix string
	//convert string to array of runes by using range
	//return index of character that is currently being read
	//for _ ignors index number
  for _, r := range infix {

		switch{
			//if it reads an open bracket it puts it onto the end of the stack
	  	case r == '(':
	   	 stack = append(stack, r)	

			//if a closing bracket is read in the we are going to pop elements off the stack until we find an open bracket.
			//everything popped off stack here will be appended onto postfix/output
			 case r == ')':
				//while last character on the stack does not equal to an open bracket
				//len gives the length of an array
				for stack[len(stack)-1] != '(' {
		  		//gets rid of last element of stack
		  		postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
				}
		
				//takes open bracket off end of the stack
				stack = stack[:len(stack)-1]
			
				//specials to be grester than 0, to check if current character from infix is a special character
			case specials[r] > 0:
				//if the presidence of the current character that we are reading
				//is less than the presidence at the top of the stack
				for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]]{
		 	 		//gets rid of element from top of stack and appends into postfix
		  		postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
				}
				//append is a built in function that takes an array and appends an element onto the end of it
				//when the element at the top of the stack has less presidence
		  	//than the current character that we are reading
		  	//append current character onto the stack
				stack = append(stack, r)
			
			//r is neither a bracket or a special character	
	  	default:
				postfix = append(postfix, r)

		}//switch
  }//for

  for len(stack) > 0 {
		//gets rid of last element of stack
		postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
  }
	//cast postfix to a string as this function has to return a string
  return string(postfix) 
}//intopost

//works as linked list
type state struct {
  symbol rune
  edge1 *state
  edge2 *state
}

//pointers to states (from linked list struct)
type nfa struct {
  initial *state
  accept *state
}

func poregtonfa(postfix string) *nfa{
  //create stack
  nfastack := []*nfa{}

  for _, r := range postfix {

		switch r {
			case '.':
				//pops 2 frags off stack
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]
				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				//concat frag1 & frag2
				frag1.accept.edge1 = frag2.initial
			
				//push new frag to stack
				nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		
	  	case '|':
				//pops 2 frags off stack
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]
				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				//new accept states
				accept := state{}
				initial := state{edge1: frag1.initial, edge2: frag2.initial}
				frag1.accept.edge1 = &accept
				frag2.accept.edge1 = &accept

				//push new frag to stack
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
	
	 		case '*':
				//pops 1 frag off stack
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				//new accept state
				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}
				frag.accept.edge1 = frag.initial
				frag.accept.edge2 = &accept

				//push new frag to stack
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			
			case '+':
				//pops 1 frag off stack
				frag := nfaStack[len(nfaStack)-1]
				nfaStack = nfaStack[:len(nfaStack)-1]

				//new accept state
				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}
				frag.accept.edge1 = frag.initial

				//push new frag to stack
				nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

			case '?':
				//pops 1 frag off stack
				frag := nfaStack[len(nfaStack)-1]
				nfaStack = nfaStack[:len(nfaStack)-1]

				//new accept state
				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}
				frag.accept.edge1 = frag.initial

				//push new frag to stack
				nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

	 		default:
				//new accept state - empty
				accept:= state{}
				initial := state{symbol: r, edge1: &accept}
				//push new frag to stack
				nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
  	}//switch
  }//for

	if len(nfastack) != 1 {
		fmt.Println("Uh oh:", len(nfastack), nfastack)
	}
  return nfastack[0]
}//poregtonfa

//function gets current array, add s state to it 
func addState(l []*state, s *state, a *state) []*state {
	//append the state to the list that has been passed in
	l = append(l, s)

	//check that s is not equal to a, and any state that has its symbol as 0 (rune) 
	if s != a && s.symbol == 0{

		//checks if state has edge
		l = addState(l, s.edge1, a)

		//checks if state has a second edge
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}//addState

//checking if postfix reg expression po string matches s string
func pomatch(po string, s string) bool {
	//default position false
	ismatch := false
	//create Non-deterministic Finite Automaton (NFAs) from the reg expression
  ponfa := poregtonfa(po)

	//two lists of states the current and next state
	current := []*state{}
	//generate next state from current state
  next := []*state{}
	
	//Want current state to have more than just initial state we want 
	//it to have all states that it can access from the initial state through the arrows.
	//In this statement we will get list of current states that you are in, add the initial state, 
	//add the accepted states - this becomes the new current state
	//current[:] is a slice
  current = addState(current[:], ponfa.initial, ponfa.accept)

	//reading s a character at a time
  for _, r := range s {
		//loop through the current states takes in all of the current states 
	  for _, c := range current{
			//checks if current states are labled by charaters/rune read from s 
			//to see if their symbol is set to r (rune character)
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		//swap current for next, reset next to be an empty array
		current, next = next, []*state{}
  }
	//in a current set of states, loop through and check if any of them are the accept state
	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}
  return ismatch
}//pomatch

func main () {
	//answer: ab.c*.
	fmt.Println ("Infix:   ", "a.b.c*") //regular expression
	//calls into post and converts regular expression to postfix notation
  fmt.Println ("Postfix: ", intopost ("a.b.c*"))

  //answer: abd|.*
	fmt.Println ("Infix:   ", "(a.(b|d))*") //regular expressions
	//calls into post and converts regular expression to postfix notation
  fmt.Println ("Postfix: ", intopost ("(a.(b|d))*"))

  //answer: abd|.c*.
	fmt.Println ("Infix:   ", "a.(b|d).c*")	//regular expressions
	//calls into post and converts regular expression to postfix notation
  fmt.Println ("Postfix: ", intopost("a.(b|d).c*"))

  //answer: abb.+.c
	fmt.Println ("Infix:   ", "a.(b.b)+.c")	//regular expressions
	//calls into post and converts regular expression to postfix notation
  fmt.Println ("Postfix: ", intopost("a.(b.b)+.c"))

  nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
	
	//checking against reg expressin
	fmt.Println(pomatch("ab.c*|", "ab"))
	
}//main