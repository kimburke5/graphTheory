package main

import  (
    "fmt"
)

 //converts infix reg expressions to postfix reg expressions
func intopost (infix string) string {
    //special characters
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}
    
    //look up rune
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
		}
	}

	for len(stack) > 0 {
	  //gets rid of last element of stack
	  postfix, stack = append(postfix, stack[len(stack)-1]), stack[:len(stack)-1]
	}

    return string(postfix) 
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

}