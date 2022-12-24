package helper

import (
	"fmt"
	"os"
)

func ReadFileToString(path string) string {
	content, error := os.ReadFile(path)

	if error != nil {
		fmt.Println(error)
	}

	return string(content)
}
