package main

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
)

func main() {
	file, err := os.Open("myfile.csv")
	if err != nil {
		fmt.Println("Ooops")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//grab a text string
		a := scanner.Text()
		matched, err := regexp.MatchString("Docker", a)
		if matched {
			fmt.Println(a)
		}
		if err != nil {
			fmt.Println("error in scanned line")
		}
	}

	if err := scanner.Err(); err != nil {
   		fmt.Println("Oh crap")
	}

}