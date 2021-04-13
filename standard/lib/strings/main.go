package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := "Hello Go"
	fmt.Println(str)
	fmt.Println(str[0])
	fmt.Println(str[0:1])
	str1 := "This is a string"
	str2 := "This is another string"
	if str1 == str2 {
		fmt.Println("Strings are identical")
	} else {
		fmt.Println("Strings do not match")
	}

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("Strings match")
	} else {
		fmt.Println("Strings do not match")
	}
	fmt.Println(strings.ToLower(str1))

	ourStr := "This is my country"
	strcoll := strings.Split(ourStr, " ")
	fmt.Println(strcoll)
	ourStr = "This is my powerbank | I use it to charge my mobile|"
	strcoll = strings.SplitAfter(ourStr, "|")
	fmt.Println(strcoll)

	file, _ := os.Open("/Users/basantkumar/Documents/gobot/src/github.com/basant93/goexamples/standard/lib/strings/power.csv")
	scanner := bufio.NewScanner(file)
	fmt.Println(scanner.Scan())
	fmt.Println(file)
	for scanner.Scan() {
		fmt.Println("After scan")
		text := scanner.Text()
		items := strings.Split(text, ",")

		for i := range items {
			fmt.Println("----- New Reocrd -----")
			fmt.Println(items[i])
		}
	}

	sampleString := "This is go"
	result := strings.HasPrefix(sampleString, "go")
	fmt.Println(result)
	result = strings.HasSuffix(sampleString, "go")
	fmt.Println(result)
	result = strings.Contains(sampleString, "go")
	fmt.Println(result)

	sampleString = strings.Replace(sampleString, "go", "go programming lang", -1)
	fmt.Println(sampleString)

}

//string are read only slice of bytes. On index byte value is returned. Using slice range characters are returned
