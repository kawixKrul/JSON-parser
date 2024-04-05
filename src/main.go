package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileFlag := flag.String("file", "", "Use this flag if you want json to be read from file")
	jsonFlag := flag.String("json", "", "Use this flag if you want directly process json")

	flag.Parse()

	if *fileFlag != "" && *jsonFlag != "" {
		fmt.Println("Usage: ./parser [--file FILE | --json JSON]")
		os.Exit(1)
	}

	var validation bool

	if *fileFlag != "" {
		jsonStr, err := readFromFile(*fileFlag)
		if err != nil {
			fmt.Println("Error reading from file:", err)
			return
		}

		validation, err = validateJsonStr(jsonStr)
		if err != nil {
			fmt.Println("Error validating json:", err)
			return
		}
	}

	if *jsonFlag != "" {
		var err error = nil
		validation, err = validateJsonStr(*jsonFlag)
		if err != nil {
			fmt.Println("Error validating json:", err)
			return
		}
	}

	fmt.Println(validation)
}
