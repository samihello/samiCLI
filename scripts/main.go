package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println()
	fmt.Println("Simulating compile and bundle of API specifications...")
	fmt.Println("(script by Anders in api/build/main.go)")
	fmt.Println("...by copying content to trigger a delta/change in API spec.")
	fmt.Println()

	output()
}

func output() {
	sourceFileWithChanges := "../api/docs.yaml"
	sourceFileWithoutChanges := "../api/no-changes.yaml"
	fmt.Println("Source file with changes: ", sourceFileWithChanges)
	fmt.Println("Source file without changes: ", sourceFileWithoutChanges)

	sourceFile := sourceFileWithChanges
	//sourceFile := sourceFileWithoutChanges

	fmt.Println()
	fmt.Println("Using source file: ", sourceFile)
	destinationFile := "../api/resources-api-docs.yaml"
	fmt.Println()

	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatalf("Failed to read source file: %v", err)
	}

	destDataBefore, err := ioutil.ReadFile(destinationFile)
	if err != nil {
		log.Fatalf("Failed to read destination file before copy: %v", err)
	}

	err = ioutil.WriteFile(destinationFile, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write to destination file: %v", err)
	}

	destDataAfter, err := ioutil.ReadFile(destinationFile)
	if err != nil {
		log.Fatalf("Failed to read destination file after copy: %v", err)
	}

	if string(destDataBefore) == string(destDataAfter) {
		fmt.Println("No changes detected in API specification.")
	} else {
		fmt.Println("Changes detected in API specification.")
	}
	fmt.Println()
	fmt.Println("---Simulation complete.")
	fmt.Println()
}
