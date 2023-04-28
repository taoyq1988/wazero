package main

import (
	"fmt"
	"github.com/tetratelabs/wazero/testasm/add"
)

func main() {
	fmt.Println(add.Add(2, 15))
}
