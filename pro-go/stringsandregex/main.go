package main

import (
	"fmt"
	"github.com/fatih/color"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	product := "Chandan Inc"
	color.Green("Contains: " + strconv.FormatBool(strings.Contains(product, "Chandan")))
	color.Green("Contains: " + strconv.FormatBool(strings.Contains(product, "chandan")))
	color.Blue("ContainsAny: " + strconv.FormatBool(strings.ContainsAny(product, "abc")))
	color.Black("ContainsRune: " + strconv.FormatBool(strings.ContainsRune(product, 'a')))
	color.Red("ContainsRune: " + strconv.FormatBool(strings.ContainsRune(product, 'z')))
	color.Red("EqualFold: " + strconv.FormatBool(strings.EqualFold(product, "Chandan")))
	color.Green("EqualFold: " + strconv.FormatBool(strings.EqualFold(product, "Chandan inc")))
	color.Green("HasPrefix: " + strconv.FormatBool(strings.HasPrefix(product, "Chandan")))
	color.Red("HasPrefix: " + strconv.FormatBool(strings.HasPrefix(product, "chandan")))
	color.Red("HasSuffix: " + strconv.FormatBool(strings.HasSuffix(product, "inc")))
	color.Green("HasSuffix: " + strconv.FormatBool(strings.HasSuffix(product, "Inc")))

	fmt.Println(strings.ToTitle(product))
	fmt.Println(strings.ToUpper(product))
	fmt.Println(strings.Title(product))
	fmt.Println(strings.ToLower(product))
	for _, char := range product {
		fmt.Println(string(char), char, "Upper Case: ", unicode.IsUpper(char))
	}

	isLetterH := func(c rune) bool {
		return c == 'H' || c == 'h'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(product, isLetterH))
	myStr := "Lorem Ipsum is simply     Lorem Ipsum of Lorem Ipsum and Lorem Ipsum.Lorem Ipsum and Lorem Ipsum."
	fields := strings.Fields(myStr)
	for _, field := range fields {
		fmt.Println("Field>>", field, "<<")
	}
	splits := strings.Split(myStr, "sum")
	for _, split := range splits {
		fmt.Println("Split>>" + split + "<<")
	}
	splitAfter := strings.SplitAfterN(myStr, "sum", 4)
	for _, split := range splitAfter {
		fmt.Println("SplitAfterN>>" + split + "<<")
	}

	pattern, compileErr := regexp.Compile("[A-z]oat")
	question := "Is that a goat?"

	if compileErr == nil {
		fmt.Println(pattern.MatchString(question))
		fmt.Printf("Type: %T\n", pattern)
		fmt.Printf("Go code: %#v\n", pattern)
		fmt.Printf("Recreate: %+v\n", pattern)
	} else {
		fmt.Println("Error: ", compileErr)
	}
}
