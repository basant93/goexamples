package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "/help" {
		fmt.Println("Usage: ./app <filepath>")
	} else {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println("Error in opening the file")
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println(text)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error: ", err)
		}

	}

}
