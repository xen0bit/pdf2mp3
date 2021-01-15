package main

import (
	"bufio"
	"os"

	"github.com/viert/go-lame"
)

func main() {

	//argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	input := argsWithoutProg[0]
	output := argsWithoutProg[1]

	of, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer of.Close()
	enc := lame.NewEncoder(of)
	defer enc.Close()

	inf, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer inf.Close()

	r := bufio.NewReader(inf)
	r.WriteTo(enc)
}
