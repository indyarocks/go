package main

type BlankInterface interface {
}

func main() {
	x := []BlankInterface{
		"hello",
		1,
	}
	Printfln("Test blank interface", x)
}
