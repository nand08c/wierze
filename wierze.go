package main

import (
	"fmt"
	"log"

	"github.com/nand08c/wierze/internal/commands"
)

func main() {
	fmt.Println("wiezre is for those unaware")

	err := commands.Init()
	if err != nil {
		log.Fatal(err)
	}
}
