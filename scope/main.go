package main

import (
	"myapp/packageone"
)

var myVar = "main package: myVar"

func main() {
	// variables
	// declare a package level variable for the main package name myVar

	// declare a block variable for the main fucntion called blockVar
	var blockVar = "Block Var: blockVar"

	// declare a package level variable in the packeone package named PackageVar

	// create an exported function in packageon called PrintMe

	// in the main function, print out the values of myVar, blockVar and PackageVar on one line
	//  using the PrintMe function in packageone
	packageone.PrintMe(myVar)
	packageone.PrintMe(blockVar)
	packageone.PrintMe(packageone.PackageVar)

}
