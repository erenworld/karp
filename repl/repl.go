package repl

import (
	"bufio"
	"fmt"
	"io"

	"karp/evaluator"
	"karp/lexer"
	"karp/object"
	"karp/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		eval := evaluator.Eval(program, env)
		if eval != nil {
			io.WriteString(out, eval.Inspect())
			io.WriteString(out, "\n")
		}
	}
}


func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Syntax Error: Unable to parse the input.\n")
	io.WriteString(out, " Details:\n")
	
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}