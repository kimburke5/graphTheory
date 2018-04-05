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

    return string(postfix) 
}
func main () {
//answer: ab.c*.
fmt.Println ("Infix:   ", "a.b.c*")
fmt.Println ("Postfix: ", intopost ("a.b.c*"))

//answer: abd|.*
fmt.Println ("Infix:   ", "(a. (b|d))*")
fmt.Println ("Postfix: ", intopost ("(a. (b|d))*"))

//answer: abd|.c*.
fmt.Println ("Infix:   ", "a. (b|d).c*")
fmt.Println ("Postfix: ", intopost("a.(b|d).c*"))

//answer: abb.+.c
fmt.Println ("Infix:   ", "a. (b.b)+.c")
fmt.Println ("Postfix: ", intopost("a. (b.b)+.c"))

}