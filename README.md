# graphTheory

# Problem statement - Brief
You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗
” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.
Other special characters you might consider allowing as input are brackets
“()” which can be used for grouping, “+” which means “at least one of”, and
“?” which means “zero or one of”. You might also decide to remove the
concatenation character, so that 1.0 becomes 10, with the concatenation
implicit.
You may initially restrict the non-special characters your program works
with to 0 and 1, if you wish. However, you should at least attempt to expand
these to all of the digits, and the characters a to z, and A to Z.
You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed. You might do that for this project as follows:
1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA

# Project
This is a project for the module graph theory as a part of third year Software Development in GMIT. I created the project through Golang programming language to build a non-deterministic finite automaton (NFA) from a regular expression and can use the NFA to check if the regular expression matches any given string of text. This project used two graphing algorithms the Shunting yard algorithm and the Thompsons construction algorithm.

# The Shunting Yard Algorithm
The shunting-yard algorithm is a method to convert an infix expression into a postfix expression. The algorithm was invented by Edsger Dijkstra and named the "shunting yard" algorithm because its operation resembles that of a railroad shunting yard. There is also a stack that holds operators not yet added to the output queue. The purpose of the stack is to reverse the order of the operators in the expression. To convert, the program reads each symbol in order and does something based on that symbol.

# The Thompsons Construction Algorithm
Thompson's construction is an algorithm for transforming a regular expression into an equivalent nondeterministic finite automaton (NFA). Hence, this algorithm is of practical interest, since it can compile regular expressions into NFAs.
The algorithm works recursively by splitting an expression into its constituent subexpressions, from which the NFA will be constructed using a set of rules.

# Regular Expression Characters
- * (Kleene Star) character in a regular expression means "match the preceding character zero or many times". For example A* matches       any number (including zero) of character 'A'. Stephen Kleene (1909-1994) was one of the early investigators of regular expressions and   finite automata to which the character was named after.
- + The plus sign is the match-one-or-more quantifier.
- ? The question mark is the match-zero-or-one quantifier. The question mark is also used in special constructs with parentheses and in   changing match behaviour.
- | the vertical pipe meaning or separates a series of alternatives.
    Example: "(a|b|c)a" matches "aa" or "ba" or "ca".
- . the dot matches any character except the newline symbol.
    Example: ".a" matches two consecutive characters where the last one is "a".
    Example: ".*\.txt$" matches all strings that end in ".txt".

# How to clone your repository in Github
1.	On GitHub, navigate to the main page of the repository.
2.	Under the repository name, click Clone or download.
3.	In the Clone with HTTPs section, click to copy the clone URL for the repository.
4.	Open Git Bash.
5.	Change the current working directory to the location where you want the cloned directory to be made.
6.	Type git clone, and then paste the URL you copied in Step 2.
7.	Press Enter.
8.	Your local clone will be created.

To run this program navigate to the folder that the program is in through the terminal, then enter command - go run "shunt.go"  (to run the program).

# References
https://en.wikipedia.org/wiki/Shunting-yard_algorithm
http://www.oxfordmathcenter.com/drupal7/node/628
https://golang.org/pkg/builtin/#rune
https://chortle.ccsu.edu/FiniteAutomata/Section07/sect07_16.html
https://en.wikipedia.org/wiki/Thompson%27s_construction
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
https://learnonline.gmit.ie/
https://stackoverflow.com/questions/22693275/what-does-asterisk-struct-notation-mean-in-golang
http://www.fon.hum.uva.nl/praat/manual/Regular_expressions_1__Special_characters.html

