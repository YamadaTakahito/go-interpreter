package repl

import (
	"bufio"
	"fmt"
	"github.com/YamadaTakahito/go-interpreter/lexer"
	"github.com/YamadaTakahito/go-interpreter/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
