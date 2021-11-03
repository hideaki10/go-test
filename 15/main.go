package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

//
func main() {
	flag.Parse()
	var input string

	// read program input
	if flag.NArg() == 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(reader.)
		for i := 0; i < 3; i++ {

			var err error
			input, err = reader.ReadString('\n')
			if err != nil {
				log.Fatalln("failed to read input")
			}
			fmt.Println(input)
		}

		// from stdin/pipe

		//input = strings.TrimSpace(input) // otherwise, we would have a blank line
	} else { // from argument
		if flag.NArg() > 1 {
			log.Fatalln("takes at most one input")
		}
		input = flag.Arg(0)
	}

	fmt.Printf("> %s\n", input)
}
