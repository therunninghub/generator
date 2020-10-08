package main

import (
	"fmt"
	"log"

	"therunninghub.net/generator"
)

func main() {
	output, err := generator.StringListWithCharset(5, generator.UpperAlphaNumeric, 100)
	if err != nil {
		log.Println("Something went wrong", err)
	}

	for _, r := range output {
		fmt.Println(r)
	}
}
