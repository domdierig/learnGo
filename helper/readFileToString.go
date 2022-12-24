package helper

import (
	"fmt"
	"os"
)

func ReadFileToString(path string) (string, error) {
	content, error := os.ReadFile(path)

	if error != nil {
		fmt.Println(error)
		return "", error
	}

	return string(content), nil
}
