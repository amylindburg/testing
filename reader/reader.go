package main

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
	"log"
)

func main() {

	file, err := os.Open("/Users/amylindburg/go/src/github.com/amylindburg/testing/reader/myfile.csv")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
   		log.Fatalf("error in scanner: %v", err)
	}

	for scanner.Scan() {
		//grab a text string
		a := scanner.Text()
		matched, err := regexp.MatchString("Docker", a)
		if err != nil {
			log.Fatalf("error in regex: %v", err)
		}
		if matched {
			fmt.Println(a)
		}
	}	
}