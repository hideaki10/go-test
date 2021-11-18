package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var clip string
var varOut []string

func init() {
	flag.StringVar(&clip, "clip", " ", "------")
	// TODO: フラグをパースする。
	flag.Parse()

}

func printResult(args []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		//fmt.Println(scanner.Text() + ":  " + clip)
		varOut = append(varOut, scanner.Text())

	}

	for _, value := range varOut {
		cmd := exec.Command(args[0], value)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println("rre")
			return
		}
		fmt.Println(string(stdout))

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "dwdw", err)
	}

}

func main() {
	args := flag.Args()
	printResult(args)
}
