package main

import (
	"github.com/fatih/color"
	"strconv"
	"strings"
)

func main() {
	product := "Chandan Inc"
	color.Green("Contains: " + strconv.FormatBool(strings.Contains(product, "Chandan")))
	color.Blue("ContainsAny: " + strconv.FormatBool(strings.ContainsAny(product, "abc")))
	color.Black("ContainsRune: " + strconv.FormatBool(strings.ContainsRune(product, 'a')))
	color.Red("ContainsRune: " + strconv.FormatBool(strings.ContainsRune(product, 'z')))
	color.Red("EqualFold: " + strconv.FormatBool(strings.EqualFold(product, "Chandan")))
	color.Green("EqualFold: " + strconv.FormatBool(strings.EqualFold(product, "Chandan inc")))
	color.Green("HasPrefix: " + strconv.FormatBool(strings.HasPrefix(product, "Chandan")))
	color.Red("HasPrefix: " + strconv.FormatBool(strings.HasPrefix(product, "chandan")))
	color.Red("HasSuffix: " + strconv.FormatBool(strings.HasSuffix(product, "inc")))
	color.Green("HasSuffix: " + strconv.FormatBool(strings.HasSuffix(product, "Inc")))
}
