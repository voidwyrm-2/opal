package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/voidwyrm-2/opal/lexer"
)

func main() {
	showTokens := flag.Bool("t", false, "Print the lexer tokens")
	// showNodes := flag.Bool("n", false, "Print the parser nodes")

	flag.Parse()

	l := lexer.New(`// solution to challenge one, part one of Advent of Code 2015

fun elevator =
  0 if #argl == 0 else
  elevator [tail [#1]] +
  (1 if head [#1] == '(' else -1);

my content = grabfile ["input.txt"];

say [elevator [content]];`)
	toks, err := l.Lex()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if *showTokens {
		for _, t := range toks {
			fmt.Println(t.Str())
		}
	}
}
