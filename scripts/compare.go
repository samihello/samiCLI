package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	compare()
}

func compare() {
	// calculate the SHA hash of file in api/docs.yaml
	filePathAfter := "../api/docs.yaml"
	hashAfter, err := calculateFileHash(filePathAfter)
	if err != nil {
		log.Fatalf("Failed to calculate hash: %v", err)
	}

	// calculate the SHA hash of file in root docs.yaml
	filePathBefore := "../docs.yaml"
	hashBefore, err := calculateFileHash(filePathBefore)
	if err != nil {
		log.Fatalf("Failed to calculate hash: %v", err)
	}

	fmt.Printf("\nSHA256 hash of file AFTER %s: \n%x\n", filePathAfter, hashAfter)
	fmt.Printf("\nSHA256 hash of file BEFORE %s: \n%x\n", filePathBefore, hashBefore)

	if hashBefore != hashAfter {
		fmt.Println()
		log.Fatalf("API specification is not in sync. Detected differences. Please run the api/build/main.go script locally and commit the updated API specification.")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("API specification file has not changed.")
		fmt.Println()
	}
}

func calculateFileHash(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}
