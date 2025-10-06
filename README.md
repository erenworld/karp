# karp

Interpreter for the Monkey programming language

Every interpreter is built to interpret a specific programming language. That’s how you “implement” a programming language. Without a compiler or an interpreter a programming language
is nothing more than an idea or a specification.

We’re going to parse and evaluate our own language called **Monkey**. It’s a language specifically
designed for this book.

* REPL

Expressed as a list of features, Monkey has the following:
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

### What Monkey looks like?
Why is it called “Monkey”? Well, because monkeys are magnificent, elegant,
fascinating and funny creatures. Exactly like our interpreter

### Interpreter vs compiler

What can be said is that the one fundamental attribute they all share is that they
take source code and evaluate it without producing some visible, intermediate result that can
later be executed. That’s in contrast to compilers, which take source code and produce output
in another language that the underlying system can understand.

 The interpreter gives them meaning!
