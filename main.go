package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/voidwyrm-2/opal/interpreter/valuetypes/listtype"
	"github.com/voidwyrm-2/opal/interpreter/valuetypes/numbertype"
	"github.com/voidwyrm-2/opal/lexer"
)

func main() {
	showTokens := flag.Bool("t", false, "Print the lexer tokens")
	// showNodes := flag.Bool("n", false, "Print the parser nodes")

	flag.Parse()

	llist := listtype.New(numbertype.New(10), numbertype.New(15))
	llist.Append(numbertype.New(20))

	fmt.Println(llist.Fmt())

	if nl, err := llist.Concat(numbertype.New(25)); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println(nl.Fmt())
	}

	fmt.Println(llist.Fmt())

	l := lexer.New(``)
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
