package helpers

import (
	"log"
	"os"
	"strings"
)

func LoadData(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	str := string(data)
	return strings.Split(str, "\n")
}
