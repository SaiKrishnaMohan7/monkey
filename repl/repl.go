package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/saikrishnamohan7/monkey/lexer"
	"github.com/saikrishnamohan7/monkey/token"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.NewLexer(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok) // "%+v\n" prints struct fields
		}
	}
}