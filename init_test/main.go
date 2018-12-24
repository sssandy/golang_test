package main

import "fmt"
import "./pack3"
import "./pack1"
// import "./pack3"

func init () {
	fmt.Println("main init")
}

func main () {
	pack3.TestPack3()
	pack1.TestPack1()
	pack1.TestPack11()

	fmt.Println("main")
}
