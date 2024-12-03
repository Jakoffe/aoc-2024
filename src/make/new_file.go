package main

import (
	"aoc/src/utils"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < 24; i++ {
		folderPath := fmt.Sprintf("src/solutions/day%02d", i)

		filePath := fmt.Sprintf("%s/solution.go", folderPath)

		if _, err := os.Stat(filePath); err == nil {
			continue
		}

		// Create folder if not already present
		os.MkdirAll(folderPath, 0700)

		// Create file and exit
		templateContent := string(utils.ReadFile("src/templates/solution.tmpl"))
		template := []byte(fmt.Sprintf(templateContent, i, i, i))
		err := os.WriteFile(filePath, template, 0644)

		if err != nil {
			fmt.Println(err.Error())
		}

		os.Exit(0)
	}
}
