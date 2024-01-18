package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora/v4"
)

func main() {
	byt, err := os.ReadFile("./in.txt")
	if err != nil {
		fmt.Printf("Error: open in.txt got %s\n", aurora.Red(err))
		os.Exit(0)
	}
	var data map[string]interface{}
	err = json.Unmarshal(byt, &data)
	if err != nil {
		fmt.Printf("Error: umarshal %s got %s\n", aurora.Gray(0, string(byt)), aurora.Red(err))
		os.Exit(0)
	}

	err = os.WriteFile("./out.graphql", []byte(data["query"].(string)), 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", aurora.Red(err))
		os.Exit(0)
	}

	fmt.Printf("%s\n", aurora.Green("Success"))
}
