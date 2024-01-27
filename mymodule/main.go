//Ref: https://www.digitalocean.com/community/tutorials/how-to-use-go-modules

package main

import (
	"fmt"
	"mymodule/newpackage"
)

func main() {
	fmt.Println("Implementing wc count")
	newpackage.PrintHello()
}
