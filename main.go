package main

import (
	"fmt"

	internal "github.com/breenbo/gator/internal/config"
)

func main() {
	config := internal.Read()
	config.SetUser("test")

	newConfig := internal.Read()

	fmt.Printf("%v\n", newConfig)
}
