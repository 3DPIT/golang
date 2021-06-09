package main

import "src/mypackage"

func main() {
	value := mypackage.MyType{}
	value.ExportedMethod()
}
