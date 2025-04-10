package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "World", "A name to say hello to")
	verbose := flag.Bool("verbose", false, "Print verbose output")
	flag.Parse()
	if *verbose {
		fmt.Printf("Starting the program...\n")
		fmt.Printf("Arguments received: name=%s, verbose=%t\n", *name, *verbose)
	}
	fmt.Printf("Hello, %s!\n", *name)
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println("Additional arguments:", args)
	}
	if *name == "" {
		fmt.Println("Error: Name cannot be empty")
		os.Exit(1)
	}
}
