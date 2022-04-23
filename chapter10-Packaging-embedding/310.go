package main

import "mypackage"

func main() {
	value := mypackage.MyType{}
	value.ExportedMethod()
//	value.unexportedMethod()
}
