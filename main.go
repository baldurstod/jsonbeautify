package main

import (
	"encoding/json"
	"os"
	"strings"
	"path"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No source file provided")
		os.Exit(1)
	}

	for _, val := range os.Args[1:] {
		processFile(val)
	}
}

func processFile(filename string) {
	fmt.Println("Processing file ", filename)
	var jsonData interface{}

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}


	j, _ := json.MarshalIndent(&jsonData, "", "\t")

	filename = strings.TrimSuffix(filename, path.Ext(filename))

	os.WriteFile(filename + "_beautified.json", j, 0666)
}
