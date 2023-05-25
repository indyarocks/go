package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var names [3]string

	names[0] = "Chandan"
	names[1] = "Tithi"
	names[2] = "Darsh"

	for i, name := range names {
		fmt.Printf("names[%d]=%s\n", i, name)
	}

	var surNames [3]string = names
	fmt.Println(surNames)

	var unKnowns = [...]string{"fatal", "error", "warn", "info"}
	fmt.Println(unKnowns)

	unKnownCopy := [4]string{"error", "warn", "info", "fatal"}
	fmt.Println(unKnowns == unKnownCopy)
	fmt.Println("unKnownCopy:", unKnownCopy)
	fmt.Println("unKnownCopy leg", len(unKnownCopy))
	fmt.Println("unKnownCopy cap", cap(unKnownCopy))

	for unKnown := range unKnowns {
		fmt.Println("unKnowns", unKnown)
	}

	var mySlice = make([]string, 3, 6)
	fmt.Println("mySlice leg", len(mySlice))
	fmt.Println("mySlice cap", cap(mySlice))

	var slice1 = make([]string, 3, 10)
	slice1[0] = "Ram"
	slice1[1] = "Lakshman"
	slice1[2] = "Bharat"
	slice2 := append(slice1, "Satrudhan", "Luv", "Kush")
	slice3 := append(slice1, "Hanuman") // NOTE: Replaces Satrudhan in all the slices
	slice4 := append(slice1, slice2...)
	slice1[0] = "RamSita"
	fmt.Println("Slice1=============")
	for i, name := range slice1 {
		fmt.Printf("slice1[%d]=%s\n", i, name)
	}
	fmt.Println("Slice2=============")
	for i, name := range slice2 {
		fmt.Printf("slice2[%d]=%s\n", i, name)
	}
	fmt.Println("Slice3=============")
	for i, name := range slice3 {
		fmt.Printf("slice3[%d]=%s\n", i, name)
	}

	fmt.Println("Slice4=============")
	for i, name := range slice4 {
		fmt.Printf("slice4[%d]=%s\n", i, name)
	}

	myArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	someNums := myArray[1:3]
	allNums := myArray[:]
	fmt.Printf("myArray: %v, Size: %d, Capacity: %d\n", myArray, len(myArray), cap(myArray))
	fmt.Printf("someNums: %v, Size: %d, Capacity: %d\n", someNums, len(someNums), cap(someNums))
	fmt.Printf("allNums: %v, Size: %d, Capacity: %d\n", allNums, len(allNums), cap(allNums))
	fmt.Println("Appending 100 to someNums========")
	someNums = append(someNums, 100)
	fmt.Printf("someNums: %v, Size: %d, Capacity: %d\n", someNums, len(someNums), cap(someNums))
	fmt.Printf("Append: allNums: %v, Size: %d, Capacity: %d\n", allNums, len(allNums), cap(allNums))

	fmt.Println("Changed myArray========")
	myArray = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("myArray: %v, Size: %d, Capacity: %d\n", myArray, len(myArray), cap(myArray))
	fmt.Printf("someNums: %v, Size: %d, Capacity: %d\n", someNums, len(someNums), cap(someNums))
	fmt.Printf("allNums: %v, Size: %d, Capacity: %d\n", allNums, len(allNums), cap(allNums))

	products := [4]string{"Uber", "Ola", "SigNoz", "Dukaan"}
	someProducts := products[1:3]
	allProducts := products[:]

	someProducts = append(someProducts, "Chandan")
	fmt.Println("someProducts", someProducts)
	fmt.Println("someProducts len:", len(someProducts), "cap:", cap(someProducts))
	fmt.Println("allProducts", allProducts)
	fmt.Println("allProducts len:", len(allProducts), "cap:", cap(allProducts))

	fmt.Println("Exceeding the size of original capacity of partial slice =====")
	products1 := [4]string{"Uber", "Ola", "SigNoz", "Dukaan"}
	someProducts1 := products1[1:3]
	allProducts1 := products1[:]

	someProducts1 = append(someProducts1, "Chandan")
	someProducts1 = append(someProducts1, "Darsh")
	fmt.Println("Products1", products1)
	fmt.Println("products1 len:", len(products1), "cap:", cap(products1))
	fmt.Println("someProducts1", someProducts1)
	fmt.Println("someProducts1 len:", len(someProducts1), "cap:", cap(someProducts1))
	fmt.Println("allProducts1", allProducts1)
	fmt.Println("allProducts1 len:", len(allProducts1), "cap:", cap(allProducts1))

	fmt.Println("Exceeding the size of original capacity of partial slice with specified capacity =====")
	products2 := [4]string{"Uber", "Ola", "SigNoz", "Dukaan"}
	newSomeProducts := products2[1:3:3]
	allProducts2 := products2[:]
	fmt.Println("newSomeProducts", newSomeProducts)
	fmt.Println("newSomeProducts len:", len(newSomeProducts), "cap:", cap(newSomeProducts))
	newSomeProducts = append(newSomeProducts, "Chandan")
	newSomeProducts = append(newSomeProducts, "Darsh")
	fmt.Println("Products2", products2)
	fmt.Println("products1 len:", len(products2), "cap:", cap(products2))
	fmt.Println("newSomeProducts", newSomeProducts)
	fmt.Println("newSomeProducts len:", len(newSomeProducts), "cap:", cap(newSomeProducts))
	fmt.Println("allProducts2", allProducts2)
	fmt.Println("allProducts2 len:", len(allProducts2), "cap:", cap(allProducts2))

	var emptySlice = []string{}
	copy(emptySlice, products[:])
	fmt.Println(emptySlice)

	var myMap = map[string]float64{
		"chandan": 10,
		"tithi":   100,
		"darshu":  200,
		"sheenu":  300,
	}
	var keySlice = make([]string, 0, len(myMap))
	for key, _ := range myMap {
		keySlice = append(keySlice, key)
	}
	sort.Strings(keySlice)
	for _, key := range keySlice {
		fmt.Printf("myMap[%s]=%.2f\n", key, myMap[key])
	}

	var myStr string = "$24.99"
	for i := 0; i < len(myStr); i++ {
		fmt.Println("byte:", myStr[i])
		fmt.Println("string:", string(myStr[i]))
	}
	var currency byte = myStr[0]
	var currencyStr string = string(myStr[0])
	fmt.Println("currency byte:", currency, "currency str:", currencyStr)
	var amountString string = myStr[1:]
	if amount, parseErr := strconv.ParseFloat(amountString, 64); parseErr == nil {
		fmt.Println("Currency", string(currency), "Amount", amount)
	} else {
		fmt.Println("Parse Error", parseErr)
	}

	var myEuro = "€"
	fmt.Println("€ size is", len(myEuro))

	var myRune []rune = []rune("€24.99")
	var rCurrency = string(myRune[0])
	var rAmount = string(myRune[1:])
	fmt.Println("currency rune:", rCurrency)
	if amount, parseErr := strconv.ParseFloat(rAmount, 64); parseErr == nil {
		fmt.Println("Currency", string(rCurrency), "Amount", amount)
	} else {
		fmt.Println("Parse Error", parseErr)
	}

	var euroStr string = "€24.99"
	fmt.Println("Iterating over rune")
	for index, char := range euroStr {
		fmt.Println("Index:", index, "char:", char, "string(char):", string(char))
	}
	fmt.Println("Iterating over bytes")
	for index, char := range []byte(euroStr) {
		fmt.Println("Index:", index, "char:", char, "string(char):", string(char))
	}
}
