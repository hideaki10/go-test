package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var clip string

func init() {
	flag.StringVar(&clip, "clip", "sdsd", "------")
	// TODO: フラグをパースする。
	flag.Parse()

}

func printResult(s string) error {
	fmt.Println(s)
	return nil
}

func run() error {
	reader := bufio.NewReader(os.Stdin)
	var err error
	input, err := reader.ReadString('\n')
	fmt.Println(input)
	if err != nil {
		return err
	}
	return printResult(input)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal("dsdsa")
	}
}
