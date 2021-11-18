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
	var input []string

	// read program input
	if flag.NArg() == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		for _, line := range input {
			fmt.Println(line + " : h")
		}
		// reader := bufio.NewReader(os.Stdin)
		// var err error
		// for {
		// 	input, err = reader.ReadString('\n')
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	fmt.Println(input)
		// }

		// from stdin/pipe

		//input = strings.TrimSpace(input) // otherwise, we would have a blank line
	} else { // from argument
		if flag.NArg() > 1 {
			log.Fatalln("takes at most one input")
		}
		//input = flag.Arg(0)
	}

	//fmt.Printf("> %s\n", input)
}
