package main

import "WB-L2/pattern"

func main() {
	demonstrateFacade()
}

func demonstrateFacade() {
	f := pattern.NewFacade()
	f.Operation1()
	f.Operation2()
}
