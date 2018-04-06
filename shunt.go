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

  return nfastack[0]
}//poregtonfa

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
}//main