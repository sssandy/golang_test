package pack1

import (
	"fmt"
	"../pack2"
)

func init () {
	fmt.Println("test pack1 init before")
}

func init () {
	fmt.Println("test pack1 init after")
}

func TestPack1 () {
	pack2.TestPack2()
}
