# 🐟 Karp

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Dependabot Enabled](https://img.shields.io/badge/Dependabot-Active-brightgreen)](https://github.com/dependabot)
[![Build Status](https://img.shields.io/github/actions/workflow/status/yourusername/karp/go.yml?branch=main)](https://github.com/yourusername/karp/actions)

---

## 🧠 Overview

**Karp** is an interpreter for the **Karp programming language**, implemented in **Go (Golang)**.

Every interpreter is built to interpret a specific programming language — that’s how you *implement* a language.  
Without a compiler or interpreter, a programming language is nothing more than an idea or a specification.

Karp is a small, expressive, C-like language built for learning and experimentation.  
It is inspired by the book *["Writing an Interpreter in Go"](https://interpreterbook.com/)*.

---

## ⚙️ Pratt Parsing (Top-Down Operator Precedence)

Karp uses **Pratt Parsing** — a parsing technique that associates parsing functions with tokens, depending on their position (prefix, infix, or postfix).

- **Prefix operators** — appear *before* their operand:
  ```karp
  -5
  !true

- **Supported infix operators**
| Operator | Description    | Example  |
| -------- | -------------- | -------- |
| `+`      | Addition       | `5 + 5`  |
| `-`      | Subtraction    | `5 - 5`  |
| `*`      | Multiplication | `5 * 5`  |
| `/`      | Division       | `5 / 5`  |
| `>`      | Greater than   | `5 > 5`  |
| `<`      | Less than      | `5 < 5`  |
| `==`     | Equality       | `5 == 5` |
| `!=`     | Inequality     | `5 != 5` |

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