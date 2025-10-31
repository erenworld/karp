# 🐟 Karp


[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Dependabot Enabled](https://img.shields.io/badge/Dependabot-Active-brightgreen)](https://github.com/dependabot)


**Karp** is an interpreter for the Karp programming language, implemented in **Go (Golang)**.

Karp is a small, expressive, C-like language built for learning and experimentation.  
It is inspired by the book *["Writing an Interpreter in Go"](https://interpreterbook.com/)*.

## Example code

```js
let add = fn(x, y) { x + y };
add(2, 3);          // => 5

let makeAdder = fn(x) {
  fn(y) { x + y };
};

let addTwo = makeAdder(2);
addTwo(5);          // => 7

let result = !false;  // => true

let nums = [1, 2, 3, 4];
len(nums);            // => 4

let person = {"name": "Karp", "type": "fish"};
person["name"];       // => "Karp"
```

## Architecture
| Component       | Description                                                         |
| --------------- | ------------------------------------------------------------------- |
| **Lexer**       | Converts raw source code into tokens                                |
| **Parser**      | Builds an AST using Pratt Parsing                                   |
| **Evaluator**   | Interprets and executes the AST                                     |
| **Environment** | Stores variables and function scopes                                |
| **Objects**     | Defines data types like integers, booleans, strings, arrays, hashes |

A program in Karp is a sequence of statements, such as:
* let statements
* return statements
* Expression statements

## Getting started 

> prerequisite golang 1.22+

### On your terminal 
`go build -o .`

then 

 `./karp`


