package main

import  (
    "fmt"
)

 //converts infix reg expressions to postfix reg expressions
func intopost (infix string) string {
  //special characters
  specials := map[rune]int{'*': 10, '.': 9, '|': 8}

  //rune arrays
  postfix := []rune{}
  stack := []rune{}
	
  //loop through infix
  //convert string to array of runes by using range
  for _, r := range infix {
	switch{
	  case r == '(':
	    stack = append(stack, r)	

	  case r == ')':
		for stack[len(stack)-1] != '(' {
		  //gets rid of last element of stack
		  postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
		}
		
		//takes open bracket off end of the stack
		stack = stack[:len(stack)-1]

	  case specials[r] > 0:
		for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]]{
		  //gets rid of last element of stack
		  postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
		}
		
		stack = append(stack, r)

	  default:
		postfix = append(postfix, r)
	}//switch
  }//for

  for len(stack) > 0 {
	//gets rid of last element of stack
	postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
  }

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
}

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
}

func main () {
  //answer: ab.c*.
  fmt.Println ("Infix:   ", "a.b.c*")
  fmt.Println ("Postfix: ", intopost ("a.b.c*"))

  //answer: abd|.*
  fmt.Println ("Infix:   ", "(a.(b|d))*")
  fmt.Println ("Postfix: ", intopost ("(a.(b|d))*"))

  //answer: abd|.c*.
  fmt.Println ("Infix:   ", "a.(b|d).c*")
  fmt.Println ("Postfix: ", intopost("a.(b|d).c*"))

  //answer: abb.+.c
  fmt.Println ("Infix:   ", "a.(b.b)+.c")
  fmt.Println ("Postfix: ", intopost("a.(b.b)+.c"))

  nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
	
	//checking against reg expressin
  fmt.Println(pomatch("ab.c*|", "ab"))
}//main