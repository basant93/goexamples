package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	sampleStr := "This is string text. String is a sentence"
	r, _ := regexp.Compile(`s([a-z]+)g`)
	fmt.Println(r.MatchString(sampleStr))
	fmt.Println(r.FindAllString(sampleStr, -1))

	str1 := "  I am string with spaces, trim me!!   "
	fmt.Println("%q \n", str1)
	fmt.Println(strings.TrimSpace(str1))
	fmt.Println(strings.TrimLeft(str1, " "))
	url := "https://www.godoc.org"
	domain := strings.TrimPrefix(url, "https://")
	fmt.Println(domain)
	filename := "abc.txt"
	file := strings.TrimSuffix(filename, ".txt")
	fmt.Println(file)
	input := "This is a paragraph."
	words := strings.Fields(input)
	fmt.Println(words)
	stopwords := "a an the"

	titledText := make([]string, 5)

	for index, word := range words {
		if strings.Contains(stopwords, word) {
			titledText[index] = word
		} else {
			titledText[index] = strings.ToUpper(word)
		}
	}

	fmt.Println(strings.Join(titledText, " "))
}
