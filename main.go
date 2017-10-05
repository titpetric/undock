package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	_ "github.com/davecgh/go-spew/spew"
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

	clean := func(s string) string {
		s = strings.Replace(s, "\t", " ", -1)
		s = strings.Replace(s, "  ", " ", -1)
		return s
	}

	fmt.Println("#!/bin/bash")
	for _, value := range result.AST.Children {
		fmt.Printf("# %s", clean(value.Original))
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
			file := path.Join(wd, value.Next.Value)
			destination := value.Next.Next.Value
			if value.Next.Value[len(value.Next.Value)-1] == '/' {
				file += "/"
			}
			fmt.Println("rsync -ia", file, destination)
		case "run":
			fmt.Println()
			fmt.Println(clean(value.Next.Value))
		default:
			fmt.Println(" -- not implemented (" + value.Value + ")")
			// spew.Dump(value, wd)
		}
	}
}
