# karp

Interpreter for the Karp programming language

Every interpreter is built to interpret a specific programming language. That’s how you “implement” a programming language. Without a compiler or an interpreter a programming language
is nothing more than an idea or a specification.

We’re going to parse and evaluate our own language called **Karp**. It’s a language specifically
designed for this book.

In the Karp programming language everything besides let and return statements is an expression.

Top Down Operator Precedence (or: Pratt Parsing)
A crucial part of this idea is that each token type can have two parsing functions associated with it, depending
on the token’s position - infix or prefix.

A prefix operator is an operator “in front of” its operand. Example:
--5

Two prefix in Karp:  ! and -

A postfix operator is an operator “after” its operand. Example:
foobar++
 
5 * 8
The * operator sits in the infix position between the two integer literals 5 and 8. Infix operators
appear in binary expressions -

## Implementation of pratt parsing
As we saw before, a program in Karp is a series of statements. Some are let statements, others return statements.
We need to add a third type of statement to our AST: expression statements.


Expressed as a list of features, Karp has the following:
* C-like syntax
* variable bindings
* integers and booleans
* arithmetic expressions
* built-in functions
* first-class and higher-order functions
* closures
* a string data structure
* an array data structure
* a hash data structure

### What Karp looks like?
Why is it called “Karp”? Well, because karps are magnificent, elegant,
fascinating and funny creatures. Exactly like our interpreter

### Interpreter vs compiler

What can be said is that the one fundamental attribute they all share is that they
take source code and evaluate it without producing some visible, intermediate result that can
later be executed. That’s in contrast to compilers, which take source code and produce output
in another language that the underlying system can understand.

 The interpreter gives them meaning!
