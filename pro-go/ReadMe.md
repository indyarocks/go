## Go packages
1. https://pkg.go.dev
2. https://github.com/golang/go/wiki/Projects

### Struct initialisation
var product = new(Product)
OR
var product = Product{}

### Compositions
field promotion - If the enclosing type has a field same as a field of
embedded type's field, then the direct access to the attribute will
render enclosing type's field value.
Fields and Methods can be promoted to the enclosing type from embedded type
```go
type Boat struct {
        *Product
        Capacity  int
        Motorized bool
}

type Product struct {
        Name, Category string
        price          float64
}

func (p *Product)Price(tax) float64{
	return p.price + p.price*tax
}
```
`Name` can be accessed by either `boat.Product.Name` and `boat.Name`
`boat.Price(0.2)` or `boat.Product.Price(0.2)` - Both are equivalent.

```go
type Product struct {
	Name, Category string
	Price float64
}

type ProductGroup []*Product

func (group ProductGroup) CalcTotal() float64{
	...
}
```
A receiver function can only be defined on `type`, so a receiver function here
can't be defined on `[]*Product`, thus we need to create `alias type` -
`ProductGroup` so that it can have a receiver method.

### Interface
Since GoLang doesn't support inheritance, it provides `interface` to provide similar functionality.

### Adapter to execute a synchornous function asynchronously
```go
calcTax := func(price float64) float64 {
	return price*0.2
}

resultChannel := make(chan float64)
go func(price float64, c chan float64) {
	c <- calcTax(price)
}(275, resultChannel) // Price and channel passed as arguments
```

The `for` loop and `range` on chan ranges over just channel, it doesn't return index
```go
   // Range over channel just returns the message on channel
   testProductChannel := make(chan *Product, 5)
	go nonBlockingSendEnumerateProducts(chan<- *Product(testProductChannel))
	time.Sleep(time.Second)
	for p := range testProductChannel {
		fmt.Println("Received", p.Name)
	}
	// Range over slice has index and object from slice
	slice := []string{"Chandan", "Alok", "Ruby", "Darsh", "Tithi", "Sheenu"}
	for i, str := range slice {
		fmt.Println(i, str)
    }
```

NOTE: goroutine functions usually doesn't have return value, instead it sends the message via channel
### Slice
When defining a slice on a struct, the elements can be direct definition of struct as below:
```go
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Product struct {
	Name string
}

var pSlice = []*Product{
	{"Chandan"},
	{"Alok"},
}

var vSlice = []Product{
	{"Chandan"},
	{"Alok"},
}

func main() {
	for _, p := range pSlice {
		fmt.Println(p.Name)
	}
	for _, p := range vSlice {
		fmt.Println(p.Name)
	}
}

```

### Error interface
Go provides a predefined interface named error that provides one way to create own error types(structs)
```go
  type error interface {
	Error()
}
```
This inteface requires errors to define a method named Error, which returns a string
Example:
```go
  type CategoryError struct {
    requestedCategory string	
  }
  
  func(category *CategoryError) Error() string {
    return "ERROR: " + category + "is invalid"	  
  }

// NOTE: As CategoryError is a struct implementing Error interface
// Thus &CategoryError{requestedCategory: category} will print the error message
// i.e. returned string of Error() function
func main() {
  // This is equivalent to fmt.Println(string) 
  // where string is returned value of Error() function above
  fmt.Println(&CategoryError{requestedCategory: category})	 
  
}

// =============== ANOTHER example
package errors_test

import (
"fmt"
"time"
)

// MyError is an error implementation that includes a time and message.
type MyError struct {
  When time.Time
  What string
}

func (e MyError) Error() string {
  return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
  return MyError{
  time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
  "the file system has gone away",
  }
}

func Example() {
  if err := oops(); err != nil {
    fmt.Println(err)
  }
// Output: 1989-03-15 22:30:00 +0000 UTC: the file system has gone away
}
```

### %v specifierType of object:
1. Use `%T` to get the Go type of a value - As we get the class of objects in other OOP languages
2. Use `%#v` format that could be used to re-create the value in Go code file
3. Use `%v` default format of the value
4. Use `%+v` includes the field name

```go
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
	// Output 
	true
	Type: *regexp.Regexp
    Go code: &regexp.Regexp{expr:"[A-z]oat", prog:(*syntax.Prog)(0x14000072240), onepass:(*regexp.onePassProg)(nil), numSubexp:0, maxBitStateLen:43690, subexpNames:[]string{""}, prefix:"", prefixBytes:[]uint8(nil), prefixRune:0, prefixEnd:0x0, mpool:0, matchcap:2, prefixComplete:false, cond:0x0, minInputLen:4, longest:false}
    Recreate: [A-z]oat

```

### Format string for a struct
User `Stringer` interface from `fmt` package
```go
type Stringer interface {
	String() string
}
```