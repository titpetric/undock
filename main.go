package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/davecgh/go-spew/spew"
	"github.com/moby/moby/builder/dockerfile/parser"
)

func main() {
	wd, _ := os.Getwd()
	f, err := os.Open("Dockerfile")
	if err != nil {
		log.Fatal(err)
	}
	result, err := parser.Parse(f)
	f.Close()

	for _, value := range result.AST.Children {
		fmt.Printf("# %s", value.Original)
		switch value.Value {
		case "add":
			fmt.Println()
			file := path.Join(wd, value.Next.Value)
			destination := value.Next.Next.Value

			if _, err := os.Stat(file); os.IsNotExist(err) {
				log.Fatal(err)
			}
			fmt.Println("cp", file, destination)
		case "copy":
			fmt.Println()
			spew.Dump(value, wd)
		case "run":
			fmt.Println()
			spew.Dump(value, wd)
		default:
			fmt.Println(" -- not implemented (" + value.Value + ")")
		}
	}
}
