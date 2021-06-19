package main

import "os"
import "fmt"

import "github.com/akamensky/argparse"

func main() {
	parser := argparse.NewParser("ycli", "pongo2 template renderer")

	renderCmd := parser.NewCommand("render", "render a file")
	filename := renderCmd.String("f", "filename", &argparse.Options{Required: true, Help: "filename to render"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if renderCmd.Happened() {
		fmt.Println("Rendering ", *filename)
	}
}
