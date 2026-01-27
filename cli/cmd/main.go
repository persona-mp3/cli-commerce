package main

import (
	"fmt"
	"log"

	"github.com/persona-mp3/internal"
)

func main() {
	userChoice := internal.MainScreen()
	res, err := internal.QueryServer(userChoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", res)
}
