package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println()
	fmt.Println("---Simulating compile and bundle of API specifications---")
	fmt.Println()

	simulateCompileAndBundle()
}

func simulateCompileAndBundle() {
	sourceFileWithChanges := "dist/compiled/with-changes.yaml"
	sourceFileWithoutChanges := "dist/compiled/no-changes.yaml"
	fmt.Println("Source file with changes: ", sourceFileWithChanges)
	fmt.Println("Source file without changes: ", sourceFileWithoutChanges)

	sourceFile := sourceFileWithChanges
	////sourceFile := sourceFileWithoutChanges

	fmt.Println()
	fmt.Println("Using source file: ", sourceFile)
	if sourceFile == sourceFileWithChanges {
		fmt.Println("---Simulating changes in API specification---")
	} else {
		fmt.Println("---Simulating no changes in API specification---")
	}
	fmt.Println()
	destinationFile := "../api/resources-api-docs.yaml"

	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatalf("Failed to read source file: %v", err)
	}

	before, err := ioutil.ReadFile(destinationFile)
	if err != nil {
		log.Fatalf("Failed to read destination file before copy: %v", err)
	}

	// Copy the source file to the destination file
	err = ioutil.WriteFile(destinationFile, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write to destination file: %v", err)
	}

	after, err := ioutil.ReadFile(destinationFile)
	if err != nil {
		log.Fatalf("Failed to read destination file after copy: %v", err)
	}

	if string(before) == string(after) {
		fmt.Println()
		fmt.Println("No changes detected in API specification.")
	} else {
		fmt.Println()
		fmt.Println("Changes detected in API specification.")
	}
	fmt.Println()
}
