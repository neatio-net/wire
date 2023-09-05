package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/nio-net/common"
	"github.com/nio-net/wire/expr"
)

func main() {
	input := ""
	if len(os.Args) > 2 {
		input = strings.Join(os.Args[1:], " ")
	} else if len(os.Args) == 2 {
		input = os.Args[1]
	} else {
		fmt.Println("Usage: wire 'u64:1 u64:2 <sig:Alice>'")
		return
	}

	got, err := expr.ParseReader(input, strings.NewReader(input))
	if err != nil {
		Exit("Error parsing input: " + err.Error())
	}
	gotBytes, err := got.(expr.Byteful).Bytes()
	if err != nil {
		Exit("Error serializing parsed input: " + err.Error())
	}

	fmt.Println(Fmt("%X", gotBytes))
}
