package main

import (
	"log"
	"os"
)

//Log levels : info, warning, error, debug, fatal
//fatal is used when program can no longer continue to run
//log state has a dat time at start
func main() {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening the file, Error: ", err)
	}

	log.SetOutput(file)
	log.Println("This is a main function")
	log.Println("This is a main function")
	logger := log.New(file, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("This is a log message")

}
