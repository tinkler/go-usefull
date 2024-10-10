package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	jsonString, err := ConvertJsonFileToString("abi.json")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(jsonString)

}

func ConvertJsonFileToString(filePath string) (string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	strppedContent := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, string(fileContent))

	return strings.ReplaceAll(strppedContent, "\"", "\\\""), nil
}
