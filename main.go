package main

import (
	"eebinary/generate"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)


func main() {

	parser := argparse.NewParser("eebinary", "Embed Binary")
	i := parser.String("i", "input", &argparse.Options{Required: true, Help: "Input Binary"})
	o := parser.String("o", "output", &argparse.Options{Required: true, Help: "Output Binary"})
	w := parser.Flag("w", "windows", &argparse.Options{Required: false, Help: "Windows binary [Default is Linux]"})


	// Print any argument errors
	err := parser.Parse(os.Args)
	if err != nil { fmt.Print(parser.Usage(err)) }

	generate.Generate(i, o, w)

}