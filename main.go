package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
)

func main() {

	path := flag.String("path", "myapp.log", "the path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "log level search for OPTIONS : DEBUG, INFO, ERROR, CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
